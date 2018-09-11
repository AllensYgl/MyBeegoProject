package common

import (
	"MyFirstProject/controllers"
)

type ExitController struct {
	controllers.BaseController
}

func (c *ExitController) Get() {
	controllers.RemoveCookie(&c.BaseController)
	controllers.RemoveSession(&c.BaseController)
	c.Redirect("/", 302)
	return
}
