package routers

import (
	"MyFirstProject/controllers/admin"
	"MyFirstProject/controllers/user"
	"MyFirstProject/controllers/common"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &common.LoginController{})
	beego.Router("/index", &common.IndexController{}, "post:Post")
	beego.Router("/admin", &admin.AdminController{})
	beego.Router("/userInformation", &user.UserInformationController{})
	beego.Router("/changeMail/", &common.ChangeMailController{}, "get:Get;post:Post")
	beego.Router("/changePwd/", &common.ChangePwdController{}, "get:Get;post:Post")
	beego.Router("/admin/delete/?:Id", &admin.UserDeleteController{})
	beego.Router("/admin/add", &admin.AdminAddController{}, "get:Get;post:Post")
	beego.Router("/exit", &common.ExitController{})
}
