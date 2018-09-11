package main

import (
	_ "MyFirstProject/routers"

	"github.com/astaxie/beego"
)

func main() {
	// user := models.User{}
	// user.Mail = "a"
	// user.Role = models.Admin
	// user.UserName = "a"
	// user.Pwd = models.MD5("1")
	// models.InsertUser(&user)
	//models.SelectUser(params)
	beego.Run()
}
