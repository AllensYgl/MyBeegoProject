package common

import (
	"MyFirstProject/controllers"
	"MyFirstProject/models"

	"github.com/astaxie/beego"
)

type IndexController struct {
	controllers.BaseController
}

func (c *IndexController) Get() {
	c.TplName = "index.tpl"
}

func (c *IndexController) Post() {

	params := getFormInformationAndEncryptPwd(c)

	if user, err := models.SelectUser(params); err == nil {

		controllers.SetCookie(&c.BaseController, &user)
		controllers.SetSession(&c.BaseController, &user)

		// redirect to corresponding tpl
		switch user.Role {
		case models.Use:
			c.Redirect("/userInformation", 302)
			return
		case models.Admin:
			c.Redirect("/admin", 302)
			return
		}
	} else {
		beego.Info("username or password is error", err)
		c.Redirect("/", 302)
		return
	}
}

//get form information
/*	form information
	key:
		UserName
		Pwd	(MD5)
*/
func getFormInformationAndEncryptPwd(c *IndexController) map[string]interface{} {
	params := make(map[string]interface{})
	params["UserName"] = c.GetString("UserName")
	params["Pwd"] = models.MD5(c.GetString("Pwd"))
	return params
}
