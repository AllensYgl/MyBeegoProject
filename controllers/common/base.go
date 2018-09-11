package common

import (
	"MyFirstProject/controllers"
	"MyFirstProject/models"
)

type CommonBaseController struct {
	controllers.BaseController
}

// the admin package role
var role models.LV = models.Use

//userid
var id int

func (b *CommonBaseController) Prepare() {
	b.BaseController.Prepare()

	//session
	controllers.SessionIsOk(&b.BaseController, &id)
}
