package dao

import (
	"github.com/astaxie/beego/orm"
	"resume/models/dao/base"
	"time"
)

type Event struct {
	Code      string        `json:"code"`
	Name      string        `json:"name"`
	Company   string        `json:"company"`
	StartTime orm.DateField `json:"start_time"`
	EndTime   orm.DateField `json:"end_time"`
	describe  string        `json:"describe"`
	harvest   string        `json:"harvest"`
	base.BaseData
}

func SelectEventAll(code string) (event []Event) {
	o := orm.NewOrm()
	o.QueryTable(new(Event)).Filter("enable", true).Filter("delete", false).Filter("code", code).All(&event)
	return
}

func (event Event) Insert() (insert int64, err error) {
	o := orm.NewOrm()
	insert, err = o.Insert(&event)
	return
}

func (event Event) Update() (update int64, err error) {
	o := orm.NewOrm()
	update, err = o.Update(&event)
	return
}

func DeleteEvent(code string, deteleUser string) (update int64, err error) {
	o := orm.NewOrm()
	update, err = o.QueryTable("skill").Filter("code", code).Update(orm.Params{
		"modifled_user": deteleUser,
		"delete_user":   deteleUser,
		"detele_time":   time.Now(),
		"delete":        1,
	})
	return
}
func StateEvent(code string, modifledUser string, enable bool) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("event").Filter("code", code).Update(orm.Params{
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
