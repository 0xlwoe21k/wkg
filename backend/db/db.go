package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

var(
	err error
	Orm *gorm.DB
)

func init()  {
	dsn := "root:root@(localhost)/wkg?charset=utf8&parseTime=true&loc=Local&charset=utf8"
	Orm,err = gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println("[!] db.go line:18 error:"+err.Error())
		os.Exit(0)
	}

	sqldb ,err :=Orm.DB()
	if err != nil{
		fmt.Println("[!] db.go line:27 [",err,"]")
		panic(err)
	}
	sqldb.SetConnMaxLifetime(time.Hour*4)
}