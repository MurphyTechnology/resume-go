package main

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "resume/db"
	_ "resume/routers"
	"resume/service"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		//	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	fmt.Println(service.Select("w"))
	beego.Run()
}
