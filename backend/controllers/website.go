package controllers

import (
	"backend/db"
	"backend/http/request"
	"backend/models"
	"backend/services/websiteService"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)


type WebSiteController struct {
	beego.Controller
}

func (c *WebSiteController)GetWebSiteInfo()  {
	var param = &PageParam{}
	body := c.Ctx.Input.RequestBody
	err := json.Unmarshal(body,param)
	if err != nil{
		res,_ := json.Marshal(Result{Code: 400,Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}

	var count int64	//查询总数
	if param.Keyword == ""{
		err = db.Orm.Debug().Model(&models.Websites{}).Count(&count).Error
	}else if param.Type == "title"{
		err = db.Orm.Debug().Model(&models.Websites{}).Where("title LIKE ?", "%"+param.Keyword+"%").Count(&count).Error
	}else if param.Type == "website"{
		err = db.Orm.Debug().Model(&models.Websites{}).Where("website LIKE ?", "%"+param.Keyword+"%").Count(&count).Error
	}else if param.Type == "favicon"{
		err = db.Orm.Debug().Model(&models.Websites{}).Where("favicon LIKE ?", "%"+param.Keyword+"%").Count(&count).Error
	}else if param.Type == "ips"{
		err = db.Orm.Debug().Model(&models.Websites{}).Where("ips LIKE ?", "%"+param.Keyword+"%").Count(&count).Error
	}
	if err != nil{
		res,_ := json.Marshal(Result{Code: 400,Msg: "database error."})
		c.Ctx.WriteString(string(res))
	}

	var website = &[]models.Websites{}
	if param.Keyword == ""{
		err = db.Orm.Debug().Model(&models.Websites{}).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Find(&website).Error
	}else if param.Type == "title"{
		err = db.Orm.Debug().Model(&models.Websites{}).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Where("title LIKE ?", "%"+param.Keyword+"%").Find(&website).Error
	}else if param.Type == "website"{
		err = db.Orm.Debug().Model(&models.Websites{}).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Where("website LIKE ?", "%"+param.Keyword+"%").Find(&website).Error
	}else if param.Type == "favicon"{
		err = db.Orm.Debug().Model(&models.Websites{}).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Where("favicon LIKE ?", "%"+param.Keyword+"%").Find(&website).Error
	}else if param.Type == "ips"{
		err = db.Orm.Debug().Model(&models.Websites{}).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Where("ips LIKE ?", "%"+param.Keyword+"%").Find(&website).Error
	}
	if err != nil{
		res,_ := json.Marshal(Result{Code: 400,Msg: "database error."})
		c.Ctx.WriteString(string(res))
	}

	data ,err := json.Marshal(website)
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


func (c *WebSiteController)ScanNew(){
	wsService := websiteService.NewWebSiteService()
	err := wsService.ScanWebsiteInfo()
	if err != nil{
		res,_ := json.Marshal(Result{Code: 400,Msg: err.Error()})
		c.Ctx.WriteString(string(res))
		return
	}else{
		res,_ := json.Marshal(Result{Code: 200,Msg: "scanning"})
		c.Ctx.WriteString(string(res))
		return
	}

}

func (c *WebSiteController)GetWebSiteInfoByCid()  {
	var param = &request.Query{}
	body := c.Ctx.Input.RequestBody
	err := json.Unmarshal(body,param)
	if err != nil{
		res,_ := json.Marshal(Result{Code: 400,Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}
	var count int64	//查询总数
	if param.Keyword == ""{
		err = db.Orm.Debug().Model(&models.Websites{}).Where("cid=?",param.Cid).Count(&count).Error
	}else if param.Type == "title"{
		err = db.Orm.Debug().Model(&models.Websites{}).Where("title LIKE ? and cid=?", "%"+param.Keyword+"%",param.Cid).Count(&count).Error
	}else if param.Type == "website"{
		err = db.Orm.Debug().Model(&models.Websites{}).Where("website LIKE ? and cid=?", "%"+param.Keyword+"%",param.Cid).Count(&count).Error
	}else if param.Type == "favicon"{
		err = db.Orm.Debug().Model(&models.Websites{}).Where("favicon LIKE ? and cid=?", "%"+param.Keyword+"%",param.Cid).Count(&count).Error
	}else if param.Type == "ips"{
		err = db.Orm.Debug().Model(&models.Websites{}).Where("ips LIKE ? and cid=?", "%"+param.Keyword+"%",param.Cid).Count(&count).Error
	}
	if err != nil{
		res,_ := json.Marshal(Result{Code: 400,Msg: "database error."})
		c.Ctx.WriteString(string(res))
		return
	}

	var website = &[]models.Websites{}
	if param.Keyword == ""{
		err = db.Orm.Debug().Model(&models.Websites{}).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Where("cid=?",param.Cid).Find(&website).Error
	}else if param.Type == "title"{
		err = db.Orm.Debug().Model(&models.Websites{}).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Where("title LIKE ? and cid=?", "%"+param.Keyword+"%",param.Cid).Find(&website).Error
	}else if param.Type == "website"{
		err = db.Orm.Debug().Model(&models.Websites{}).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Where("website LIKE ? and cid=?", "%"+param.Keyword+"%",param.Cid).Find(&website).Error
	}else if param.Type == "favicon"{
		err = db.Orm.Debug().Model(&models.Websites{}).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Where("favicon LIKE ? and cid=?", "%"+param.Keyword+"%",param.Cid).Find(&website).Error
	}else if param.Type == "ips"{
		err = db.Orm.Debug().Model(&models.Websites{}).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Where("ips LIKE ? and cid=?", "%"+param.Keyword+"%",param.Cid).Find(&website).Error
	}
	if err != nil{
		res,_ := json.Marshal(Result{Code: 400,Msg: "database error."})
		c.Ctx.WriteString(string(res))
		return
	}

	data ,err := json.Marshal(website)
	if err != nil{
		res,_ := json.Marshal(Result{Code: 400,Msg: "json marshal error.."})
		c.Ctx.WriteString(string(res))
		return
	}
	type Res struct {
		Code int `json:"code"`
		Msg string `json:"msg"`
		Total int `json:"total"`
	}
	res,_ := json.Marshal(Res{Code: 200,Msg: string(data),Total: int(count)})

	c.Ctx.WriteString(string(res))
}

func (c *WebSiteController)GetNewWebSiteInfo()  {
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
	err = db.Orm.Model(&models.Websites{}).Where("isNew=true").Count(&count).Error
	if err != nil {
		res,_ := json.Marshal(Result{Code: 400,Msg: "query error."})
		c.Ctx.WriteString(string(res))
		return
	}
	//根据page和pagesize查询数据
	var dom []models.Websites
	err = db.Orm.Model(&models.Websites{}).Where("isNew=true").Limit(page.PageSize).Offset((page.Page - 1) * page.PageSize).Find(&dom).Error
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

func (c *WebSiteController) ScanWebsiteInfo()  {
	type Res struct {
		Code int `json:"code"`
		Msg string `json:"msg"`
	}

	ws := websiteService.NewWebSiteService()
	go ws.ScanWebsiteInfo()

	res,_ := json.Marshal(Res{Code: 200,Msg: "scanning"})

	c.Ctx.WriteString(string(res))
}


func (c *WebSiteController)GetWebSiteInfoByKey()  {
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


	var website = &[]models.Websites{}
	if dss.Type == "title"{
		err = db.Orm.Debug().Model(&models.Websites{}).Where("title LIKE ?", "%"+dss.KeyWord+"%").Find(&website).Error
	}else if dss.Type == "website"{
		err = db.Orm.Debug().Model(&models.Websites{}).Where("domain LIKE ?", "%"+dss.KeyWord+"%").Find(&website).Error
	}else if dss.Type == "favicon"{
		err = db.Orm.Debug().Model(&models.Websites{}).Where("favicon LIKE ?", "%"+dss.KeyWord+"%").Find(&website).Error
	}else if dss.Type == "cid"{
		err = db.Orm.Debug().Model(&models.Websites{}).Where("cid LIKE ?", "%"+dss.KeyWord+"%").Find(&website).Error
	}
	if err != nil {
		res,_ := json.Marshal(Result{Code: 400,Msg: "query error."})
		c.Ctx.WriteString(string(res))
		return
	}

	data ,err := json.Marshal(website)
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