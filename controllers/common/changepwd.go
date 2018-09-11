package common

import (
	"MyFirstProject/models"

	"github.com/astaxie/beego"
)

type ChangePwdController struct {
	CommonBaseController
}

func (c *ChangePwdController) Get() {
	c.Data["time"] = c.GetSession("time")
	c.TplName = "changePwd.tpl"
}

func (c *ChangePwdController) Post() {
	changePwd(c, id)

	c.Redirect("/userInformation", 302)
	return
}

// reset password
// pwd and user not same at the sametime
func changePwd(c *ChangePwdController, id int) {
	// judge the user is have ?
	// err ==nil
	// user is have
	if user, err := models.SelectUserById(id); err == nil {
		// pwd is ok
		if user.Pwd == models.MD5(c.GetString("Pwd")) {
			user.Pwd = models.MD5(c.GetString("NewPwd"))
			models.UpdateUser(&user, "UserName", "Pwd")
			c.DealError(err)

			c.Redirect("/userInformation", 302)
			return
		} else {
			beego.Info("password is don't correct")
			c.Redirect("/changePwd", 302)
			return
		}
	} else {
		beego.Info("user is none")
		c.Redirect("/userInformation", 302)
		return
	}
	return
}
