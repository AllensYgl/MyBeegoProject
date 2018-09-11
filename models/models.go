package models

import (
	"crypto/md5"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type LV int

const (
	Admin LV = 1
	Use      = 2
	//...
)

type User struct {
	Id       int
	UserName string
	Role     LV
	Mail     string
	Pwd      string
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysqlparams"), 30)
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
	beego.Info("start model success")
}

//Select..
// if select return user else return nil and err
func SelectUserById(id int) (User, error) {
	o := orm.NewOrm()
	user := User{Id: id}
	err := o.Read(&user)
	return user, err
}

//Insert
func InsertUser(user *User) error {
	// the user is have? we need to select
	params := make(map[string]interface{})
	params["UserName"] = user.UserName
	params["Pwd"] = user.Pwd
	//err!=nil means no user
	// can insert
	if _, err := SelectUser(params); err != nil {
		if _, err := SelectUserByParams(user, "Mail"); err != nil {
			o := orm.NewOrm()
			// err==nil
			// insert ok
			if _, err := o.Insert(user); err == nil {
				beego.Info("success insert id=", user.Id)
				return nil
			} else {
				return err
			}
		} else {
			return err
		}
	} else {
		return err
	}
}

//SearchUser..
// dim select
func SearchUser(params *map[string]interface{}) ([]*User, int64) {
	o := orm.NewOrm()
	qseter := o.QueryTable("user")
	for key, value := range *params {
		qseter = qseter.Filter(key+"__icontains", value)
	}
	var result []*User
	num, _ := qseter.All(&result)
	//beego.Info(result)
	return result, num
}

//select by params options
// select params wil be fully used
func SelectUser(params map[string]interface{}) (User, error) {
	o := orm.NewOrm()
	qseter := o.QueryTable("user")
	for key, value := range params {
		if value != "" {
			qseter = qseter.Filter(key, value)
		}
	}
	//qseter = qseter.Filter(params["UserPwd"], value)
	var result User
	err := qseter.One(&result)
	if err != nil {
		beego.Error(err)
	}
	return result, err
}

//Delete
func DeleteUserById(id int) error {
	o := orm.NewOrm()
	if num, err := o.Delete(&User{Id: id}); err == nil {
		beego.Info("success delete id=", id, num)
		return nil
	} else {
		return err
	}
}

//Select..
// if select return user else return nil and err
// if have user err == nil
// select params can choice
func SelectUserByParams(user *User, col ...string) (User, error) {
	o := orm.NewOrm()
	use := user
	err := o.Read(use, col...)
	return *use, err
}

//Update
func UpdateUser(params *User, cols ...string) error {
	o := orm.NewOrm()
	// err ==nil means have user
	// can update
	if user, err := SelectUserById((*params).Id); err == nil {
		if _, err := SelectUserByParams(params, cols...); err != nil {
			user = *params
			if _, err := o.Update(&user, cols...); err == nil {
				beego.Info("success update id", (*params).Id)
				return nil
			} else {
				return err
			}
		} else {
			return err
		}
	} else {
		return err
	}
}

//md5
func MD5(s string) string {
	data := []byte(s)
	result := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", result) //将[]byte转成16进制\
	return md5str1
}
