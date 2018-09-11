package common

import (
	"MyFirstProject/models"

	"github.com/astaxie/beego"
)

type ChangeMailController struct {
	CommonBaseController
}

func (c *ChangeMailController) Get() {

	//user is have
	if user, err := models.SelectUserById(id); err == nil {
		// select is ok
		// to rander data
		c.Data["UserName"] = user.UserName
		c.Data["Mail"] = user.Mail
		c.Data["time"] = c.GetSession("time")

		c.TplName = "changeMail.tpl"
		return
	} else {
		beego.Info("user is not found", err)
	}

	c.Redirect("/userInformation", 302)
	return
}

func (c *ChangeMailController) Post() {

	// get id
	// select user
	if user, err := models.SelectUserById(id); err == nil {
		// user is have
		user.Mail = c.GetString("Mail")
		err=models.UpdateUser(&user, "Mail")
		c.DealError(err)

		c.Redirect("/userInformation", 302)
		return
	} else {
		beego.Info("user is none")
	}

	c.Redirect("/changeMail", 302)
	return
}
