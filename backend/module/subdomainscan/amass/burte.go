package main

import (
	"backend/module/subdomainscan/amass/config"
	"backend/module/subdomainscan/amass/datasrcs"
	"backend/module/subdomainscan/amass/enum"
	"backend/module/subdomainscan/amass/systems"
	"context"
	"fmt"
	"github.com/caffix/stringset"
	"math/rand"
	"time"
)

func main() {
	// Seed the default pseudo-random number generator
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println(123123)
	//Setup the most basic amass configuration
	cfg := config.NewConfig()
	cfg.AddDomain("sf-express.com")

	sys, err := systems.NewLocalSystem(cfg)
	if err != nil {
		return
	}
	sys.SetDataSources(datasrcs.GetAllSources(sys))

	e := enum.NewEnumeration(cfg, sys)
	if e == nil {
		return
	}
	defer e.Close()

	ctx := context.Background()
	e.Start(ctx)

	known := stringset.New()
	for _, o := range enum.ExtractOutput(ctx,e,known,true,100) {
		fmt.Println(o.Name)
	}
}