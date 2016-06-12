// login
package controllers

import (
	"fmt"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) UserIndex() {
	this.TplName = "login.html"
}

func (this *LoginController) UserLogin() {
	username := this.GetString("userName")
	password := this.GetString("password")
	user, err := CheckLogin(username, password)
	if err == nil {
		this.Redirect("/main", 302)
		user.Pass = "" //将密码置空
		fmt.Println(user)
		this.SetSession("userinfo", user)
	} else {
		this.Data["user"] = user
		this.Data["errmsg"] = err.Error()
		this.TplName = "login.html"
	}
}

func (this *UserController) UserExit() {
	this.DelSession("userinfo")
	this.Redirect("/", 302)
}
