package services

import (
	"backend/services/websiteService"
)

func InitService()  {

	//开启一个MSQ队列
	//ctx, _ := context.WithCancel(context.Background())
	//smq := go_smq.NewSmq()
	//err := smq.Register(ctx,subdomainscan.PassiveDomainBrute)
	//if err != nil{
	//	fmt.Println("[!] service.go InitService register error.:",err)
	//}
	//domainService.StartDomainMonitorService()
	websiteService.StartWebSiteMonitorService()
}

