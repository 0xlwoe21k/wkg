package subdomainscan

import (
	"backend/module/subdomainscan/amass/config"
	"backend/module/subdomainscan/amass/datasrcs"
	"backend/module/subdomainscan/amass/enum"
	requests "backend/module/subdomainscan/amass/requests"
	"backend/module/subdomainscan/amass/systems"
	"math/rand"
	"time"

	//systems "backend/module/subdomainscan/amass/systems"
	"github.com/caffix/stringset"

	"context"
	"fmt"
)

func PassiveDomainBrute(rootCtx context.Context, output chan *requests.Output, rootDomain string, ) {
	//rootdomain := "sf-express.com"
	//argStr := []string{"-passive","-norecursive","-noalts","-d",rootDomain}
	//amass2.RunEnumCommand(rootCtx,output,argStr)

}






func DomainBrute(rootDomain string) []requests.Output {
	var result []requests.Output
	// Setup the most basic amass configuration
	//var Resolvers = []string{"180.76.76.76","1.2.4.8","114.114.114.114","114.114.115.115","119.29.29.29","223.5.5.5","1.1.1.1:53","1.0.0.1:53","8.8.8.8:53","8.8.4.4:53", "9.9.9.9:53",
	//	"9.9.9.10:53", "77.88.8.8:53", "77.88.8.1:53", "208.67.222.222:53", "208.67.220.220:53"}
	//
	//cfg := config.NewConfig()
	//cfg.Resolvers = Resolvers
	//cfg.Passive = true
	//cfg.LocalDatabase =false
	//cfg.Recursive = false
	//cfg.Alterations = false
	//cfg.Ports = []int{443,8443}
	//
	//cfg.AddDomain(rootDomain)
	//sys, err := systems.NewLocalSystem(cfg)
	//
	//if err != nil {
	//	return nil
	//}
	//defer func() {
	//	sys.Pool().Stop()
	//	sys.Shutdown()
	//}()
	//sys.SetDataSources(datasrcs.GetAllSources(sys))
	//
	//e := enum.NewEnumeration(cfg, sys)
	//defer e.Close()
	//if e == nil {
	//	return nil
	//}
	//defer e.Close()

	rand.Seed(time.Now().UTC().UnixNano())

	// Setup the most basic amass configuration
	cfg := config.NewConfig()
	cfg.AddDomain("example.com")

	sys, err := systems.NewLocalSystem(cfg)
	if err != nil {
		return nil
	}
	sys.SetDataSources(datasrcs.GetAllSources(sys))

	e := enum.NewEnumeration(cfg, sys)
	if e == nil {
		return nil
	}
	defer e.Close()


	ctx, _ := context.WithCancel(context.Background())
	err = e.Start(ctx)
	if err != nil {
		fmt.Println("scan.go enum start error.")
		return nil
	}
	known := stringset.New()
	for _, o := range enum.ExtractOutput(ctx, e, known, true, 100) {
		result = append(result, *o)
		//fmt.Println("name:",o.Name,"    source:",o.Sources)
	}

	return result
}

func ActiveDomainBrute(rootCtx context.Context, output chan *requests.Output, rootDomain string, ) {
	//rootdomain := "sf-express.com"
	//argStr := []string{"-active","-ip","-src","-d",rootDomain}
	//amass2.RunEnumCommand(rootCtx,output,argStr)

}
