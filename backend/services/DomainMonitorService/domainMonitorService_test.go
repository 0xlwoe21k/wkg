package DomainMonitorService

import (
	subdomainscan2 "backend/module/subdomainscan"
	"fmt"
	"testing"
	"time"
)

func TestDomainMoniter_UpdateCmpInfo(t *testing.T) {


	var domain = "sxjdfreight.com"
	result := subdomainscan2.DomainBrute(domain)
	fmt.Println(result)


	//domain  = []string{"sf-express.com"}
	//result = subdomainscan2.DomainBrute(domain)
	//fmt.Println(result)
	//ticker := time.NewTicker(5*time.Second)
	//for {
	//	select {
	//	case  <-ticker.C:
	//		fmt.Println("123")
	//	}
	//}
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