package main

import (
	. "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"hanloon/controllers"
)

type HanloonAdmin struct {
	Id         int
	Account    string
	Password   string
	Add_time   int
	Last_login int
	Login_ip   string
	Real_name  string
	Privilage  int
}

func init() {
	orm.RegisterModel(new(HanloonAdmin))
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:123123@/hanloon_customization?charset=utf8", 30)
}

func main() {
	beego.AppName = "Hanloon"

	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")

	admin := new(HanloonAdmin)
	admin.Id = 1
	admin.Account = "abellee"
	admin.Password = "lijinbei"
	admin.Add_time = 123123
	admin.Last_login = 321321
	admin.Login_ip = "192.168.1.1"
	admin.Real_name = "李金贝"
	admin.Privilage = 1

	Print(o.Insert(admin))

	//定义登录页
	beego.Router("/", &controllers.LoginController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/index", &controllers.LoginController{})

	beego.Run()
}
