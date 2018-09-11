package admin

import (
	"MyFirstProject/controllers"
	"MyFirstProject/models"
)

type AdminBaseController struct {
	controllers.BaseController
}

// the admin package role
var role models.LV = models.Admin

func (b *AdminBaseController) Prepare() {
	b.BaseController.Prepare()

	//session and role check
	var id int
	controllers.RoleAndSessionIsOk(&b.BaseController, &id, role)

}
