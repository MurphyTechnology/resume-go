package dao

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"resume/models/dao/base"
	"time"
)

type Event struct {
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	CompanyCode string    `json:"company"`
	StartTime   time.Time `json:"-"`
	EndTime     time.Time `json:"-"`
	Start       string    `json:"start_time" orm:"-"`
	End         string    `json:"end_time" orm:"-"`
	Describe    string    `json:"describe"`
	Harvest     string    `json:"harvest"`
	base.BaseData
}

func SelectEventAll(companyCode string) (event []Event) {
	o := orm.NewOrm()
	all, err := o.QueryTable(new(Event)).Filter("enable", true).Filter("delete", false).Filter("company_code", companyCode).All(&event)
	for i := 0; i < len(event); i++ {
		event[i].End = fmt.Sprintf(`%s`, event[i].EndTime.Format("2006/01"))
		event[i].Start = fmt.Sprintf(`%s`, event[i].StartTime.Format("2006/01"))
	}
	fmt.Println(all, err)
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
	if err != nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}
	return
}

func (event Event) Delete(companyCode string, deteleUser string) (update int64, err error) {
	o := orm.NewOrm()
	update, err = o.QueryTable("skill").Filter("company_code", companyCode).Update(orm.Params{
		"modifled_user": deteleUser,
		"delete_user":   deteleUser,
		"detele_time":   time.Now(),
		"delete":        1,
	})
	return
}
func (event Event) State(companyCode string, modifledUser string, enable bool) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("event").Filter("company_code", companyCode).Update(orm.Params{
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
