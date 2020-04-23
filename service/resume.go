package service

import (
	"resume/models/dao"
	"resume/models/response"
	"resume/utils"
)

func AddUser(user dao.User) (int64, error) {
	user.CreateUser = "用户"
	insert, err := user.Insert()
	return insert, err
}

func AddSkill(skill dao.Skill) (int64, error) {
	skill.CreateUser = "用户"
	insert, err := skill.Insert()
	return insert, err
}

func AddEvent(event dao.Event) (int64, error) {
	event.CreateUser = "用户"
	insert, err := event.Insert()
	return insert, err
}

func Display(code string) string {
	uid := utils.GetRandomSalt()
	dao.UpdateKeyword(code, uid)
	return uid
}

func UpdateUser(user dao.User) (int64, error) {
	user.ModifledUser = "用户"
	update, err := user.Update()
	return update, err
}
func UpdateEvent(event dao.Event) (int64, error) {
	event.ModifledUser = "用户"
	update, err := event.Update()
	return update, err
}
func UpdateSkill(skill dao.Skill) (int64, error) {
	skill.ModifledUser = "用户"
	update, err := skill.Update()
	return update, err
}

func DeleteSkill(code string) (int64, error) {
	skill, err := dao.DeleteSkill(code, "用户")
	return skill, err
}

func DeleteEvent(code string) (int64, error) {
	event, err := dao.DeleteEvent(code, "用户")
	return event, err
}

func EnableSkill(code string, enable bool) error {
	err := dao.StateSkill(code, "用户", enable)
	return err
}

func EnableEvent(code string, modifledUser string, enable bool) {
	dao.StateEvent(code, "用户", enable)
}

func SelectKeyword(keyword string) (response.Resume, error) {
	user, err := dao.SelectUserByKeyword(keyword)
	skillAll := dao.SelectSkillAll(user.Code)
	eventAll := dao.SelectEventAll(user.Code)
	for i := 0; i < len(skillAll); i++ {
		skillAll[i].Code = ""
	}
	for i := 0; i < len(eventAll); i++ {
		eventAll[i].Code = ""
	}
	user.Code = ""
	return response.Resume{User: user, Skill: skillAll, Event: eventAll}, err
}
func Select(code string) (response.Resume, error) {
	user, err := dao.SelectUserByCode(code)
	skillAll := dao.SelectSkillAll(code)
	eventAll := dao.SelectEventAll(code)
	return response.Resume{User: user, Skill: skillAll, Event: eventAll}, err
}
