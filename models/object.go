package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"auth/models/table"

)



func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default","mysql",
		"admin:yangyang123@tcp(192.168.174.131:3306)/beego?charset=utf8",30)
	orm.RegisterModel(new(table.User))
	orm.Debug =true
}
