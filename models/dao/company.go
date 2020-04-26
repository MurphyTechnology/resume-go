package dao

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"log"
	"resume/models/dao/base"
	"time"
)

type Company struct {
	Code        string    `json:"code"`
	CompanyName string    `json:"company_name"`
	CompanyCode string    `json:"company_code"`
	Start       time.Time `json:"-"`
	End         time.Time `json:"-"`
	StartYM     string    `json:"start" orm:"-"`
	EndYM       string    `json:"end" orm:"-"`
	base.BaseData
}

func (com Company) Insert() (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(&com)
	return id, err
}

func (com Company) Update() (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Update(&com)
	if err != nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}
	return
}

func (com Company) Delete(code string, deteleUser string) (int64, error) {
	o := orm.NewOrm()
	update, err := o.QueryTable(new(Company)).Filter("code", code).Update(orm.Params{
		"modifled_user": deteleUser,
		"delete_user":   deteleUser,
		"detele_time":   time.Now(),
		"delete":        1,
	})
	return update, err
}
func (com Company) State(code string, modifledUser string, enable bool) error {
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

func SelectCompanyByCode(code string) (com []Company, err error) {
	o := orm.NewOrm()
	o.QueryTable("company").Filter("enable", true).Filter("delete", false).Filter("code", code).All(&com)
	for i := 0; i < len(com); i++ {
		com[i].EndYM = fmt.Sprintf(`%s`, com[i].End.Format("2006/01"))
		com[i].StartYM = fmt.Sprintf(`%s`, com[i].Start.Format("2006/01"))
	}
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
