package base

import "github.com/astaxie/beego/orm"

type BaseData struct {
	Base
	CreateTime   orm.DateTimeField `orm:"auto_now_add;type(date)" json:"-"`
	DeleteTime   orm.DateTimeField `json:"-"`
	ModifledTime orm.DateTimeField `orm:"auto_now;type(date)" json:"-"`
	CreateUser   string            `json:"-"`
	DeleteUser   string            `json:"-"`
	ModifledUser string            `json:"-"`
	Delete       bool              `orm:"default(0)" json:"-"`
	Enable       bool              `orm:"default(1)" json:"-"`
}
