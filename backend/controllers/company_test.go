package controllers

import (
	"backend/db"
	"backend/models"
	"fmt"
	"testing"
)

func TestCompanyController_DelCompanyByid(t *testing.T) {
	var OneCmp models.Company
	err := db.Orm.Debug().Delete(&OneCmp, 8).Error
	if err != nil {
		fmt.Println(err)
	}

}
