package domainService

import (
	"backend/db"
	"backend/global"
	"backend/libs/helper"
	"backend/libs/safe"
	"backend/models"
	subdomainscan2 "backend/module/subdomainscan"
	"context"
	"crypto/tls"
	"fmt"
	//"backend/module/subdomainscan/amass/requests"
	"backend/module/subdomainscan/amass/requests"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

type DomainMoniter struct {
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

func NewDomainMoniter() *DomainMoniter {
	ctx, cancel := context.WithCancel(context.Background())
	return &DomainMoniter{
		threadNum:          1,
		wg:                 &sync.WaitGroup{},
		newCmpFlag:         make(chan bool),
		ctx:                ctx,
		once:               &sync.Once{},
		cancel:             cancel,
		lock:               &sync.RWMutex{},
		updateCmpInfoTime:  domainMonitorTime,
		ScanDomainInfoTime: domainMonitorTime,
	}
}

var domainMonitorTime = time.Duration(24) * time.Hour

func StartDomainMonitorService() {
	rand.Seed(time.Now().UTC().UnixNano())
	dm := NewDomainMoniter()
	dm.startDomainMonitorService()

}

func ScanDomain(cmp models.Company)  {
	sd := strings.Split(cmp.Domain, "|")
	for _, vd := range sd {
		fmt.Println("[*] domain monitor:", vd)
		var domainResult []requests.Output
		domainResult = subdomainscan2.DomainBrute(vd)
		fmt.Println("[*] domain scan done.")
		if domainResult != nil {
			for _, v := range domainResult {
				//如果域名长度大于150，丢弃这个域名。
				if len(v.Name) > 150 {
					continue
				}
				//判断是否已经存在，如果已经存在就不储存
				var count int64
				err := db.Orm.Model(&models.Domain{}).Where("domain=?",v.Name).Count(&count).Error
				if err != nil {
					continue
				}
				if count <= 0 {
					//如果记录不存在就插到数据库里
					fp := true
					var newDomain = models.Domain{IsNew: &fp}
					newDomain.Domain = v.Name
					newDomain.Cid = cmp.Id
					*newDomain.IsNew = true
					if len(v.Sources) > 3 {
						newDomain.Source = v.Sources[0]

					} else {
						newDomain.Source = fmt.Sprintf("%v", v.Sources)

					}
					newDomain.UpdateTime = helper.GetCurTime()
					err := db.Orm.Model(&models.Domain{}).Create(&newDomain).Error
					if err != nil {
						fmt.Println("[!] domainService.go line:155 insert into error.   [", err, "]")
					}
				}
			}
		}
	}
}

//每5秒查看一次是否有新公司
func (d *DomainMoniter) monitorNewCompany() {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			if global.HasNewCompanyFlag {
				d.newCmpFlag <- true
				global.HasNewCompanyFlag = false
			}
		case <-d.ctx.Done():
			return

		}
	}
}

func (d *DomainMoniter) startDomainMonitorService() {
	fmt.Println("[*] Start domain monitor service.")
	//按给定时间读取数据库中的公司信息

	//监控是否有新公司被添加
	go d.monitorNewCompany()
	go d.loopUpdateCmpInfo()
	//域名监控功能
	//按设定的总时间一次性从数据库里读取信息，进行域名监控扫描。
	go d.loopMonitorDomain()
	//1.每个公司开一个协程，根据monitorstatus来确定是否需要监控

}

func (d *DomainMoniter) loopMonitorDomain() {
	//ticker := time.NewTicker(tim*time.Hour)
	ticker := time.NewTicker(d.ScanDomainInfoTime)
	for {
		select {
		case <-ticker.C:
			d.scanDomainByAmass()
		case <-d.newCmpFlag:
			d.reloadCmpInfo()
			d.onceCcanDomainByAmass()
		case <-d.ctx.Done():
			return
		}
	}
}

func (d *DomainMoniter) onceCcanDomainByAmass() {
	//log.Println("[*] start domain monitor queue.")
	if len(d.cmpInfo) <= 0 {
		return
	}
	//判断监控状态
	if !*d.cmpInfo[len(d.cmpInfo)-1].MonitorStatus {
		return
	}
	sd := strings.Split(d.cmpInfo[len(d.cmpInfo)-1].Domain, "|")
	err := db.Orm.Model(&models.Domain{}).Where("cid=?", d.cmpInfo[len(d.cmpInfo)-1].Id).Find(&d.DomainCache).Error
	if err != nil {
		fmt.Println("[!] domainService.go line:147 insert into error.   [", err, "]")
	}
	for _, vd := range sd {
		fmt.Println("[*] domain monitor:", vd)

		var domainResult []requests.Output
		var errCh = make(chan error, 1)
		safe.Go(func() error {
			domainResult = subdomainscan2.DomainBrute(vd)
			return nil
		}, errCh)
		if <-errCh != nil {
			fmt.Println("go error. ", <-errCh)
			continue
		}
		fmt.Println("[+] found new domain count:", len(domainResult))
		fp := true
		if domainResult != nil {
			for _, v := range domainResult {
				//如果域名长度大于150，丢弃这个域名。
				if len(v.Name) > 150 {
					continue
				}
				//判断是否已经存在，如果已经存在就不储存
				var isExsit = false
				for _, s := range d.DomainCache {
					if v.Name == s.Domain {
						isExsit = true
						break
					}
				}
				if !isExsit {
					//如果记录不存在就插到数据库里
					var newDomain = models.Domain{IsNew: &fp}
					newDomain.Domain = v.Name
					newDomain.Cid = d.cmpInfo[len(d.cmpInfo)-1].Id
					*newDomain.IsNew = true
					if len(v.Sources) > 3 {
						newDomain.Source = v.Sources[0]

					} else {
						newDomain.Source = fmt.Sprintf("%v", v.Sources)

					}
					newDomain.UpdateTime = helper.GetCurTime()
					d.DomainCache = append(d.DomainCache, newDomain)
					err := db.Orm.Model(&models.Domain{}).Create(&newDomain).Error
					if err != nil {
						fmt.Println("[!] domainService.go line:155 insert into error.   [", err, "]")
					}
					isExsit = false
				}
			}
		}
	}

	fmt.Println("domain cache:", len(d.DomainCache))

}

func (d *DomainMoniter) scanDomainByAmass() {
	//log.Println("[*] start domain monitor queue.")
	if len(d.cmpInfo) <= 0 {
		return
	}
	for i := len(d.cmpInfo) - 1; i > 0; i-- {
		//判断监控状态
		if !*d.cmpInfo[i].MonitorStatus {
			continue
		}
		sd := strings.Split(d.cmpInfo[i].Domain, "|")
		err := db.Orm.Model(&models.Domain{}).Where("cid=?", d.cmpInfo[i].Id).Find(&d.DomainCache).Error
		if err != nil {
			fmt.Println("[!] domainService.go line:147 insert into error.   [", err, "]")
		}
		for _, vd := range sd {
			fmt.Println("[*] domain monitor:", vd)

			var domainResult []requests.Output
			var errCh = make(chan error, 1)
			safe.Go(func() error {
				domainResult = subdomainscan2.DomainBrute(vd)
				return nil
			}, errCh)
			if <-errCh != nil {
				fmt.Println("go error. ", <-errCh)
				continue
			}
			fmt.Println("[+] found new domain count:", len(domainResult))
			fp := true
			if domainResult != nil {
				for _, v := range domainResult {
					//如果域名长度大于150，丢弃这个域名。
					if len(v.Name) > 150 {
						continue
					}
					//判断是否已经存在，如果已经存在就不储存
					var isExsit = false
					for _, s := range d.DomainCache {
						if v.Name == s.Domain {
							isExsit = true
							break
						}
					}
					if !isExsit {
						//如果记录不存在就插到数据库里
						var newDomain = models.Domain{IsNew: &fp}
						newDomain.Domain = v.Name
						newDomain.Cid = d.cmpInfo[i].Id
						*newDomain.IsNew = true
						if len(v.Sources) > 3 {
							newDomain.Source = v.Sources[0]

						} else {
							newDomain.Source = fmt.Sprintf("%v", v.Sources)

						}
						newDomain.UpdateTime = helper.GetCurTime()
						d.DomainCache = append(d.DomainCache, newDomain)
						err := db.Orm.Model(&models.Domain{}).Create(&newDomain).Error
						if err != nil {
							fmt.Println("[!] domainService.go line:155 insert into error.   [", err, "]")
						}
						isExsit = false
					}
				}
			}
		}

		fmt.Println("domain cache:", len(d.DomainCache))
	}
}

func (d *DomainMoniter) reloadCmpInfo() () {
	err := db.Orm.Model(&models.Company{}).Find(&d.cmpInfo).Error
	if err != nil {
		fmt.Println("[!] domainService.go  line:137 db error:", err)
	}
}

func (d *DomainMoniter) loopUpdateCmpInfo() {

	err := db.Orm.Model(&models.Company{}).Find(&d.cmpInfo).Error
	if err != nil {
		fmt.Println("[!] domainService.go  line:137 db error:", err)
	}

	ticker := time.NewTicker(d.updateCmpInfoTime)
	for {
		select {
		case <-ticker.C:
			//获取各公司的域名信息和待监控状态
			err := db.Orm.Model(&models.Company{}).Find(&d.cmpInfo).Error
			if err != nil {
				fmt.Println("[!] domainService.go domainService line:148 db error:", err)
			}

		}
	}

}


func (d *DomainMoniter) getIPAndTitle(domain string) (ip string, title string) {

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
