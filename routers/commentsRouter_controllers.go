package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

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
			Method:           "Aprobar_Rechazar_solicitud",
			Router:           "/aprobar-rechazar",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SolicitudController"],
		beego.ControllerComments{
			Method:           "GetFormulariosByDocumentoId",
			Router:           "/formularios/:documentoId",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sabaticos_mid/controllers:SolicitudController"],
		beego.ControllerComments{
			Method:           "Radicar",
			Router:           "/radicar/:id",
			AllowHTTPMethods: []string{"post"},
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

}
