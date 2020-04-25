package dao

import (
	"github.com/astaxie/beego/orm"
	"log"
	"resume/models/dao/base"
	"time"
)

type User struct {
	Code     string    `json:"code"`
	Name     string    `json:"name"`
	Phone    string    `json:"phone"`
	Describe string    `json:"describe"`
	Sex      bool      `json:"sex"`
	Status   string    `json:"status"`
	Birthday time.Time `json:"birthday"`
	Photo    string    `json:"photo"`
	Email    string    `json:"email"`
	Wechat   string    `json:"wechat"`
	Github   string    `json:"github"`
	Native   string    `json:"native"`
	Keyword  string    `json:"keyword"`
	Watch    int       `json:"watch"`
	City     string    `json:"city"`
	Expect   string    `json:"expect"`
	Year     string    `json:"year"`
	base.BaseData
}

func SelectUserByKeyword(keyword string) (user User, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("user").Filter("keyword", keyword).One(&user)
	user.Watch += 1
	o.Update(&user)
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
func SelectUserByCode(code string) (user User, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("user").Filter("code", code).One(&user)
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

func UpdateKeyword(code string, keyword string) (update int64, err error) {
	o := orm.NewOrm()
	update, err = o.QueryTable("user").Filter("code", code).Update(orm.Params{
		"keyword": keyword,
	})
	return
}

func (user User) State(code string, modifledUser string, enable bool) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("user").Filter("code", code).Update(orm.Params{
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

func (user User) Insert() (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(&user)
	return id, err
}

func (user User) Delete(code string, deteleUser string) (int64, error) {
	o := orm.NewOrm()
	update, err := o.QueryTable(new(User)).Filter("code", code).Update(orm.Params{
		"modifled_user": deteleUser,
		"delete_user":   deteleUser,
		"detele_time":   time.Now(),
		"delete":        1,
	})
	return update, err
}

func (user User) Update() (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Update(&user)
	if err != nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}
	return
}
