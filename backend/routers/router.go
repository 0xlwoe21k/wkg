package routers

import (
	"backend/controllers"
	Gjwt "backend/libs/jwt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func init() {

	beego.InsertFilter("/v1/*", beego.BeforeRouter, func(ctx *context.Context) {
		cookie := ctx.Request.Header.Get("token")
		if !Gjwt.CheckToken(cookie) {
			//res ,_ := json.Marshal(&Result{Code: 302,Msg: "invild auth."})
			ctx.ResponseWriter.WriteHeader(401)
			//http.Redirect(ctx.ResponseWriter, ctx.Request, "/api", http.StatusMovedPermanently)
		}
	})

	beego.Router("/debug/pprof", &controllers.ProfController{})
	beego.Router("/debug/pprof/:app([\\w]+)", &controllers.ProfController{})

	beego.Router("/v1/getDashboardInfo", &controllers.DashBoard{}, "get:GetDashboardInfo")

	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.AuthController{}, "post:Post")
	beego.Router("/v1/checklogin", &controllers.AuthController{}, "get:CheckLogin")

	beego.Router("/v1/scanWebsiteInfo", &controllers.WebSiteController{}, "get:ScanWebsiteInfo")

	beego.Router("/v1/ex", &controllers.CompanyController{}, "get:Ex")

	beego.Router("/v1/export", &controllers.CompanyController{}, "post:Export")

	beego.Router("/v1/scanCompanyInfo", &controllers.CompanyController{}, "post:ScanCompanyInfo")
	beego.Router("/v1/getCompanyInfo", &controllers.CompanyController{}, "post:GetCompanyInfo")
	beego.Router("/v1/newCompany", &controllers.CompanyController{}, "post:NewCompanyInfo")
	beego.Router("/v1/getCompanyByid", &controllers.CompanyController{}, "post:GetCompanyByid")
	beego.Router("/v1/getCompanyInfoByKey", &controllers.CompanyController{}, "post:GetCompanyByKey")
	beego.Router("/v1/updateCompanyInfo", &controllers.CompanyController{}, "post:UpdateCompanyInfo")
	beego.Router("/v1/delCompanyByid", &controllers.CompanyController{}, "post:DelCompanyByid")

	beego.Router("/v1/getSelectOption", &controllers.CompanyController{}, "get:GetSelectOption")

	//beego.Router("/v1/searchCompanyByid",&controllers.CompanyController{},"post:SearchCompanyByid")

	beego.Router("/v1/getDomainInfoByCid", &controllers.DomainController{}, "post:GetDomainInfoByCid")

	beego.Router("/v1/getNewDomainInfo", &controllers.DomainController{}, "post:GetNewDomainInfo")
	beego.Router("/v1/getDomainInfo", &controllers.DomainController{}, "post:GetDomainInfo")
	beego.Router("/v1/getDomainInfoByKey", &controllers.DomainController{}, "post:GetDomainInfoByKey")
	beego.Router("/v1/readFlagDomainInfoById", &controllers.DomainController{}, "post:ReadFlagDomainInfoById")
	beego.Router("/v1/readAllFlagDomainInfo", &controllers.DomainController{}, "get:ReadAllFlagDomainInfo")


	beego.Router("/v1/scanNew", &controllers.WebSiteController{}, "get:ScanNew")
	beego.Router("/v1/getWebSiteInfoByCid", &controllers.WebSiteController{}, "post:GetWebSiteInfoByCid")
	beego.Router("/v1/getWebSiteInfo", &controllers.WebSiteController{}, "post:GetWebSiteInfo")
	beego.Router("/v1/getWebSiteInfoByKey", &controllers.WebSiteController{}, "post:GetWebSiteInfoByKey")
	beego.Router("/v1/getNewWebSiteInfo", &controllers.WebSiteController{}, "post:GetNewWebSiteInfo")

	beego.Router("/v1/getIPsInfo", &controllers.IPsController{}, "post:GetIPsInfo")
	beego.Router("/v1/getNewIPsInfo", &controllers.IPsController{}, "post:GetNewIPsInfo")

	beego.Router("/v2/domainScan", &controllers.DomainScanController{}, "get:DomainScan")
	beego.Router("/v2/VulnScanSingle", &controllers.VulnController{}, "get:VulnScanSingle")
	beego.Router("/v2/VulnScanMulti", &controllers.VulnController{}, "get:VulnScanMulti")
}
