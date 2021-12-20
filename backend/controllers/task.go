package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

type TaskController struct {
	beego.Controller
}

type VulnScan struct {

}

type TaskStrut struct {
	Type string `json:"type"`			//vulnscan:漏洞扫描    dirscan：目录扫描       
}

func (c *TaskController)NewTask()  {
	var param = &PageParam{}
	body := c.Ctx.Input.RequestBody
	err := json.Unmarshal(body,param)
	if err != nil{
		res,_ := json.Marshal(Result{Code: 400,Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}
}