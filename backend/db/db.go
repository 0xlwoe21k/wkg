package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var(
	err error
	Orm *gorm.DB
)

func init()  {
	dsn := "root:root@(localhost)/wkg?charset=utf8&parseTime=true&loc=Local"
	Orm,err = gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println("[!] db.go line:18 error:"+err.Error())
		os.Exit(0)
	}
}