package dao

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"resume/models/dao/base"
	"time"
)

type Honor struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Url  string `json:"url"`
	base.BaseData
}

func SelectHonorAll(code string) (honor []Honor) {
	o := orm.NewOrm()
	all, err := o.QueryTable(new(Honor)).Filter("enable", true).Filter("delete", false).Filter("code", code).All(&honor)
	fmt.Println(all, err)
	return
}

func (honor Honor) Insert() (insert int64, err error) {
	o := orm.NewOrm()
	insert, err = o.Insert(&honor)
	return
}

func (honor Honor) Update() (update int64, err error) {
	o := orm.NewOrm()
	update, err = o.Update(&honor)
	if err != nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}
	return
}

func (honor Honor) Delete(code string, deteleUser string) (update int64, err error) {
	o := orm.NewOrm()
	update, err = o.QueryTable(new(Honor)).Filter("code", code).Update(orm.Params{
		"modifled_user": deteleUser,
		"delete_user":   deteleUser,
		"detele_time":   time.Now(),
		"delete":        1,
	})
	return
}
func (honor Honor) State(code string, modifledUser string, enable bool) error {
	o := orm.NewOrm()
	_, err := o.QueryTable(new(Honor)).Filter("code", code).Update(orm.Params{
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
