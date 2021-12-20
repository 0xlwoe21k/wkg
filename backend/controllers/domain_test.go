package controllers

import (
	"backend/db"
	"backend/models"
	"fmt"
	"testing"
)

func TestDomainController_DelDomainInfoById(t *testing.T) {
	err := db.Orm.Model(&models.Domain{}).Where("id=?",11).Update("isNew",false).Error
	if err != nil {
		fmt.Println("error",err)
		return
	}
}

func TestDomainController_ReadAllFlagDomainInfoById(t *testing.T) {
	err := db.Orm.Model(&models.Domain{}).Where("1=1").Update("isNew",true).Error
	if err != nil {
		fmt.Println("12313123")
	}
}