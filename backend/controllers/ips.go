package controllers

import (
	"backend/db"
	"backend/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type IPsController struct {
	beego.Controller
}



func (c *IPsController)GetIPsInfo()  {
	var param = &PageParam{}
	body := c.Ctx.Input.RequestBody
	err := json.Unmarshal(body,param)
	if err != nil{
		res,_ := json.Marshal(Result{Code: 400,Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}
	//fmt.Println("page:",page.Page,"  pagesize:",page.PageSize)


	var count int64
	//查询总数
	err = db.Orm.Model(&models.Ips{}).Count(&count).Error
	if err != nil {
		res,_ := json.Marshal(Result{Code: 400,Msg: "query error."})
		c.Ctx.WriteString(string(res))
		return
	}
	//根据page和pagesize查询数据
	var IPs []models.Ips
	err = db.Orm.Model(&models.Ips{}).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Find(&IPs).Error
	if err != nil {
		res,_ := json.Marshal(Result{Code: 400,Msg: "query error."})
		c.Ctx.WriteString(string(res))
		return
	}

	data ,err := json.Marshal(IPs)
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


func (c *IPsController)GetNewIPsInfo()  {
	var page = &PageParam{}
	body := c.Ctx.Input.RequestBody
	err := json.Unmarshal(body,page)
	if err != nil{
		res,_ := json.Marshal(Result{Code: 400,Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}
	//fmt.Println("page:",page.Page,"  pagesize:",page.PageSize)


	var count int64
	//查询总数
	err = db.Orm.Model(&models.Ips{}).Where("isNew=true").Count(&count).Error
	if err != nil {
		res,_ := json.Marshal(Result{Code: 400,Msg: "query error."})
		c.Ctx.WriteString(string(res))
		return
	}
	//根据page和pagesize查询数据
	var dom []models.Domain
	err = db.Orm.Model(&models.Domain{}).Where("isNew=true").Limit(page.PageSize).Offset((page.Page - 1) * page.PageSize).Find(&dom).Error
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


func (c *IPsController) ReadAllFlagIPsInfoById(){
	err := db.Orm.Model(&models.Domain{}).Where("1=1").Update("isNew",false).Error
	if err != nil {
		res,_ := json.Marshal(Result{Code: 400,Msg: "modify error."})
		c.Ctx.WriteString(string(res))
		return
	}
	res,_ := json.Marshal(Result{Code: 200,Msg: "modify success."})
	c.Ctx.WriteString(string(res))
}

func (c *IPsController)ReadFlagIPsInfoById()  {
	type DelId struct {
		Id  string `json:"id"`
	}
	var ei = &DelId{}
	body := c.Ctx.Input.RequestBody
	err := json.Unmarshal(body,ei)
	if err != nil{
		res,_ := json.Marshal(Result{Code: 400,Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}
	id := strings.Split(ei.Id,",")

	for _,v := range id {
		err := db.Orm.Model(&models.Domain{}).Where("id=?",v).Update("isNew",false).Error
		if err != nil {
			res,_ := json.Marshal(Result{Code: 400,Msg: "modify error."})
			c.Ctx.WriteString(string(res))
			return
		}
	}
	type Res struct {
		Code int `json:"code"`
		Msg string `json:"msg"`
		Total int `json:"total"`
	}
	res,_ := json.Marshal(Result{Code: 200,Msg: "modify success."})
	c.Ctx.WriteString(string(res))
}

func (c *IPsController)GetIPsInfoByKey()  {
	var err error
	var dss  = &SearchStrut{}
	body := c.Ctx.Input.RequestBody
	err = json.Unmarshal(body,dss)
	if err != nil{
		res,_ := json.Marshal(Result{Code: 400,Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}
	fmt.Println("type:",dss.Type,"  keyword:",dss.KeyWord)

	var dom = &[]models.Domain{}
	if dss.Type == "ip"{
		err = db.Orm.Debug().Model(&models.Domain{}).Where("ip LIKE ?", "%"+dss.KeyWord+"%").Find(&dom).Error
	}else if dss.Type == "domain"{
		err = db.Orm.Debug().Model(&models.Domain{}).Where("domain LIKE ?", "%"+dss.KeyWord+"%").Find(&dom).Error
	}else if dss.Type == "title"{
		err = db.Orm.Debug().Model(&models.Domain{}).Where("title LIKE ?", "%"+dss.KeyWord+"%").Find(&dom).Error
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
	res,_ := json.Marshal(Res{Code: 200,Msg: string(data),Total: 1})

	c.Ctx.WriteString(string(res))
}