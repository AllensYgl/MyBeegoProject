package controllers

import (
	"MyFirstProject/models"
	"fmt"
	"time"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

var Version string
var ApiServer string

func (b *BaseController) Prepare() {
	if beego.BConfig.Listen.EnableHTTPS && !b.Ctx.Input.IsSecure() {
		redirect := fmt.Sprintf("https://%s:%d%s", b.Ctx.Input.Host(), beego.BConfig.Listen.HTTPSPort, b.Ctx.Input.URL())
		b.Redirect(redirect, 302)
	}
	b.Data["Version"] = Version
}

//render data
func Render(c *BaseController, user models.User) {
	c.Data["UserName"] = user.UserName
	c.Data["Email"] = user.Mail
	c.Data["time"] = c.GetSession("time")
}

// judge the session is ok?
//	return id
// 0 means none
// session is none redirect /login
func SessionIsOk(c *BaseController, userid *int) {
	if id, ok := c.GetSession("Id").(int); ok {
		*userid = id
	} else {
		beego.Info("please go to login")
		c.Redirect("/", 302)
	}
}

//dedge the role is ok?
// role is permission denied
// redirect /user
func RoleIsOk(c *BaseController, rol models.LV) {
	if role, ok := c.GetSession("Role").(models.LV); !ok || role != rol {
		beego.Info("permission denied")
		c.Redirect("/userInformation", 302)
	}
}

func RoleAndSessionIsOk(c *BaseController, userid *int, rol models.LV) {
	SessionIsOk(c, userid)
	RoleIsOk(c, rol)
}

//set session
/*	session information
	key:
		UserName
		Id
		Role
		time
*/
func SetSession(c *BaseController, user *models.User) {
	c.SetSession("UserName", user.UserName)
	c.SetSession("Id", user.Id)
	c.SetSession("Role", user.Role)
	c.SetSession("time", time.Now().Format("15:04:05"))
}

//set cookie
/*	cookie information
	key:
		UserName
		Pwd (MD5)
*/
func SetCookie(c *BaseController, user *models.User) {
	c.Ctx.SetCookie("UserName", user.UserName, 1000, "/") // 设置cookie  name
	c.Ctx.SetCookie("Pwd", user.Pwd, 1000, "/")           // 设置cookie	 password
}

//remove cookie
func RemoveCookie(c *BaseController) {
	c.Ctx.SetCookie("UserName", "", -1)
	c.Ctx.SetCookie("Pwd", "", -1)
}

//remove session
func RemoveSession(c *BaseController) {
	c.DelSession("UserName")
	c.DelSession("Id")
	c.DelSession("Role")
	c.DelSession("time")
}

//deal with error
func (b *BaseController) DealError(err error) {
	if err != nil {
		beego.Info(err)
	}
}
