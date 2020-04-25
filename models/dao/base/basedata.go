package base

import (
	"time"
)

type BaseData struct {
	Base
	CreateTime   time.Time `orm:"auto_now_add;type(date)" json:"-"`
	DeleteTime   time.Time `json:"-"`
	ModifledTime time.Time `orm:"auto_now;type(date)" json:"-"`
	CreateUser   string    `json:"-"`
	DeleteUser   string    `json:"-"`
	ModifledUser string    `json:"-"`
	Delete       bool      `orm:"default(0)" json:"-"`
	Enable       bool      `orm:"default(1)" json:"-"`
}
