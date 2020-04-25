package dao

import (
	"github.com/astaxie/beego/orm"
	"log"
	"resume/models/dao/base"
	"time"
)

type Works struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Url   string `json:"url"`
	Photo string `json:"photo"`
	base.BaseData
}

func (w Works) Insert() (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(&w)
	return id, err
}

func (w Works) Update() (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Update(&w)
	if err != nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}
	return
}

func (w Works) Delete(code string, deteleUser string) (int64, error) {
	o := orm.NewOrm()
	update, err := o.QueryTable(new(Works)).Filter("code", code).Update(orm.Params{
		"modifled_user": deteleUser,
		"delete_user":   deteleUser,
		"detele_time":   time.Now(),
		"delete":        1,
	})
	return update, err
}
func (w Works) State(code string, modifledUser string, enable bool) error {
	o := orm.NewOrm()
	_, err := o.QueryTable(new(Company)).Filter("code", code).Update(orm.Params{
		"enable":        enable,
		"modifled_user": modifledUser,
	})
	if err != nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}
	return err
}

func SelectWorksByCode(code string) (w []Works, err error) {
	o := orm.NewOrm()
	o.QueryTable("works").Filter("enable", true).Filter("delete", false).Filter("code", code).All(&w)
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		log.Printf("Returned Multi Rows Not One")
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		log.Printf("Not row found")
	}
	return
}
