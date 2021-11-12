package controllers

import (
	"backend/db"
	"backend/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type DomainController struct {
	beego.Controller
}

func (c *DomainController)GetDomainInfo()  {
	var page = &PageParam{}
	body := c.Ctx.Input.RequestBody
	err := json.Unmarshal(body,page)
	if err != nil{
		res,_ := json.Marshal(Result{Code: 400,Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}
	fmt.Println("page:",page.Page,"  pagesize:",page.PageSize)


	var count int64
	//查询总数
	err = db.Orm.Model(&models.Domain{}).Count(&count).Error
	if err != nil {
		res,_ := json.Marshal(Result{Code: 400,Msg: "query error."})
		c.Ctx.WriteString(string(res))
		return
	}
	//根据page和pagesize查询数据
	var dom []models.Domain
	err = db.Orm.Model(&models.Domain{}).Limit(page.PageSize).Offset((page.Page - 1) * page.PageSize).Find(&dom).Error
	if err != nil {
		res,_ := json.Marshal(Result{Code: 400,Msg: "query error."})
		c.Ctx.WriteString(string(res))
		return
	}

	data ,err := json.Marshal(dom)
	if err != nil{
		res,_ := json.Marshal(Result{Code: 400,Msg: "json marshal error.."})
		c.Ctx.WriteString(string(res))
	}
	type Res struct {
		Code int `json:"code"`
		Msg string `json:"msg"`
		Total int `json:"total"`
	}
	res,_ := json.Marshal(Res{Code: 200,Msg: string(data),Total: int(count)})

	c.Ctx.WriteString(string(res))
}

func (c *DomainController)GetDomainInfoByKey()  {
	var err error
	var dss  = &DomaSearchStrut{}
	body := c.Ctx.Input.RequestBody
	err = json.Unmarshal(body,dss)
	if err != nil{
		res,_ := json.Marshal(Result{Code: 400,Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}
	fmt.Println("type:",dss.Type,"  keyword:",dss.KeyWord)

	var count int64
	var dom = &[]models.Domain{}
	if dss.Type == "ip"{
		err = db.Orm.Debug().Model(&models.Domain{}).Where("ip LIKE ?", "%"+dss.KeyWord+"%").Find(&dom).Count(&count).Error
	}else if dss.Type == "domain"{
		err = db.Orm.Debug().Model(&models.Domain{}).Where("domain LIKE ?", "%"+dss.KeyWord+"%").Count(&count).Find(&dom).Error
	}else if dss.Type == "title"{
		err = db.Orm.Debug().Model(&models.Domain{}).Where("title LIKE ?", "%"+dss.KeyWord+"%").Count(&count).Find(&dom).Error
	}
	if err != nil {
		res,_ := json.Marshal(Result{Code: 400,Msg: "query error."})
		c.Ctx.WriteString(string(res))
		return
	}

	data ,err := json.Marshal(dom)
	if err != nil{
		res,_ := json.Marshal(Result{Code: 400,Msg: "json marshal error.."})
		c.Ctx.WriteString(string(res))
	}
	type Res struct {
		Code int `json:"code"`
		Msg string `json:"msg"`
		Total int `json:"total"`
	}
	res,_ := json.Marshal(Res{Code: 200,Msg: string(data),Total: int(count)})

	c.Ctx.WriteString(string(res))
}