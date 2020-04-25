package dao

import (
	"github.com/astaxie/beego/orm"
	"log"
	"resume/models/dao/base"
	"time"
)

type Education struct {
	Code   string `json:"code"`
	School string `json:"school"`
	Formal string `json:"formal"`
	Start  int    `json:"start"`
	End    int    `json:"end"`
	Job    string `json:"job"`
	base.BaseData
}

func (edu Education) Insert() (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(&edu)
	if err != nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}
	return id, err
}

func (edu Education) Update() (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Update(&edu)
	return
}

func (edu Education) Delete(code string, deteleUser string) (int64, error) {
	o := orm.NewOrm()
	update, err := o.QueryTable(new(Education)).Filter("code", code).Update(orm.Params{
		"modifled_user": deteleUser,
		"delete_user":   deteleUser,
		"detele_time":   time.Now(),
		"delete":        1,
	})
	return update, err
}
func (edu Education) State(code string, modifledUser string, enable bool) error {
	o := orm.NewOrm()
	_, err := o.QueryTable(new(Education)).Filter("code", code).Update(orm.Params{
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

func SelectEducation(code string) (edu []Education, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("education").Filter("code", code).All(&edu)

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
