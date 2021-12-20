package websiteService

import (
	"backend/db"
	"backend/libs/helper"
	"backend/libs/safe"
	"backend/models"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

//https://user.95516.com/favicon.ico
//检测域名库和IP库isNew字段是否为true，为true就拿过来测试站点是否能访问，能访问就记录相关数据

type mq struct {
	domain string
	cid    int
}

type WebSiteService struct {
	wsCache      sync.Map
	ctx          context.Context
	cancel       context.CancelFunc
	monitorTime  time.Duration
	monitorQueue []mq
	threadnum    int
	tch          chan mq
	client       http.Client
}

func NewWebSiteService() *WebSiteService {
	ctx, cancel := context.WithCancel(context.Background())
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	return &WebSiteService{
		ctx:         ctx,
		cancel:      cancel,
		monitorTime: wesiteMonitorTime,
		threadnum:   30,
		tch:         make(chan mq, 10),
		client:      http.Client{Timeout: 2 * time.Second, Transport: tr},
	}
}

var wesiteMonitorTime = time.Hour * 24

func StartWebSiteMonitorService() {
	fmt.Println("[*] start website monitor service.")
	ws := NewWebSiteService()
	ws.initCache()
	go ws.loopUpdateCache()
	go ws.worker()
	ws.startWebSiteMonitorService()

}

func (ws *WebSiteService) startWebSiteMonitorService() {

	ticker := time.NewTicker(ws.monitorTime)
	for {
		select {
		case <-ticker.C:
			safe.Go(func() error {
				ws.monitorWebSiteFromDomain()
				return nil
			}, nil)
			safe.Go(func() error {
				ws.monitorWebSiteFromIP()
				return nil
			}, nil)
		case <-ws.ctx.Done():
			ws.cancel()
		}
	}
}

func (ws *WebSiteService) initCache() {
	var wesites []string
	err := db.Orm.Debug().Model(&models.Websites{}).Select("domain").Find(&wesites).Error
	if err != nil {
		fmt.Println("[!] websiteService.go line:201 [", err, "]")
	}

	for _, val := range wesites {
		ws.wsCache.Store(val, val)
	}
}

func (ws *WebSiteService) loopUpdateCache() {
	var domains []string
	ticker := time.NewTicker(24 * time.Hour)
	for {
		select {
		case <-ticker.C:
			err := db.Orm.Debug().Model(&models.Domain{}).Select("domain").Where("isNew=false").Find(&domains).Error
			if err != nil {
				fmt.Println("[!] websiteService.go line:80 [", err, "]")
			}

			for _, d := range domains {
				if _, ok := ws.wsCache.Load(d); ok {
					ws.wsCache.Delete(d)
				}
			}

		case <-ws.ctx.Done():
			return

		}
	}
}

func (ws *WebSiteService) monitorWebSiteFromDomain() {

	var domains []models.Domain

	////每次重新读取数据库的时候清空一次
	//ws.monitorQueue=nil
	if len(ws.monitorQueue) <= 0 {
		err := db.Orm.Debug().Model(&models.Domain{}).Select("domain", "cid").Where("isNew=true").Find(&domains).Error
		if err != nil {
			fmt.Println("[!] websiteService.go line:80 [", err, "]")
		}
		if len(domains) <= 0 {
			return
		}
		for _, v := range domains {
			var m mq
			m.domain = v.Domain
			m.cid = v.Cid
			ws.monitorQueue = append(ws.monitorQueue, m)

		}
	}

}

func (ws *WebSiteService) monitorWebSiteFromIP() {

}

func (ws *WebSiteService) worker() {
	tk := time.NewTicker(ws.monitorTime)

	//开启多个扫描groutine，太多数据库怕抗不住
	for i := 0; i < ws.threadnum; i++ {
		go ws.run()
	}

	for {
		select {
		case <-tk.C:
			for i := len(ws.monitorQueue) - 1; i > 0; i-- {
				ws.tch <- ws.monitorQueue[i]
			}
			ws.monitorQueue = nil
		}
	}
}

func (ws *WebSiteService) ScanWebsiteInfo() error {

	var domains []models.Domain
	err := db.Orm.Debug().Model(&models.Domain{}).Select("domain", "cid").Where("isNew=true").Find(&domains).Error
	if err != nil || len(domains) <= 0 {
		fmt.Println("[!] websiteService.go line:80 [", err, "]")
		return errors.New("database error or not found new?")
	}
	go func() {
		for _, v := range domains {
			alivesite, err := ws.getAliveSite(v.Domain)
			if err != nil {
				continue
			}
			for _, alv := range alivesite {
				//f:= true
				website := models.Websites{IsNew: new(bool), CDN: new(bool)}
				title, headers, errx := ws.getTitleandHeader(alv)
				if errx != nil {
					continue
				}
				website.Cid = v.Cid
				website.Website = alv
				website.Headers = headers
				var tmpstr string
				for _, v := range title {
					if ('\u4e00' < v && v < '\u9fa5') || ('\x00' < v && v < '\x7e') {
						tmpstr += string(v)
					}
				}
				website.Domain = v.Domain
				website.Title = tmpstr
				tmpstr = ""
				website.UpdateTime = helper.GetCurTime()
				favhash, favhashUrl, err := getFaviconHash(alv)
				website.FaviconUrl = favhashUrl
				ips, err := ws.getIP(v.Domain)
				if err == nil {
					if len(ips) > 1 {
						*website.CDN = true
					} else {
						website.Ips = ips[0]
						//添加IP到IP数据库
						var tip = models.Ips{IsNew: new(bool)}
						tip.Ip = ips[0]
						tip.Cid = v.Cid
						*tip.IsNew = true
						tip.UpdateTime = helper.GetCurTime()
						err = db.Orm.Debug().Model(&models.Ips{}).Create(&tip).Error
						if err != nil {
							fmt.Println("[!] websiteService.go line:244 [", err, "]")
						}
					}
				}
				*website.IsNew = true
				if err == nil {
					website.Favicon = favhash
				}
				err = db.Orm.Debug().Model(&models.Websites{}).Create(&website).Error
				if err != nil {
					fmt.Println("[!] websiteService.go line:164 [", err, "]")
				}

			}
		}
	}()

	return nil
}

func (ws *WebSiteService) run() {
	for v := range ws.tch {
		//查找当前缓存，当前站点是不是已经存在了
		_, ok := ws.wsCache.Load(v.domain)
		if ok {
			continue
		}
		//只要走到这一步就把域名添加到缓存里。
		ws.wsCache.Store(v.domain, v.domain)
		alivesite, err := ws.getAliveSite(v.domain)
		if err != nil {
			continue
		}
		for _, alv := range alivesite {
			//f:= true
			website := models.Websites{IsNew: new(bool), CDN: new(bool)}
			title, headers, errx := ws.getTitleandHeader(alv)
			if errx != nil {
				continue
			}
			website.Cid = v.cid
			website.Website = alv
			website.Headers = headers
			var tmpstr string
			for _, v := range title {
				if ('\u4e00' < v && v < '\u9fa5') || ('\x00' < v && v < '\x7e') {
					tmpstr += string(v)
				}
			}
			website.Domain = v.domain
			website.Title = tmpstr
			tmpstr = ""
			website.UpdateTime = helper.GetCurTime()
			favhash, favhashUrl, err := getFaviconHash(alv)
			website.FaviconUrl = favhashUrl
			ips, err := ws.getIP(v.domain)
			if err == nil {
				if len(ips) > 1 {
					*website.CDN = true
				} else {
					website.Ips = ips[0]
					//添加IP到IP数据库
					var tip = models.Ips{IsNew: new(bool)}
					tip.Ip = ips[0]
					tip.Cid = v.cid
					*tip.IsNew = true
					tip.UpdateTime = helper.GetCurTime()
					err = db.Orm.Debug().Model(&models.Ips{}).Create(&tip).Error
					if err != nil {
						fmt.Println("[!] websiteService.go line:244 [", err, "]")
					}
				}
			}
			*website.IsNew = true
			if err == nil {
				website.Favicon = favhash
			}
			err = db.Orm.Debug().Model(&models.Websites{}).Create(&website).Error
			if err != nil {
				fmt.Println("[!] websiteService.go line:164 [", err, "]")
			}

		}
	}
	fmt.Println("run done")
}

func (ws *WebSiteService) getAliveSite(target string) ([]string, error) {
	var testPort = []string{"80", "443", "8080"}
	var aliveUrl []string
	var url string
	for _, v := range testPort {
		if v == "443" {
			url = "https://" + target
		} else {
			url = "http://" + target + ":" + v
		}
		_, err := ws.client.Get(url)
		if err != nil {
			continue
		}
		aliveUrl = append(aliveUrl, url)
		url = ""
	}
	if len(aliveUrl) == 0 {
		return nil, errors.New("no alive")
	}
	return aliveUrl, nil
}

func (ws *WebSiteService) getIP(domain string) (ip []string, err error) {
	//先获取IP地址
	ips, err := net.LookupIP(domain)
	if err != nil {
		//fmt.Println("Resolution error", err.Error())
		return
	}

	for _, v := range ips {
		ip = append(ip, v.String())
	}
	return
}

func (ws *WebSiteService) getTitleandHeader(url string) (title string, header string, errs error) {

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36 Edg/95.0.1020.53")
	if err != nil {
		return
	}
	resp, err := ws.client.Do(req)
	if err != nil {
		return "", "", err
	}

	for k, v := range resp.Header {
		header += k
		header += ": "
		for _, vx := range v {
			header += vx

		}
		header += "\n"
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}
	exp := regexp.MustCompile(`<title.*>(.*?)</title>`)
	result := exp.FindAllStringSubmatch(string(body), -1)
	for _, text := range result {
		title = text[1]
	}
	if len(title) >= 0 {
		return
	}

	return
}

func getFaviconHash(rooturl string) (string, string, error) {
	url := rooturl
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := http.Client{Transport: tr, Timeout: 3 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		return "", "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//fmt.Println("[!] websiteService.go line:197 [ not 200 or ", err, "]")
		return "", "", err
	}

	var path string
	rex := regexp.MustCompile(`<link.*rel="{0,1}.*icon"{0,1}.*href="{0,1}(.*ico|png|icon|jpg)"{0,1}/{0,1}>`)
	resslice := rex.FindAllStringSubmatch(string(body), -1)
	for _, v := range resslice {
		if len(v[1]) > 0 {
			path = v[1]
		}
	}
	if len(path) <= 0 {
		rex = regexp.MustCompile("<link.*href=\"(.*)\".*rel=\".*icon\"")
		resslice = rex.FindAllStringSubmatch(string(body), -1)
		for _, v := range resslice {
			if len(v[1]) > 0 {
				path = v[1]
			}
		}
	}

	if len(path) <= 0 {
		fUrl := url + "/favicon.ico"
		resp, err = client.Get(fUrl)
		if err != nil {
			return "", "", err
		}
		if resp.StatusCode == 200 && strings.Contains(resp.Header.Get("Content-Type"), "image") {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("[!] websiteService.go line:197 [ not 200 or ", err, "]")
				return "", "", err
			}
			b64body := helper.StandBase64(body)
			res := helper.Mmh3Hash32([]byte(b64body))
			return res, path, nil
		}

	}

	if len(path) > 0 && !strings.Contains(path, "http") {
		fUrl := url + path
		resp, err = http.Get(fUrl)
		if err != nil {
			return "", "", err
		}
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("[!] websiteService.go line:191 [", err, "]")
			return "", "", err
		}

		b64body := helper.StandBase64(body)
		res := helper.Mmh3Hash32([]byte(b64body))
		return res, fUrl, nil
	} else {
		resp, err = http.Get(path)
		if err != nil {
			return "", "", err
		}
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("[!] websiteService.go line:194 [", err, "]")
			fmt.Println(err)
		}

		b64body := helper.StandBase64(body)
		res := helper.Mmh3Hash32([]byte(b64body))
		return res, path, nil
	}

}
