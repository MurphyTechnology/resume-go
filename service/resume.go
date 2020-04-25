package service

import (
	"bytes"
	"encoding/base64"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/png"
	"resume/models/dao"
	"resume/models/response/resume"
	"resume/utils"
	"unsafe"
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
	var skill dao.Skill
	i, err := skill.Delete(code, "用户")
	return i, err
}

func DeleteEvent(code string) (int64, error) {
	var event dao.Event
	i, err := event.Delete(code, "用户")
	return i, err
}

func EnableSkill(code string, enable bool) error {
	var skill dao.Skill
	err := skill.State(code, "用户", enable)
	return err
}

func EnableEvent(code string, modifledUser string, enable bool) {
	var skill dao.Skill
	skill.State(code, "用户", enable)
}

func SelectKeyword(keyword string) (resume.Resume, error) {
	user, err := dao.SelectUserByKeyword(keyword)
	skillAll := dao.SelectSkillAll(user.Code)
	edu, _ := dao.SelectEducation(user.Code)
	company, _ := dao.SelectCompanyByCode(user.Code)
	works, err := dao.SelectWorksByCode(user.Code)
	honor := dao.SelectHonorAll(user.Code)
	for i := 0; i < len(works); i++ {
		works[i].Code = ""
	}

	var com = make([]resume.CompanyRes, 0)
	for i := 0; i < len(company); i++ {
		var comItem resume.CompanyRes
		all := dao.SelectEventAll(company[i].CompanyCode)
		for i := 0; i < len(all); i++ {
			all[i].Code = ""
		}
		company[i].Code = ""
		comItem.Event = all
		comItem.Info = company[i]
		com = append(com, comItem)
	}

	for i := 0; i < len(skillAll); i++ {
		skillAll[i].Code = ""
	}
	user.Code = ""
	return resume.Resume{User: user, Skill: skillAll, Company: com, Education: edu, Works: works, Honor: honor}, err
}
func Select(code string) (resume.Resume, error) {
	user, err := dao.SelectUserByCode(code)
	skillAll := dao.SelectSkillAll(code)
	edu, _ := dao.SelectEducation(user.Code)
	company, _ := dao.SelectCompanyByCode(user.Code)
	works, err := dao.SelectWorksByCode(user.Code)
	honor := dao.SelectHonorAll(user.Code)
	var com = make([]resume.CompanyRes, 0)
	for i := 0; i < len(company); i++ {
		var comItem resume.CompanyRes
		all := dao.SelectEventAll(company[i].CompanyCode)
		comItem.Event = all
		comItem.Info = company[i]
		com = append(com, comItem)
	}
	return resume.Resume{User: user, Skill: skillAll, Company: com, Education: edu, Works: works, Honor: honor}, err
}

func CreateQrcode(url string) (string, error) {
	// Create the barcode
	qrCode, err := qr.Encode(url, qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, err = barcode.Scale(qrCode, 200, 200)
	emptyBuff := bytes.NewBuffer(nil)
	png.Encode(emptyBuff, qrCode)
	dist := make([]byte, 50000)
	base64.StdEncoding.Encode(dist, emptyBuff.Bytes())
	index := bytes.IndexByte(dist, 0)
	baseImage := dist[0:index]
	return *(*string)(unsafe.Pointer(&baseImage)), err
}
