package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"auth/models/table"

	"fmt"
	"github.com/astaxie/beego"
)



func init() {
	sqlurl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		beego.AppConfig.String("mysqluser"),
		beego.AppConfig.String("mysqlpass"),
		beego.AppConfig.String("mysqlurls"),
		beego.AppConfig.String("mysqlport"),
		beego.AppConfig.String("mysqldb"),
	)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default","mysql", sqlurl,30)
	orm.RegisterModel(new(table.User),new(table.Group),new(table.Permission))
	orm.Debug =true
}
