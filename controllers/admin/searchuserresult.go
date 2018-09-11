package admin

import (
	"MyFirstProject/models"
)



type AdminController struct {
	AdminBaseController
}

func (c *AdminController) Get() {

	//get search condition
	params := getSearchInputInformation(c)
	users, num := models.SearchUser(&params)

	// set search information
	c.Data["num"] = num
	c.Data["users"] = users
	c.Data["time"] = c.GetSession("time")

	c.TplName = "admin.tpl"

}

func getSearchInputInformation(c *AdminController) map[string]interface{} {
	params := make(map[string]interface{})
	params["UserName"] = c.GetString("UserName")
	params["Mail"] = c.GetString("Mail")
	params["Role"] = c.GetString("Role")
	params["Id"] = c.GetString("Id")
	//...
	return params
}
