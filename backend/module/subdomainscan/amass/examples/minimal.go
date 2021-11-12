package main

import (
	"context"
	"math/rand"
	"time"

	"backend/module/subdomainscan/amass/config"
	"backend/module/subdomainscan/amass/datasrcs"
	"backend/module/subdomainscan/amass/enum"
	"backend/module/subdomainscan/amass/systems"
)

func main() {
	// Seed the default pseudo-random number generator
	rand.Seed(time.Now().UTC().UnixNano())

	// Setup the most basic amass configuration
	cfg := config.NewConfig()
	cfg.AddDomain("example.com")

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

	e.Start(context.TODO())
}
