package controllers

import (
	"backend/db"
	"backend/models"
	"backend/services/domainService"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tealeg/xlsx"
	"log"
	"net/http"
	"strconv"
	"time"
)

type CompanyController struct {
	beego.Controller
}

type XID struct {
	Id int `json:"id"`
}

func (c *CompanyController) Export() {
	var xid XID
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell

	body := c.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &xid)
	if err != nil {
		res, _ := json.Marshal(Result{Code: 400, Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}

	wb := xlsx.NewFile()
	myStyle := xlsx.NewStyle()
	myStyle.Font.Size = 11
	myStyle.Font.Name = "Microsoft YaHei UI Light"
	//取出域名信息
	var dom []models.Domain
	err = db.Orm.Model(&models.Domain{}).Where("cid=?",xid.Id).Find(&dom).Error
	if err != nil {
		res,_ := json.Marshal(Result{Code: 400,Msg: "query error."})
		c.Ctx.WriteString(string(res))
		return
	}
	sheet,err = wb.AddSheet("domain")
	err=sheet.SetColWidth(20,20,100)
	if err!=nil {}
	if err != nil {
		res,_ := json.Marshal(Result{Code: 400,Msg: "craete excel."})
		c.Ctx.WriteString(string(res))
		return
	}
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.SetStyle(myStyle)
	cell.Value = "id"
	cell = row.AddCell()
	cell.SetStyle(myStyle)
	cell.Value = "cid"
	cell = row.AddCell()
	cell.SetStyle(myStyle)
	cell.Value = "domain"
	cell = row.AddCell()
	cell.SetStyle(myStyle)
	cell.Value = "type"
	cell = row.AddCell()
	cell.SetStyle(myStyle)
	cell.Value = "updateTime"
	cell = row.AddCell()
	cell.SetStyle(myStyle)
	cell.Value = "isNew"


	for _,v:=range dom {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = strconv.Itoa(v.Id)
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = strconv.Itoa(v.Cid)
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = v.Domain
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = v.Type
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = v.UpdateTime
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		if *v.IsNew{cell.Value="true"}
	}
	//取出域名信息
	var ws []models.Websites
	err = db.Orm.Model(&models.Websites{}).Where("cid=?",xid.Id).Find(&ws).Error
	if err != nil {
		res,_ := json.Marshal(Result{Code: 400,Msg: "query error."})
		c.Ctx.WriteString(string(res))
		return
	}
	sheet,err = wb.AddSheet("website")

	if err != nil {
		res,_ := json.Marshal(Result{Code: 400,Msg: "query error."})
		c.Ctx.WriteString(string(res))
		return
	}
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.SetStyle(myStyle)
	cell.Value = "#"
	cell = row.AddCell()
	cell.SetStyle(myStyle)
	cell.Value = "Website"
	cell = row.AddCell()
	cell.SetStyle(myStyle)
	cell.Value = "Title"
	cell = row.AddCell()
	cell.SetStyle(myStyle)
	cell.Value = "Ip"
	cell = row.AddCell()
	cell.SetStyle(myStyle)
	cell.Value = "Favicon"
	cell = row.AddCell()
	cell.SetStyle(myStyle)
	cell.Value = "Headers"
	cell = row.AddCell()
	cell.SetStyle(myStyle)
	cell.Value = "UpdateTime"
	cell = row.AddCell()
	cell.SetStyle(myStyle)
	cell.Value = "CDN"
	cell = row.AddCell()
	cell.SetStyle(myStyle)
	cell.Value = "Cert"
	for _,v:=range ws {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = strconv.Itoa(v.Id)
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = strconv.Itoa(v.Cid)
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = v.Domain
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = v.Website
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = v.Ips
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = v.Favicon
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = v.FaviconUrl
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = v.Title
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = v.Headers
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = v.UpdateTime
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		if *v.CDN {cell.Value = "Y"}
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = v.Cert

	}
	path:="tmp/result.xlsx"
	err = wb.Save(path)
	if err != nil {
		res,_ := json.Marshal(Result{Code: 400,Msg: "xlsx save error."})
		c.Ctx.WriteString(string(res))
		return
	}


	c.Ctx.ResponseWriter.Header().Add("Content-Disposition", `attachment; filename=result.xlsx`)
	c.Ctx.ResponseWriter.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet; "+"filename=result.xlsx")
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, path)
}

func (c *CompanyController) Ex() {
	c.Ctx.Output.Download("tmp/result.xlsx","result.xlsx")

}

func (c *CompanyController) GetCompanyInfo() {
	var param = &PageParam{}
	body := c.Ctx.Input.RequestBody
	err := json.Unmarshal(body, param)
	if err != nil {
		res, _ := json.Marshal(Result{Code: 400, Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}
	var count int64
	if param.Keyword == "" {
		err = db.Orm.Debug().Model(&models.Company{}).Count(&count).Error
	} else if param.Type == "companyName" {
		err = db.Orm.Debug().Model(&models.Company{}).Where("companyName LIKE ?", "%"+param.Keyword+"%").Count(&count).Error
	} else if param.Type == "domain" {
		err = db.Orm.Debug().Model(&models.Company{}).Where("domain LIKE ?", "%"+param.Keyword+"%").Count(&count).Error
	} else if param.Type == "keyWord" {
		err = db.Orm.Debug().Model(&models.Company{}).Where("keyWord LIKE ?", "%"+param.Keyword+"%").Count(&count).Error
	}
	if err != nil {
		res, _ := json.Marshal(Result{Code: 400, Msg: "database error."})
		c.Ctx.WriteString(string(res))
	}

	var dom = &[]models.Company{}
	if param.Keyword == "" {
		err = db.Orm.Debug().Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Model(&models.Company{}).Find(&dom).Error
	} else if param.Type == "companyName" {
		err = db.Orm.Debug().Limit(param.PageSize).Offset((param.Page-1)*param.PageSize).Model(&models.Company{}).Where("companyName LIKE ?", "%"+param.Keyword+"%").Find(&dom).Error
	} else if param.Type == "domain" {
		err = db.Orm.Debug().Limit(param.PageSize).Offset((param.Page-1)*param.PageSize).Model(&models.Company{}).Where("domain LIKE ?", "%"+param.Keyword+"%").Find(&dom).Error
	} else if param.Type == "keyWord" {
		err = db.Orm.Debug().Limit(param.PageSize).Offset((param.Page-1)*param.PageSize).Model(&models.Company{}).Where("keyWord LIKE ?", "%"+param.Keyword+"%").Find(&dom).Error
	}
	if err != nil {
		res, _ := json.Marshal(Result{Code: 400, Msg: "database error."})
		c.Ctx.WriteString(string(res))
	}

	data, err := json.Marshal(dom)
	if err != nil {
		res, _ := json.Marshal(Result{Code: 400, Msg: "json marshal error.."})
		c.Ctx.WriteString(string(res))
	}
	type Res struct {
		Code  int    `json:"code"`
		Msg   string `json:"msg"`
		Total int    `json:"total"`
	}
	res, _ := json.Marshal(Res{Code: 200, Msg: string(data), Total: int(count)})

	c.Ctx.WriteString(string(res))
}

func (c *CompanyController) ScanCompanyInfo() {
	type TId struct {
		Id int `json:"id"`
	}
	var Tid = &TId{}
	body := c.Ctx.Input.RequestBody
	err := json.Unmarshal(body, Tid)
	if err != nil {
		log.Println("company.go line:73 error:[", err, "]")
		res, _ := json.Marshal(Result{Code: 400, Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}
	var cmp models.Company
	err = db.Orm.Model(&models.Company{}).Where("id=?", Tid.Id).Find(&cmp).Error
	if err != nil {
		log.Println("company.go line:81 error:[", err, "]")
		res, _ := json.Marshal(Result{Code: 400, Msg: "database error."})
		c.Ctx.WriteString(string(res))
		return
	}
	if len(cmp.Domain) > 0 {
		go domainService.ScanDomain(cmp)
	} else {
		res, _ := json.Marshal(Result{Code: 200, Msg: "not found root domain."})
		c.Ctx.WriteString(string(res))
		return
	}
	res, _ := json.Marshal(Result{Code: 200, Msg: "scanning."})
	c.Ctx.WriteString(string(res))

}

//{"projectType":"src","companyName":"asd","domain":"asd","srcUrl":{"_shallow":false,"dep":{"w":0,"n":0},"__v_isRef":true,"_rawValue":"asd","_value":"asd"},"keyWord":"asd","monitorStatus":"1","monitorRate":"30","vulnScanStatus":"1","vulnScanRate":"30"}

func (c *CompanyController) GetCompanyByKey() {
	var err error
	var dss = &SearchStrut{}
	body := c.Ctx.Input.RequestBody
	err = json.Unmarshal(body, dss)
	if err != nil {
		res, _ := json.Marshal(Result{Code: 400, Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}
	fmt.Println("type:", dss.Type, "  keyword:", dss.KeyWord)

	var count int64
	var dom = &[]models.Company{}
	if dss.Type == "companyName" {
		err = db.Orm.Debug().Model(&models.Company{}).Where("companyName LIKE ?", "%"+dss.KeyWord+"%").Find(&dom).Count(&count).Error
	} else if dss.Type == "domain" {
		err = db.Orm.Debug().Model(&models.Company{}).Where("domain LIKE ?", "%"+dss.KeyWord+"%").Count(&count).Find(&dom).Error
	} else if dss.Type == "keyWord" {
		err = db.Orm.Debug().Model(&models.Company{}).Where("keyWord LIKE ?", "%"+dss.KeyWord+"%").Count(&count).Find(&dom).Error
	}
	if err != nil {
		res, _ := json.Marshal(Result{Code: 400, Msg: "query error."})
		c.Ctx.WriteString(string(res))
		return
	}

	data, err := json.Marshal(dom)
	if err != nil {
		res, _ := json.Marshal(Result{Code: 400, Msg: "json marshal error.."})
		c.Ctx.WriteString(string(res))
	}
	type Res struct {
		Code  int    `json:"code"`
		Msg   string `json:"msg"`
		Total int    `json:"total"`
	}
	res, _ := json.Marshal(Res{Code: 200, Msg: string(data), Total: int(count)})

	c.Ctx.WriteString(string(res))
}

func (c *CompanyController) NewCompanyInfo() {
	var tc = &models.Company{}
	body := c.Ctx.Input.RequestBody
	err := json.Unmarshal(body, tc)
	if err != nil {
		fmt.Println(err)
		res, _ := json.Marshal(Result{Code: 400, Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}

	tc.LastUpdateTime = time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")

	err = db.Orm.Model(&models.Company{}).Create(&tc).Error
	if err != nil {
		res, _ := json.Marshal(Result{Code: 400, Msg: "update data error."})
		c.Ctx.WriteString(string(res))
		return
	} else {
		go domainService.ScanDomain(*tc)
		res, _ := json.Marshal(Result{Code: 200, Msg: "update success. scanning..."})
		c.Ctx.WriteString(string(res))
		return
	}

}

func (c *CompanyController) GetSelectOption() {

	cmp := []models.Company{}
	type Option struct {
		Value string `json:"value"`
		Lable string `json:"label"`
	}
	err := db.Orm.Model(&models.Company{}).Find(&cmp).Error
	if err != nil {
		fmt.Println(err)
		res, _ := json.Marshal(Result{Code: 400, Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}
	var option []Option
	for _, v := range cmp {
		option = append(option, Option{Value: strconv.Itoa(v.Id), Lable: v.CompanyName})
	}
	data, _ := json.Marshal(option)
	res, _ := json.Marshal(Result{Code: 200, Msg: string(data)})
	c.Ctx.WriteString(string(res))
	return

}

func (c *CompanyController) DelCompanyByid() {
	type TId struct {
		Id int `json:"id"`
	}
	var Tid = &TId{}
	body := c.Ctx.Input.RequestBody
	err := json.Unmarshal(body, Tid)
	if err != nil {
		fmt.Println(err)
		res, _ := json.Marshal(Result{Code: 400, Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}

	var OneCmp models.Company
	err = db.Orm.Debug().Delete(&OneCmp, Tid.Id).Error
	if err != nil {
		res, _ := json.Marshal(Result{Code: 400, Msg: "delete data error."})
		c.Ctx.WriteString(string(res))
		return
	}

	var tDom models.Domain
	err = db.Orm.Debug().Where("cid=?", Tid.Id).Delete(&tDom).Error
	if err != nil {
		res, _ := json.Marshal(Result{Code: 400, Msg: "delete data error."})
		c.Ctx.WriteString(string(res))
		return
	}

	var tws models.Websites
	err = db.Orm.Debug().Where("cid=?", Tid.Id).Delete(&tws).Error
	if err != nil {
		res, _ := json.Marshal(Result{Code: 400, Msg: "delete data error."})
		c.Ctx.WriteString(string(res))
		return
	}

	res, _ := json.Marshal(Result{Code: 200, Msg: "delete success."})
	c.Ctx.WriteString(string(res))
}

func (c *CompanyController) GetCompanyByid() {
	type TId struct {
		Id string `json:"id"`
	}
	var Tid = &TId{}
	body := c.Ctx.Input.RequestBody
	err := json.Unmarshal(body, Tid)
	if err != nil {
		fmt.Println(err)
		res, _ := json.Marshal(Result{Code: 400, Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}

	var OneCmp models.Company
	err = db.Orm.Debug().Model(&models.Company{}).Find(&OneCmp, "id=?", Tid.Id).Error
	if err != nil {
		res, _ := json.Marshal(Result{Code: 400, Msg: "update data error."})
		c.Ctx.WriteString(string(res))
		return
	} else {
		msg, err := json.Marshal(OneCmp)
		if err != nil {
			res, _ := json.Marshal(Result{Code: 400, Msg: "Marshal onecmp error."})
			c.Ctx.WriteString(string(res))
			return
		}
		res, _ := json.Marshal(Result{Code: 200, Msg: string(msg)})
		c.Ctx.WriteString(string(res))
		return
	}

}

func (c *CompanyController) UpdateCompanyInfo() {
	var tc = &models.Company{}
	body := c.Ctx.Input.RequestBody
	err := json.Unmarshal(body, tc)
	if err != nil {
		fmt.Println(err)
		res, _ := json.Marshal(Result{Code: 400, Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}

	tc.LastUpdateTime = time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")

	//输入的时候就把回车换

	//data := map[string]
	//fmt.Println(helper.Struct2Map(*tc))

	err = db.Orm.Debug().Model(&models.Company{}).Where("id=?", tc.Id).Updates(&tc).Error
	if err != nil {
		res, _ := json.Marshal(Result{Code: 400, Msg: "update data error."})
		c.Ctx.WriteString(string(res))
		return
	} else {
		res, _ := json.Marshal(Result{Code: 200, Msg: "update success."})
		c.Ctx.WriteString(string(res))
		return
	}

}
