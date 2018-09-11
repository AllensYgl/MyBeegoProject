package admin

import (
	"MyFirstProject/models"

	"github.com/astaxie/beego"
)

type UserDeleteController struct {
	AdminBaseController
}

func (c *UserDeleteController) Get() {

	if userid, err := c.GetInt(":Id"); err == nil {
		err=models.DeleteUserById(userid)
		c.DealError(err)
	} else {
		beego.Info("no Id ", err)
	}

	c.Redirect("/admin", 302)
	return
}
