package controllers

import (
	m "easyui_crud/models"

	"easyui_crud/lib"
	"errors"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Rsp(status bool, str string) {
	this.Data["json"] = &map[string]interface{}{"succ": status, "msg": str}
	this.ServeJSON()
}

func CheckLogin(username string, password string) (user m.User, err error) {
	user = m.GetUserByUsername(username)
	if user.Id == 0 {
		user.Pass = password
		return user, errors.New("用户不存在")
	}
	if user.Pass != lib.Pwdhash(password) {
		user.Pass = password
		return user, errors.New("密码错误")
	}
	return user, nil
}
