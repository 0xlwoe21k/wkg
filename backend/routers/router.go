package routers

import (
	"backend/controllers"
	Gjwt "backend/libs/jwt"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)


type Result struct {
	Code 	int `json:"code"`
	Msg		string `json:"msg"`
}
func init() {



	beego.InsertFilter("/v1/*", beego.BeforeRouter, func(ctx *context.Context) {
		cookie, err := ctx.Request.Cookie("Authorization")
		if err != nil || !Gjwt.CheckToken(cookie.Value) {
			res ,_ := json.Marshal(&Result{Code: 302,Msg: "invild jwt."})
			ctx.WriteString(string(res))
			//http.Redirect(ctx.ResponseWriter, ctx.Request, "/api", http.StatusMovedPermanently)
		}
	})

	beego.Router("/debug/pprof", &controllers.ProfController{})
	beego.Router("/debug/pprof/:app([\\w]+)", &controllers.ProfController{})


	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.AuthController{},"post:Post")

	beego.Router("/v1/getCompanyInfo",&controllers.CompanyController{},"post:GetCompanyInfo")
	beego.Router("/v1/newCompany",&controllers.CompanyController{},"post:NewCompanyInfo")
	beego.Router("/v1/getCompanyByid",&controllers.CompanyController{},"post:GetCompanyByid")
	beego.Router("/v1/updateCompanyInfo",&controllers.CompanyController{},"post:UpdateCompanyInfo")

	beego.Router("/v1/getDomainInfo",&controllers.DomainController{},"post:GetDomainInfo")
	beego.Router("/v1/getDomainInfoByKey",&controllers.DomainController{},"post:GetDomainInfoByKey")




	beego.Router("/v2/domainScan",&controllers.DomainScanController{},"get:DomainScan")
	beego.Router("/v2/VulnScanSingle",&controllers.VulnController{},"get:VulnScanSingle")
	beego.Router("/v2/VulnScanMulti",&controllers.VulnController{},"get:VulnScanMulti")
}
