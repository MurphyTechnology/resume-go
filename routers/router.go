// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego/plugins/cors"
	"resume/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		//AllowOrigins:      []string{"https://192.168.0.102"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	ns := beego.NewNamespace("/v1",
		//beego.NSNamespace("/object",
		//	beego.NSInclude(
		//		&controllers.ObjectController{},
		//	),
		//),
		beego.NSNamespace("/resume",
			beego.NSInclude(
				&controllers.ResumeController{},
			),
		),
		//beego.NSNamespace("/user",
		//	beego.NSInclude(
		//		&controllers.UserController{},
		//	),
		//),
	)
	beego.AddNamespace(ns)
}
