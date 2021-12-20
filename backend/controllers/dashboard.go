package controllers

import (
	"backend/db"
	"backend/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type DashBoard struct {
	beego.Controller
}



func (c *DashBoard) GetDashboardInfo() {
	var TotalDomain int64
	err := db.Orm.Model(&models.Domain{}).Count(&TotalDomain).Error
	if err != nil {
		fmt.Println("[!] dashboard.go line:20 [", err, "]")
		res,_ := json.Marshal(Result{Code: 400,Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}

	var NewDomain int64
	err = db.Orm.Model(&models.Domain{}).Where("isNew=true").Count(&NewDomain).Error
	if err != nil {
		fmt.Println("[!] dashboard.go line:20 [", err, "]")
		res,_ := json.Marshal(Result{Code: 400,Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}

	var TotalWebwite int64
	err = db.Orm.Model(&models.Websites{}).Count(&TotalWebwite).Error
	if err != nil {
		fmt.Println("[!] dashboard.go line:20 [", err, "]")
		res,_ := json.Marshal(Result{Code: 400,Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}

	var NewWebsite int64
	err = db.Orm.Model(&models.Websites{}).Where("isNew=true").Count(&NewWebsite).Error
	if err != nil {
		fmt.Println("[!] dashboard.go line:20 [", err, "]")
		res,_ := json.Marshal(Result{Code: 400,Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}

	var TotalIp int64
	err = db.Orm.Model(&models.Ips{}).Count(&TotalIp).Error
	if err != nil {
		fmt.Println("[!] dashboard.go line:20 [", err, "]")
		res,_ := json.Marshal(Result{Code: 400,Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}

	var NewIp int64
	err = db.Orm.Model(&models.Ips{}).Where("isNew=true").Count(&NewIp).Error
	if err != nil {
		fmt.Println("[!] dashboard.go line:20 [", err, "]")
		res,_ := json.Marshal(Result{Code: 400,Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}

	var TotalService int64
	err = db.Orm.Model(&models.Services{}).Count(&TotalService).Error
	if err != nil {
		fmt.Println("[!] dashboard.go line:20 [", err, "]")
		res,_ := json.Marshal(Result{Code: 400,Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}

	var NewService int64
	err = db.Orm.Model(&models.Services{}).Where("isNew=true").Count(&NewService).Error
	if err != nil {
		fmt.Println("[!] dashboard.go line:20 [", err, "]")
		res,_ := json.Marshal(Result{Code: 400,Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}

	res,err := json.Marshal(struct {
		TotalDomain int64		`json:"totalTomain"`
		NewDomain int64			`json:"newDomain"`
		TotalWebsite int64		`json:"totalWebsite"`
		NewWebsite	int64		`json:"newWebsite"`
		TotalIps	int64		`json:"totalIps"`
		NewIps		int64		`json:"newIps"`
		TotalService	int64	`json:"totalService"`
		NewService		int64	`json:"newService"`
	}{
		TotalDomain:TotalDomain,
		NewDomain: NewDomain,
		TotalWebsite: NewWebsite,
		NewWebsite: NewWebsite,
		TotalIps: TotalIp,
		NewIps: NewIp,
		TotalService: TotalService,
		NewService:NewService,
	})
	if err != nil{
		res,_ := json.Marshal(Result{Code: 400,Msg: "json marshal error.."})
		c.Ctx.WriteString(string(res))
	}
	data,_ := json.Marshal(Result{Code: 200,Msg: string(res)})
	c.Ctx.WriteString(string(data))

}