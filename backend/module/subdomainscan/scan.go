package subdomainscan

import (
	amass2 "backend/module/subdomainscan/amass/cmd/amass"
	config2 "backend/module/subdomainscan/amass/config"
	datasrcs2 "backend/module/subdomainscan/amass/datasrcs"
	enum2 "backend/module/subdomainscan/amass/enum"
	requests2 "backend/module/subdomainscan/amass/requests"
	systems2 "backend/module/subdomainscan/amass/systems"
	"context"
	"log"
	"math/rand"
	"time"
)

func PassiveDomainBrute(rootCtx context.Context, output chan *requests2.Output, rootDomain string, ) {
	//rootdomain := "sf-express.com"
	//argStr := []string{"-passive","-norecursive","-noalts","-d",rootDomain}
	//amass2.RunEnumCommand(rootCtx,output,argStr)

}

func DomainBrute(rootDomain string) []string {
	rand.Seed(time.Now().UTC().UnixNano())
	var result []string
	// Setup the most basic amass configuration
	cfg := config2.NewConfig()
	cfg.Dir = "/tmp"
	cfg.Passive = true
	cfg.Recursive = false
	cfg.Alterations = false
	cfg.AddDomain(rootDomain)
	sys, err := systems2.NewLocalSystem(cfg)

	if err != nil {
		return nil
	}
	sys.SetDataSources(datasrcs2.GetAllSources(sys))

	e := enum2.NewEnumeration(cfg, sys)
	if e == nil {
		return nil
	}
	defer e.Close()

	ctx, _ := context.WithCancel(context.Background())
	err = e.Start(ctx)
	if err != nil {
		log.Fatalln("scan.go enum start error.")
		return nil
	}
	//known := stringset.New()
	for _, o := range amass2.ExtractOutput(ctx, e, nil, true, 0) {
		result = append(result, o.Name)
	}

	return result
}

func ActiveDomainBrute(rootCtx context.Context, output chan *requests2.Output, rootDomain string, ) {
	//rootdomain := "sf-express.com"
	//argStr := []string{"-active","-ip","-src","-d",rootDomain}
	//amass2.RunEnumCommand(rootCtx,output,argStr)

}
