package common

import (
	"MyFirstProject/controllers"
	"MyFirstProject/models"
)

type LoginController struct {
	controllers.BaseController
}

func (c *LoginController) Get() {

	name := c.Ctx.GetCookie("UserName")
	password := c.Ctx.GetCookie("Pwd")

	user, err := models.SelectUserByParams(&models.User{UserName: name, Pwd: password}, "UserName", "Pwd")
	if err == nil {
		controllers.SetSession(&c.BaseController, &user)
		c.Redirect("/userInformation", 302)
		return
	}

	c.TplName = "login.tpl"
}
