package websiteService

import (
	"backend/db"
	"backend/libs/helper"
	"backend/models"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"

	"net/http"
	"regexp"
	"strings"
	"testing"
	"time"
)

//https://user.95516.com/favicon.ico

func Test_title(t *testing.T) {
	//ws := NewWebSiteService()
	//_, title, errx := ws.getTitle("www.sf-laas.com", "https://autodiscover.immomo.com")
	//if errx != nil {
	//	fmt.Println("21:", errx)
	//}
	//
	////var IsStdCodeFlag = true
	//var tmpstr string
	//for _, v := range title {
	//	if ('\u4e00' < v && v < '\u9fa5') || ('\x00' < v && v < '\x7e') {
	//		tmpstr += string(v)
	//	}
	//}
	//fmt.Println(tmpstr)

	//if IsStdCodeFlag {
	//	fmt.Println(title)
	//}else{
	//	fmt.Println("[code error]")
	//}
}

func httpGet(url string) (result string, err error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36 Edg/95.0.1020.53")
	if err != nil {
		return
	}
	var c = http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func Test_favicon(t *testing.T) {


	url := "http://p2.sf-express.com:80"
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := http.Client{Transport: tr, Timeout: 3 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil  {
		//fmt.Println("[!] websiteService.go line:197 [ not 200 or ", err, "]")
		return
	}

	var path string
	rex := regexp.MustCompile(`<link.*rel=".*icon".*href="{0,1}(.*ico|png|icon|jpg)"{0,1}/{0,1}>`)
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
			return
		}
		if resp.StatusCode == 200 && strings.Contains(resp.Header.Get("Content-Type"), "image") {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("[!] websiteService.go line:197 [ not 200 or ", err, "]")
				return
			}
			b64body := helper.StandBase64(body)
			res := helper.Mmh3Hash32([]byte(b64body))
			fmt.Println(res)
			return
		}

	}

	if len(path) > 0 && !strings.Contains(path, "http") {
		fUrl := url + path
		resp, err = http.Get(fUrl)
		if err != nil {
			return
		}
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("[!] websiteService.go line:191 [", err, "]")
			return
		}

		b64body := helper.StandBase64(body)
		res := helper.Mmh3Hash32([]byte(b64body))
		fmt.Println(res)
		return
	} else {
		resp, err = http.Get(path)
		if err != nil {
			return
		}
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("[!] websiteService.go line:194 [", err, "]")
			fmt.Println(err)
		}

		b64body := helper.StandBase64(body)
		res := helper.Mmh3Hash32([]byte(b64body))
		fmt.Println(res)
		return
	}

}

func Test_mmh3(t *testing.T) {

	//url := "http://wcas.sf-express.com:80"

	//var title string
	//body, err := httpGet(url)
	//if err != nil {
	//	return
	//}
	//exp := regexp.MustCompile(`<title.*>(.*?)</title>`)
	//result := exp.FindAllStringSubmatch(string(body), -1)
	//for _, text := range result {
	//	title = text[1]
	//}
	//title = "不知道"

	//resp ,err :=http.Get("https://user.95516.com/favicon.ico")
	//if err != nil{
	//	fmt.Println("Get err",err)
	//	return
	//}
	//
	//body ,err := ioutil.ReadAll(resp.Body)
	//if err != nil{
	//	fmt.Println("ReadAll err",err)
	//	return
	//}
	//
	//b64body := helper.StandBase64(body)
	//foo := helper.Mmh3Hash32([]byte(b64body))
	//fmt.Println(foo)
}

func Test_Select(t *testing.T) {

	var domains []models.Domain

	////每次重新读取数据库的时候清空一次
	//ws.monitorQueue=nil
	err := db.Orm.Debug().Model(&models.Domain{}).Select("domain", "isNew").Where("isNew=true").Find(&domains).Error
	if err != nil {
		fmt.Println("[!] websiteService.go line:80 [", err, "]")
	}

	fmt.Println(domains)

}

func Test_IP(t *testing.T) {

	//先获取IP地址
	ips, err :=net.LookupIP("www.baidu.com")
	//ips, err := net.ResolveIPAddr("ip", )
	if err != nil {
		//fmt.Println("Resolution error", err.Error())
		return
	}

	fmt.Println(ips)

	return
}



func Test_RemoveRepeta(t *testing.T) {

	var wesites []string
	err := db.Orm.Debug().Model(&models.Websites{}).Select("domain").Find(&wesites).Error
	if err != nil {
		fmt.Println("[!] websiteService.go line:201 [", err, "]")
	}

	var domains []string
	err = db.Orm.Debug().Model(&models.Domain{}).Select("domain").Where("isNew=true").Find(&domains).Error
	if err != nil {
		fmt.Println("[!] websiteService.go line:207 [", err, "]")
	}

	hSection := make(map[string]string,0)
	ans := make([]string,0)
	for _,val:= range wesites{
		hSection[val]=val
	}
	for _,val := range domains{
		_,ok :=hSection[val]
		if ok{
			ans = append(ans,val)
			delete(hSection,val)
		}
	}

	fmt.Println(ans)
	//去重
}
