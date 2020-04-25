package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {
	beego.GlobalControllerRouter["resume/controllers:ResumeController"] =
		append(beego.GlobalControllerRouter["resume/controllers:ResumeController"],
			beego.ControllerComments{
				Method:           "Select",
				Router:           `/getInfoByCode`,
				AllowHTTPMethods: []string{"get"},
				MethodParams:     param.Make(),
				Filters:          nil,
				Params:           nil})
	beego.GlobalControllerRouter["resume/controllers:ResumeController"] =
		append(beego.GlobalControllerRouter["resume/controllers:ResumeController"],
			beego.ControllerComments{
				Method:           "SelectKeyword",
				Router:           `/getInfoByKeyword`,
				AllowHTTPMethods: []string{"get"},
				MethodParams:     param.Make(),
				Filters:          nil,
				Params:           nil})
	beego.GlobalControllerRouter["resume/controllers:ResumeController"] =
		append(beego.GlobalControllerRouter["resume/controllers:ResumeController"],
			beego.ControllerComments{
				Method:           "CreateKeyword",
				Router:           `/createKeyword`,
				AllowHTTPMethods: []string{"post"},
				MethodParams:     param.Make(),
				Filters:          nil,
				Params:           nil})
	beego.GlobalControllerRouter["resume/controllers:ResumeController"] =
		append(beego.GlobalControllerRouter["resume/controllers:ResumeController"],
			beego.ControllerComments{
				Method:           "UpdateUser",
				Router:           `/updateUser`,
				AllowHTTPMethods: []string{"post"},
				MethodParams:     param.Make(),
				Filters:          nil,
				Params:           nil})
	beego.GlobalControllerRouter["resume/controllers:ResumeController"] =
		append(beego.GlobalControllerRouter["resume/controllers:ResumeController"],
			beego.ControllerComments{
				Method:           "AddUser",
				Router:           `/addUser`,
				AllowHTTPMethods: []string{"post"},
				MethodParams:     param.Make(),
				Filters:          nil,
				Params:           nil})
	beego.GlobalControllerRouter["resume/controllers:ResumeController"] =
		append(beego.GlobalControllerRouter["resume/controllers:ResumeController"],
			beego.ControllerComments{
				Method:           "AddEvent",
				Router:           `/addEvent`,
				AllowHTTPMethods: []string{"post"},
				MethodParams:     param.Make(),
				Filters:          nil,
				Params:           nil})
	beego.GlobalControllerRouter["resume/controllers:ResumeController"] =
		append(beego.GlobalControllerRouter["resume/controllers:ResumeController"],
			beego.ControllerComments{
				Method:           "UpdateEvent",
				Router:           `/updateEvent`,
				AllowHTTPMethods: []string{"post"},
				MethodParams:     param.Make(),
				Filters:          nil,
				Params:           nil})
	beego.GlobalControllerRouter["resume/controllers:ResumeController"] =
		append(beego.GlobalControllerRouter["resume/controllers:ResumeController"],
			beego.ControllerComments{
				Method:           "DelEvent",
				Router:           `/delEvent`,
				AllowHTTPMethods: []string{"post"},
				MethodParams:     param.Make(),
				Filters:          nil,
				Params:           nil})
	beego.GlobalControllerRouter["resume/controllers:ResumeController"] =
		append(beego.GlobalControllerRouter["resume/controllers:ResumeController"],
			beego.ControllerComments{
				Method:           "EnbleEvent",
				Router:           `/enbleEvent`,
				AllowHTTPMethods: []string{"post"},
				MethodParams:     param.Make(),
				Filters:          nil,
				Params:           nil})
	beego.GlobalControllerRouter["resume/controllers:ResumeController"] =
		append(beego.GlobalControllerRouter["resume/controllers:ResumeController"],
			beego.ControllerComments{
				Method:           "UpdateSkill",
				Router:           `/updateSkill`,
				AllowHTTPMethods: []string{"post"},
				MethodParams:     param.Make(),
				Filters:          nil,
				Params:           nil})
	beego.GlobalControllerRouter["resume/controllers:ResumeController"] =
		append(beego.GlobalControllerRouter["resume/controllers:ResumeController"],
			beego.ControllerComments{
				Method:           "AddSkill",
				Router:           `/addSkill`,
				AllowHTTPMethods: []string{"post"},
				MethodParams:     param.Make(),
				Filters:          nil,
				Params:           nil})
	beego.GlobalControllerRouter["resume/controllers:ResumeController"] =
		append(beego.GlobalControllerRouter["resume/controllers:ResumeController"],
			beego.ControllerComments{
				Method:           "EnableSkill",
				Router:           `/enableSkill`,
				AllowHTTPMethods: []string{"post"},
				MethodParams:     param.Make(),
				Filters:          nil,
				Params:           nil})
	beego.GlobalControllerRouter["resume/controllers:ResumeController"] =
		append(beego.GlobalControllerRouter["resume/controllers:ResumeController"],
			beego.ControllerComments{
				Method:           "DelSkill",
				Router:           `/delSkill`,
				AllowHTTPMethods: []string{"post"},
				MethodParams:     param.Make(),
				Filters:          nil,
				Params:           nil})
	beego.GlobalControllerRouter["resume/controllers:ResumeController"] =
		append(beego.GlobalControllerRouter["resume/controllers:ResumeController"],
			beego.ControllerComments{
				Method:           "CreateQrcode",
				Router:           `/createQrcode`,
				AllowHTTPMethods: []string{"post"},
				MethodParams:     param.Make(),
				Filters:          nil,
				Params:           nil})
}
