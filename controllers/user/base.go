package user

import (
	"MyFirstProject/controllers"
	"MyFirstProject/models"
)

type UserBaseController struct {
	controllers.BaseController
}

// the admin package role
var role models.LV = models.Use

//userid
var id int

func (b *UserBaseController) Prepare() {
	b.BaseController.Prepare()

	//session
	controllers.SessionIsOk(&b.BaseController, &id)
}
