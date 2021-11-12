package controllers

import (
	"backend/db"
	"backend/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type CompanyController struct {
	beego.Controller
}


func (c *CompanyController)GetCompanyInfo()  {
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
	err = db.Orm.Model(&models.Company{}).Count(&count).Error
	if err != nil {
		res,_ := json.Marshal(Result{Code: 400,Msg: "query error."})
		c.Ctx.WriteString(string(res))
		return
	}
	//根据page和pagesize查询数据
	var cmp []models.Company
	err = db.Orm.Model(&models.Company{}).Limit(page.PageSize).Offset((page.Page - 1) * page.PageSize).Find(&cmp).Error
	if err != nil {
		res,_ := json.Marshal(Result{Code: 400,Msg: "query error."})
		c.Ctx.WriteString(string(res))
		return
	}

	data ,err := json.Marshal(cmp)
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


//{"projectType":"src","companyName":"asd","domain":"asd","srcUrl":{"_shallow":false,"dep":{"w":0,"n":0},"__v_isRef":true,"_rawValue":"asd","_value":"asd"},"keyWord":"asd","monitorStatus":"1","monitorRate":"30","vulnScanStatus":"1","vulnScanRate":"30"}

func (c *CompanyController)NewCompanyInfo()  {
	var tc = &models.Company{}
	body := c.Ctx.Input.RequestBody
	err := json.Unmarshal(body,tc)
	if err != nil{
		fmt.Println(err)
		res,_ := json.Marshal(Result{Code: 400,Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}

	tc.LastUpdateTime = time.Unix(time.Now().Unix(),0).Format("2006-01-02 15:04:05")

	err = db.Orm.Debug().Model(&models.Company{}).Create(&tc).Error

	var count int64
	db.Orm.Where("username = ?","gelen").Count(&count)

	if err !=nil{
		res,_ := json.Marshal(Result{Code: 400,Msg: "update data error."})
		c.Ctx.WriteString(string(res))
		return
	}else{
		res,_ := json.Marshal(Result{Code: 200,Msg: "update success."})
		c.Ctx.WriteString(string(res))
		return
	}

}

func (c *CompanyController)GetCompanyByid()  {
	type TId struct {
		Id string `json:"id"`
	}
	var Tid = &TId{}
	body := c.Ctx.Input.RequestBody
	err := json.Unmarshal(body,Tid)
	if err != nil{
		fmt.Println(err)
		res,_ := json.Marshal(Result{Code: 400,Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}

	var OneCmp models.Company
	err = db.Orm.Debug().Model(&models.Company{}).Find(&OneCmp,"id=?",Tid.Id).Error
	if err !=nil{
		res,_ := json.Marshal(Result{Code: 400,Msg: "update data error."})
		c.Ctx.WriteString(string(res))
		return
	}else{
		msg ,err:= json.Marshal(OneCmp)
		if err != nil{
			res,_ := json.Marshal(Result{Code: 400,Msg: "Marshal onecmp error."})
			c.Ctx.WriteString(string(res))
			return
		}
		res,_ := json.Marshal(Result{Code: 200,Msg: string(msg)})
		c.Ctx.WriteString(string(res))
		return
	}

}


func (c *CompanyController)UpdateCompanyInfo()  {
	var tc = &models.Company{}
	body := c.Ctx.Input.RequestBody
	err := json.Unmarshal(body,tc)
	if err != nil{
		fmt.Println(err)
		res,_ := json.Marshal(Result{Code: 400,Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}

	tc.LastUpdateTime = time.Unix(time.Now().Unix(),0).Format("2006-01-02 15:04:05")

	//输入的时候就把回车换

	//data := map[string]
	//fmt.Println(helper.Struct2Map(*tc))

	err = db.Orm.Debug().Model(&models.Company{}).Where("id=?",tc.Id).Updates(&tc).Error
	if err !=nil{
		res,_ := json.Marshal(Result{Code: 400,Msg: "update data error."})
		c.Ctx.WriteString(string(res))
		return
	}else{
		res,_ := json.Marshal(Result{Code: 200,Msg: "update success."})
		c.Ctx.WriteString(string(res))
		return
	}

}