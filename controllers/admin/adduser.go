package admin

import (
	"MyFirstProject/models"
	"github.com/astaxie/beego"
)

type AdminAddController struct {
	AdminBaseController
}

func (c *AdminAddController) Get() {
	c.Data["time"] = c.GetSession("time")
	c.TplName = "adminAdd.tpl"
}

func (c *AdminAddController) Post() {
	// create the insert user model
	user := models.User{}
	rol, err := c.GetInt("Role")
	if err != nil {
		beego.Info("error role", err)
	}
	user.Role = models.LV(rol)
	user.Pwd = models.MD5(c.GetString("Pwd"))
	user.Mail = c.GetString("Mail")
	user.UserName = c.GetString("UserName")

	//insert
	err=models.InsertUser(&user)
	c.DealError(err)
	
	c.Redirect("/admin", 302)
	return
}
