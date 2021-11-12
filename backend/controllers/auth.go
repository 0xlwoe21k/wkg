package controllers

import (
	"backend/db"
	Gjwt "backend/libs/jwt"
	"backend/models"
	"encoding/json"
	"github.com/astaxie/beego"

	"github.com/mailru/easyjson/jlexer"
)

type Result struct {
	Code 	int `json:"code"`
	Msg		string `json:"msg"`
}

type AuthController struct {
	beego.Controller
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u User) UnmarshalEasyJSON(w *jlexer.Lexer) {
	panic("implement me")
}

func (c *AuthController) Post() {

	user := &User{}
	body := c.Ctx.Input.RequestBody
	err := json.Unmarshal(body,user)
	if err != nil{
		c.Ctx.WriteString("json unmarshal error.")
		return
	}

	var count int64
	err = db.Orm.Debug().Model(&models.Users{}).Where("username=? and password=?",user.Username,user.Password).Count(&count).Error
	if err != nil {
		c.Ctx.WriteString("query error.")
		return
	}

	if count > 0{
		res,err := json.Marshal(Result{Code: 200,Msg: "login sucess."})
		token := Gjwt.GenToken()
		//cookie := http.Cookie{Name: "Authorization", Value: token, Path: "/", MaxAge: 3600}
		c.Ctx.SetCookie("Authorization",token)
		if err != nil{
			c.Ctx.WriteString("json marshal error.")
			return
		}
		c.Ctx.WriteString(string(res))
		return
	}else{
		res,err := json.Marshal(Result{Code: 400,Msg: "login failed."})
		if err != nil{
			c.Ctx.WriteString("json marshal error.")
			return
		}
		c.Ctx.WriteString(string(res))
	}
}