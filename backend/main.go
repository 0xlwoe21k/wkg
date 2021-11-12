package main

import (
	"backend/db"
	_ "backend/routers"
	"backend/services"
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"os"
	"os/signal"
)

func main() {

	//开启几个服务.域名监控域名、漏洞扫描服务
	go services.InitService()

	beego.Run()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")

	defer func() {
		sqldb,err := db.Orm.DB()
		if err !=nil {
			fmt.Println("[!] main.go get sqldb failed. line:29  .   [",err,"]")
		}
		sqldb.Close()
	}()
}

