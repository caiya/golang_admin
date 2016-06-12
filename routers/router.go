package routers

import (
	"easyui_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/user/index", &controllers.UserController{}, "get:UserIndex")
	beego.Router("/user/list", &controllers.UserController{}, "post:UserList")
	beego.Router("/user/del", &controllers.UserController{}, "post:UserDel")
	beego.Router("/user/add", &controllers.UserController{}, "post:UserAdd")
	beego.Router("/user/update", &controllers.UserController{}, "post:UserUpdate")
	beego.Router("/user/upload", &controllers.UserController{}, "post:UserUpload")

	beego.Router("/", &controllers.LoginController{}, "get:UserIndex")
	beego.Router("/", &controllers.LoginController{}, "post:UserLogin")
	beego.Router("/main", &controllers.UserController{}, "get:Index")
	beego.Router("/exit", &controllers.UserController{}, "get:UserExit")
}
