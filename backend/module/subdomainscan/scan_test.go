package subdomainscan

import (
	amass2 "backend/module/subdomainscan/amass/cmd/amass"
	config2 "backend/module/subdomainscan/amass/config"
	datasrcs2 "backend/module/subdomainscan/amass/datasrcs"
	enum2 "backend/module/subdomainscan/amass/enum"
	systems2 "backend/module/subdomainscan/amass/systems"
	"context"
	"fmt"
	"backend/libs/stringset"
	"math/rand"
	"testing"
	"time"
)

func TestDomainBrute(t *testing.T) {

	//rootdomain := "sf-express.com"
	//argStr := []string{"-passive","-norecursive","-noalts","-d",rootdomain}
	//amass.RunEnum(context.Background(),rootdomain)


}

func TestActiveDomainBrute(t *testing.T) {
	// Seed the default pseudo-random number generator
	rand.Seed(time.Now().UTC().UnixNano())

	// Setup the most basic amass configuration
	cfg := config2.NewConfig()
	cfg.Passive = true
	cfg.AddDomain("sf-express.com")

	sys, err := systems2.NewLocalSystem(cfg)
	if err != nil {
		return
	}
	sys.SetDataSources(datasrcs2.GetAllSources(sys))

	e := enum2.NewEnumeration(cfg, sys)
	if e == nil {
		return
	}
	defer e.Close()

	ctx := context.Background()
	e.Start(ctx)
	known := stringset.New()
	for _, o := range amass2.ExtractOutput(ctx,e,known,true,0) {
		fmt.Println(o.Name)
	}
}
