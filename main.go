package main

import (
	_ "api_mid_sabaticos/routers"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/errorhandler"
	"github.com/udistrital/utils_oas/xray"
)

func main() {
	xray.InitXRay()

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.ErrorController(&errorhandler.ErrorHandlerController{})

	beego.Run()
}
