package domainMonitorService

import (
	"backend/db"
	"backend/libs/helper"
	"backend/libs/safe"
	"backend/models"
	subdomainscan2 "backend/module/subdomainscan"
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

type domainMoniter struct {
	DomainCache           []models.Domain
	UpdateDomainCacheTime time.Duration
	threadNum             int
	cmpInfo               []models.Company
	updateCmpInfoTime     time.Duration
	ScanDomainInfoTime    time.Duration
	ctx                   context.Context
	cancel                context.CancelFunc
	newCmpFlag            chan bool
	once                  *sync.Once
	lock                  *sync.RWMutex
	wg                    *sync.WaitGroup
}

func NewDomainMoniter(updateCmpTime time.Duration) *domainMoniter {
	ctx, cancel := context.WithCancel(context.Background())
	return &domainMoniter{
		threadNum:             1,
		wg:                    &sync.WaitGroup{},
		newCmpFlag:            make(chan bool),
		ctx:                   ctx,
		once:                  &sync.Once{},
		cancel:                cancel,
		lock:                  &sync.RWMutex{},
		updateCmpInfoTime:     updateCmpTime * time.Second,
		ScanDomainInfoTime:    ScanDomainTime * time.Second,
		UpdateDomainCacheTime: UpdateDomainCacheTime * time.Second,
	}
}

var updateCmpTime = time.Duration(20)
var ScanDomainTime = time.Duration(20)
var UpdateDomainCacheTime = time.Duration(60)

func StartDomainMonitorService() {
	dm := NewDomainMoniter(updateCmpTime)
	dm.startDomainMonitorService()

}

func (d *domainMoniter) startDomainMonitorService() {
	fmt.Println("[*] Start domain monitor service.")
	//按给定时间读取数据库中的公司信息
	d.initDomainCache()
	//time.Sleep(3 * time.Second)
	//go d.loopUpdateDomainCache()
	go d.loopUpdateCmpInfo()
	//域名监控功能
	//按设定的总时间一次性从数据库里读取信息，进行域名监控扫描。
	go d.loopMonitorDomain()
	//1.每个公司开一个协程，根据monitorstatus来确定是否需要监控

}
func (d *domainMoniter) loopUpdateDomainCache() {
	ticker := time.NewTicker(d.updateCmpInfoTime)
	for {
		select {
		case <-ticker.C:
			if len(d.DomainCache) < 1000 {
				err := db.Orm.Debug().Model(&models.Domain{}).Updates(&d.DomainCache).Error
				if err != nil {
					log.Fatalln("[!] domainMonitorService.go save domain cache failed. line:75.  [", err, "]")
					return
				}
			} else {
				for i := 0; i < len(d.DomainCache)/1000; i += 1000 {
					tmpDomain := d.DomainCache[i : i*1000]
					err := db.Orm.Model(&models.Domain{}).Updates(&tmpDomain).Error
					if err != nil {
						log.Fatalln("[!] domainMonitorService.go save domain cache failed. line:91.  [", err, "]")
						return
					}
				}
			}
		default:
		}
	}
}

func (d *domainMoniter) initDomainCache() {
	err := db.Orm.Model(&models.Domain{}).Find(&d.DomainCache).Error
	if err != nil {
		log.Fatalln("[!] domainMonitorService.go init domain cache failed. line:83,   [", err, "]")
		return
	}
}

func (d *domainMoniter) loopMonitorDomain() {
	rctx, _ := context.WithCancel(d.ctx)
	//ticker := time.NewTicker(tim*time.Hour)
	ticker := time.NewTicker(d.ScanDomainInfoTime)
	for {
		select {
		case <-ticker.C:
			d.scanDomainByAmass()
		case <-rctx.Done():
			return
		default:

		}
	}
}

func (d *domainMoniter) scanDomainByAmass() {
	//log.Println("[*] start domain monitor queue.")
	if len(d.cmpInfo) <= 0 {
		return
	}
	for _, vcmp := range d.cmpInfo {
		//不同操作系统换行不同，windows \r\n  unix   \n    mac \r所以要做区别
		sd := strings.Split(vcmp.Domain, "|")
		for _, vd := range sd {
			fmt.Println("[*] domain monitor:", vd)

			var domainResult []string
			var errCh = make(chan error,1)
			safe.Go(func() error {
				domainResult = subdomainscan2.DomainBrute(vd)
				return nil
			},errCh)
			if <-errCh != nil{
				fmt.Println("go error. ",<-errCh)
				continue
			}
			fmt.Println("new domain count:", len(domainResult))
			fp := true
			if domainResult != nil {
     				for _, v := range domainResult {
					//如果域名长度大于150，丢弃这个域名。
					if len(v) > 150 {
						continue
					}
					if len(d.DomainCache) == 0 {
						var newDomain = models.Domain{IsNew: &fp}
						newDomain.Domain = v
						newDomain.Cid = vcmp.Id
						*newDomain.IsNew = true
						newDomain.UpdateTime = helper.GetCurTime()
						d.DomainCache = append(d.DomainCache, newDomain)
						err := db.Orm.Debug().Model(&models.Domain{}).Create(&newDomain).Error
						if err != nil {
							log.Fatalln("[!] domainMonitorService.go line:147 insert into error.   [", err, "]")
						}
						continue
					}
					//判断是否已经存在，如果已经存在就不储存
					var isExsit = false
					for _, s := range d.DomainCache {
						if v == s.Domain {
							isExsit = true
							break
						}
					}
					if !isExsit {
						//如果记录不存在就插到数据库里
						var newDomain = models.Domain{IsNew: &fp}
						newDomain.Domain = v
						newDomain.Cid = vcmp.Id
						*newDomain.IsNew = true
						newDomain.UpdateTime = helper.GetCurTime()
						d.DomainCache = append(d.DomainCache, newDomain)
						err := db.Orm.Model(&models.Domain{}).Create(&newDomain).Error
						if err != nil {
							log.Fatalln("[!] domainMonitorService.go line:155 insert into error.   [", err, "]")
						}
						isExsit = false
					}
				}
			}
		}
		fmt.Println("domain cache:", len(d.DomainCache))
	}
}

func (d *domainMoniter) loopUpdateCmpInfo() {

	err := db.Orm.Model(&models.Company{}).Find(&d.cmpInfo).Error
	if err != nil {
		log.Fatalln("[!] domainMonitorService.go  line:137 db error:", err)
	}
	ticker := time.NewTicker(d.updateCmpInfoTime)
	for {
		select {
		case <-ticker.C:
			//var cmpInfo = []models.Company{}
			//获取各公司的域名信息和待监控状态
			err := db.Orm.Model(&models.Company{}).Find(&d.cmpInfo).Error
			if err != nil {
				log.Fatalln("[!] domainMonitorService.go domainMonitorService line:148 db error:", err)
				continue
			}
		default:

		}
	}

}

//获取Domain的其它信息,如果标题,IP等信息
func (d *domainMoniter) getDomainDetailInfo() {
	log.Println("[+] start update dommain detail information.")
	ticker := time.NewTicker(24 * time.Hour)
	for {
		select {
		case <-ticker.C:
			var domains []models.Domain
			err := db.Orm.Debug().Model(&models.Domain{}).Find(&domains).Error
			if err != nil {
				log.Fatalln("[!] get domains info error.:", err)
				break
			}
			if len(domains) <= 0 {
				break
			}
			for i := 0; i < len(domains); i++ {
				go func(dom models.Domain) {
					ip, title := d.getIPAndTitle(dom.Domain)
					if strings.Contains(dom.Ip, ip) || strings.Contains(dom.Title, title) {
						dom.Ip = ip
						dom.Title = title
						err = db.Orm.Debug().Model(&models.Domain{}).Where("id=?", domains[i].Id).Updates(&domains[i]).Error
						if err != nil {
							//log.Fatalln()
						}
					}
				}(domains[i])
			}
			log.Println("update domain information done.")
		}
	}

}

func (d *domainMoniter) getIPAndTitle(domain string) (ip string, title string) {

	//先获取IP地址
	ips, err := net.ResolveIPAddr("ip", domain)
	if err != nil {
		//fmt.Println("Resolution error", err.Error())
		return
	}

	ip = ips.String()
	if len(ips.String()) <= 0 {
		return
	}

	domain = strings.ReplaceAll(domain, "\r", "")
	domain = strings.ReplaceAll(domain, "\n", "")
	v1 := "http://" + domain
	body, err := httpGet(v1)
	if err != nil {
		return
	}
	exp := regexp.MustCompile(`<title>(.*?)</title>`)
	result := exp.FindAllStringSubmatch(string(body), -1)
	for _, text := range result {
		title = text[1]
	}
	if len(title) >= 0 {
		return
	}

	//https请求
	v1 = "https://" + domain
	body, err = httpGet(v1)
	if err != nil {
		return
	}
	exp = regexp.MustCompile(`<title>(.*?)</title>`)
	result = exp.FindAllStringSubmatch(string(body), -1)
	for _, text := range result {
		title = text[1]
	}

	return
}

func httpGet(url string) (result string, err error) {
	tr := &http.Transport{
		//Proxy:               p,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		TLSHandshakeTimeout:   8 * time.Second,
		ResponseHeaderTimeout: 4 * time.Second,
		DisableKeepAlives:     false,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	client := http.Client{Transport: tr, Timeout: 2 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
