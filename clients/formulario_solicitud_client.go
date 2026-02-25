package clients

import (
	"api_mid_sabaticos/helpers"
	"api_mid_sabaticos/models"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

func RegistrarFormularioSolicitud(solicitudId int) (*models.FormularioSolicitud, error) {
	var formularioResp interface{}

	formulario := models.FormularioSolicitudCreateRequest{
		Contenido: "{}",
		SolicitudId: models.IdReference{
			Id: solicitudId,
		},
		Activo: true,
	}

	if err := request.SendJson(beego.AppConfig.String("sabaticosService")+"/formulario_solicitud/", "POST", &formularioResp, formulario); err != nil {
		beego.Error("Error POST formulario:", err)
	}

	var formularioFinal *models.FormularioSolicitud

	if err := helpers.ExtractDataApi(formularioResp, &formularioFinal); err != nil {
		beego.Error("Error extrayendo datos de formulario:", err)
		return nil, err
	}

	return formularioFinal, nil

}
