package user

import (
	"MyFirstProject/models"
)

type UserInformationController struct {
	UserBaseController
}

func (c *UserInformationController) Get() {
	user, err := models.SelectUserById(id)
	if err == nil {

		//render datas
		c.Data["Role"] = user.Role
		c.Data["UserName"] = user.UserName
		c.Data["Email"] = user.Mail
		c.Data["time"] = c.GetSession("time")

		c.TplName = "userInformation.tpl"
	} else {
		c.Redirect("/", 302)
		return
	}
}
