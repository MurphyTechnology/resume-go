package dao

import (
	"github.com/astaxie/beego/orm"
	"resume/models/dao/base"
	"time"
)

/**
技能
*/
type Skill struct {
	//逻辑code
	Code string `json:"code"`
	//技能名称
	Name string `json:"name"`
	//排序
	Sequence int `json:"sequence" orm:digits(2)`
	//技能掌握百分比
	Percentage int `json:"percentage" orm:digits(3)`
	base.BaseData
}

func SelectSkillAll(code string) (skills []Skill) {
	o := orm.NewOrm()
	o.QueryTable(new(Skill)).Filter("enable", true).Filter("delete", false).Filter("code", code).All(&skills)
	return
}

func (skill Skill) Insert() (insert int64, err error) {
	o := orm.NewOrm()
	insert, err = o.Insert(&skill)
	return
}

func (skill Skill) Update() (update int64, err error) {
	o := orm.NewOrm()
	update, err = o.Update(&skill)
	if err != nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}
	return
}

func (skill Skill) Delete(code string, deteleUser string) (int64, error) {
	o := orm.NewOrm()
	update, err := o.QueryTable(new(Skill)).Filter("code", code).Update(orm.Params{
		"modifled_user": deteleUser,
		"delete_user":   deteleUser,
		"detele_time":   time.Now(),
		"delete":        1,
	})
	return update, err
}
func (skill Skill) State(code string, modifledUser string, enable bool) error {
	o := orm.NewOrm()
	_, err := o.QueryTable(new(Skill)).Filter("code", code).Update(orm.Params{
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
