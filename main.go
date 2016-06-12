package main

import (
	_ "easyui_crud/models"
	_ "easyui_crud/routers" //调用routes下的init()方法来注册路由

	"github.com/astaxie/beego"
)

const VERSION = "1.0.0"

func main() {
	beego.AppConfig.Set("version", VERSION)
	beego.Run() //运行程序
}
