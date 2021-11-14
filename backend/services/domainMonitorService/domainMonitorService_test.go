package domainMonitorService

import (
	"backend/db"
	"backend/models"
	"log"
	"testing"
	"time"
)

func TestDomainMoniter_UpdateDomainInfo(t *testing.T) {
	var DomainCache  []models.Domain

	err := db.Orm.Model(&models.Domain{}).Find(&DomainCache).Error
	if err != nil {
		log.Fatalln("[!] domainMonitorService.go init domain cache failed. line:83,   [", err, "]")
		return
	}
	DomainCache[0].Ip= "123"

	err = db.Orm.Debug().Model(&DomainCache).Where("1=1").Save(&DomainCache).Error
	if err != nil {
		log.Fatalln("[!] domainMonitorService.go save domain cache failed. line:75.  [", err, "]")
		return
	}
}


func addrtram(str *string){
	*str = "abc"
}

func TestDetDomaininfo(t *testing.T) {


	//var s = "123"
	//addrtram(&s)
	//fmt.Println(s)
	//
	//s= "123"
	//go addrtram(&s)
	//time.Sleep(1*time.Second)
	//fmt.Println(s)

	//var dom []models.Domain
	//
	//for i := 0; i < 1000; i ++ {
	//	var tmp models.Domain
	//	tmp.Domain = fmt.Sprintf("%d ",i)
	//	dom = append(dom,tmp)
	//}
	//for i := 0; i < 1000; i ++ {
	//	go func(d *models.Domain,i int ) {
	//		d.Domain += fmt.Sprintf("%d ",i)
	//	}(&dom[i],i)
	//}
	//
	//
	//time.Sleep(1*time.Second)
	//for i := 0; i < 1000; i ++ {
	//	fmt.Println(dom[i].Domain)
	//}

	dm := NewDomainMoniter(time.Duration(updateCmpTime))
	dm.getDomainDetailInfo()
	//
	//ip,title := dm.getIPAndTitle("sfsrc.sf-express.com")
	//fmt.Println(ip,"    ",title)
}