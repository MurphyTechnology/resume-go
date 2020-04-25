package db

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"resume/models/dao"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	maxIdle := 30
	maxConn := 30
	sqltype := beego.AppConfig.String("sqltype")
	sqladdress := beego.AppConfig.String("sqladd")
	sqluser := beego.AppConfig.String("sqluser")
	sqlpwd := beego.AppConfig.String("sqlpwd")
	sqlother := beego.AppConfig.String("sqlother")
	orm.RegisterDataBase("default",
		sqltype,
		sqluser+":"+sqlpwd+"@tcp("+sqladdress+")"+sqlother,
		maxIdle,
		maxConn)
	orm.RegisterModel(new(dao.User), new(dao.Skill), new(dao.Event), new(dao.Education), new(dao.Company), new(dao.Works), new(dao.Honor))
}
