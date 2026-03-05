package main

import (
	_ "api_mid_sabaticos/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/udistrital/utils_oas/errorhandler"
	"github.com/udistrital/utils_oas/xray"
)

// CORS middleware
var CORS = func(ctx *context.Context) {
	ctx.Output.Header("Access-Control-Allow-Origin", "*")
	ctx.Output.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
	ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type,Authorization,Accept,X-Requested-With")
	ctx.Output.Header("Access-Control-Allow-Credentials", "true")

	if ctx.Input.Method() == "OPTIONS" {
		ctx.Output.SetStatus(200)
		return
	}
}

func main() {
	xray.InitXRay()

	// Insertar filtro CORS
	beego.InsertFilter("*", beego.BeforeRouter, CORS)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.ErrorController(&errorhandler.ErrorHandlerController{})

	beego.Run()
}
