package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SabaticoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SabaticoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SabaticoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SabaticoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SabaticoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SabaticoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SabaticoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SabaticoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SabaticoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SabaticoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SolicitudController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SolicitudController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SolicitudController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SolicitudController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SolicitudController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SoporteSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SoporteSolicitudController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SoporteSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SoporteSolicitudController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SoporteSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SoporteSolicitudController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SoporteSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SoporteSolicitudController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SoporteSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SoporteSolicitudController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/:uid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:uid",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:uid",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           "/login",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Logout",
			Router:           "/logout",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
