package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"auth/models/table"

	"fmt"
	"github.com/astaxie/beego"
)



func init() {
	if beego.AppConfig.String("db") == "mysql"{
		sqlurl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
			beego.AppConfig.String("mysql::mysqluser"),
			beego.AppConfig.String("mysql::mysqlpass"),
			beego.AppConfig.String("mysql::mysqlurls"),
			beego.AppConfig.String("mysql::mysqlport"),
			beego.AppConfig.String("mysql::mysqldb"),
		)
		orm.RegisterDriver("mysql", orm.DRMySQL)
		orm.RegisterDataBase("default","mysql", sqlurl,30)

	}else if beego.AppConfig.String("db") == "sqlite"{
		orm.RegisterDriver("sqlite", orm.DRSqlite)
		orm.RegisterDataBase("default","sqlite3",beego.AppConfig.String("sqlite::sqliteurls"))
	}else {
		orm.RegisterDriver("sqlite", orm.DRSqlite)
		orm.RegisterDataBase("default","sqlite3","default.db")
	}

	orm.RegisterModel(new(table.User),new(table.Group),new(table.Permission))
	orm.Debug =true
}
