package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	"resume/models/dao"
	"resume/models/response"
	"resume/service"
)

type ResumeController struct {
	beego.Controller
}

var r *response.Result = new(response.Result)

// @Success 200 {string} logout success
// @router /getInfoByCode [get]
func (this *ResumeController) Select() {
	code := this.GetString("code")
	if code != "" {
		resume, err := service.Select(code)
		if err != nil {
			this.Data["json"] = r.Fail()
		} else {
			this.Data["json"] = r.Success(resume)
		}
	}
	this.ServeJSON()
}

// @Success 200 {string} logout success
// @router /getInfoBykeyWord [get]
func (this *ResumeController) SelectKeyword() {
	keyword := this.GetString("keyword")
	if keyword != "" {
		resume, err := service.SelectKeyword(keyword)
		if err != nil {
			this.Data["json"] = r.Fail()
		} else {
			this.Data["json"] = r.Success(resume)
		}
	}
	this.ServeJSON()
}

// @Success 200 {string} logout success
// @router /createKeyword [post]
func (this *ResumeController) CreateKeyword() {
	data := this.Ctx.Input.RequestBody
	m := make(map[string]string)
	err := json.Unmarshal([]byte(data), &m)
	if err == nil {
		display := service.Display(m["code"])
		this.Data["json"] = r.Success(display)
	} else {
		this.Data["json"] = r.Fail()
	}
	this.ServeJSON()
}

// @Success 200 {string} logout success
// @router /updateUser [post]
func (this *ResumeController) UpdateUser() {
	var user dao.User
	data := this.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &user)
	if err == nil {
		update, err := service.UpdateUser(user)
		if err != nil {
			this.Data["json"] = r.Fail()
		} else {
			this.Data["json"] = r.Success(update)
		}
	}
	this.ServeJSON()
}

// @Success 200 {string} logout success
// @router /addUser [post]
func (this *ResumeController) AddUser() {
	var user dao.User
	data := this.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &user)
	if err == nil {
		insert, err := service.AddUser(user)
		if err != nil {
			this.Data["json"] = r.Fail()
		} else {
			this.Data["json"] = r.Success(insert)
		}
	}
	this.ServeJSON()
}

// @Success 200 {string} logout success
// @router /addEvent [post]
func (this *ResumeController) AddEvent() {
	var event dao.Event
	data := this.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &event)
	if err == nil {
		insert, err := service.AddEvent(event)
		if err != nil {
			this.Data["json"] = r.Fail()
		} else {
			this.Data["json"] = r.Success(insert)
		}
	}
	this.ServeJSON()
}

// @Success 200 {string} logout success
// @router /UpdateEvent [post]
func (this *ResumeController) UpdateEvent() {
	var event dao.Event
	data := this.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &event)
	if err == nil {
		update, err := service.UpdateEvent(event)
		if err != nil {
			this.Data["json"] = r.Fail()
		} else {
			this.Data["json"] = r.Success(update)
		}
	}
	this.ServeJSON()
}

// @Success 200 {string} logout success
// @router /delEvent [post]
func (this *ResumeController) DelEvent() {
	data := this.Ctx.Input.RequestBody
	m := make(map[string]string)
	err := json.Unmarshal([]byte(data), &m)
	if err == nil {
		event, err := service.DeleteEvent(m["code"])
		if err != nil {
			this.Data["json"] = r.Fail()
		} else {
			this.Data["json"] = r.Success(event)
		}
	}
	this.ServeJSON()
}

// @Success 200 {string} logout success
// @router /enbleEvent [post]
func (this *ResumeController) EnbleEvent() {
	data := this.Ctx.Input.RequestBody
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(data), &m)
	var c string
	var e bool
	if c, ok := m["code"].(string); ok {
		log.Fatal(c)
	}
	if e, ok2 := m["enable"].(bool); ok2 {
		log.Fatal(e)
	}
	if err == nil {

		err := service.EnableSkill(c, e)
		if err != nil {
			this.Data["json"] = r.Fail()
		} else {
			this.Data["json"] = r.Success("ok")
		}
	}
	this.ServeJSON()
}

// @Success 200 {string} logout success
// @router /UpdateSkill [post]
func (this *ResumeController) UpdateSkill() {
	var skill dao.Skill
	data := this.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &skill)
	if err == nil {
		update, err := service.UpdateSkill(skill)
		if err != nil {
			this.Data["json"] = r.Fail()
		} else {
			this.Data["json"] = r.Success(update)
		}
	}
	this.ServeJSON()
}

// @Success 200 {string} logout success
// @router /addSkill [post]
func (this *ResumeController) AddSkill() {
	var skill dao.Skill
	data := this.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &skill)
	if err == nil {
		update, err := service.AddSkill(skill)
		if err != nil {
			this.Data["json"] = r.Fail()
		} else {
			this.Data["json"] = r.Success(update)
		}
	}
	this.ServeJSON()
}

// @Success 200 {string} logout success
// @router /enableSkill [post]
func (this *ResumeController) EnableSkill() {
	data := this.Ctx.Input.RequestBody
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(data), &m)
	var c string
	var e bool
	if c, ok := m["code"].(string); ok {
		log.Fatal(c)
	}
	if e, ok2 := m["enable"].(bool); ok2 {
		log.Fatal(e)
	}
	if err == nil {
		err := service.EnableSkill(c, e)
		if err != nil {
			this.Data["json"] = r.Fail()
		} else {
			this.Data["json"] = r.Success("ok")
		}
	}
	this.ServeJSON()
}

// @Success 200 {string} logout success
// @router /delSkill [post]
func (this *ResumeController) DelSkill() {
	data := this.Ctx.Input.RequestBody
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(data), &m)
	var c string
	if c, ok := m["code"].(string); ok {
		log.Fatal(c)
	}
	if err == nil {
		skill, err := service.DeleteSkill(c)
		if err != nil {
			this.Data["json"] = r.Fail()
		} else {
			this.Data["json"] = r.Success(skill)
		}
	}
	this.ServeJSON()
}
