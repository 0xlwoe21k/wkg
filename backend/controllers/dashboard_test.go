package controllers

import (
	"backend/db"
	"backend/models"
	"fmt"
	"testing"
)

func Test_Count(T *testing.T)  {
	var TotalService int64
	err := db.Orm.Debug().Model(&models.Domain{}).Count(&TotalService).Error
	if err != nil {
		fmt.Println("[!] dashboard.go line:20 [", err, "]")
		return
	}
	fmt.Println(TotalService)

	var NewService int64
	err = db.Orm.Model(&models.Domain{}).Where("isNew=true").Count(&NewService).Error
	if err != nil {
		fmt.Println("[!] dashboard.go line:20 [", err, "]")
		return
	}

	fmt.Println(NewService)

}