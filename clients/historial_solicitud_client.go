package clients

import (
	"api_mid_sabaticos/enums"
	"api_mid_sabaticos/models"

	"api_mid_sabaticos/helpers"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

func RegistrarHistorialSolicitud(solicitudId int, terceroId int) (*models.HistorialSolicitud, error) {
	var historicoResp interface{}

	tipoSolicitud, err := ConsultarEstadoSolicitud(string(enums.ENVIADA))
	if err != nil {
		return nil, err
	}

	historial := models.HistorialSolicitud{
		TerceroId:         terceroId,
		Justificacion:     "Nueva Solicitud Creada",
		Activo:            true,
		EstadoSolicitudId: tipoSolicitud.Id,
		SolicitudId:       solicitudId,
	}

	if err := request.SendJson(beego.AppConfig.String("sabaticosService")+"/historial_solicitud/", "POST", &historicoResp, historial); err != nil {
		beego.Error("Error POST histórico:", err)
	}

	var historialFinal *models.HistorialSolicitud

	if err := helpers.ExtractDataApi(historicoResp, &historialFinal); err != nil {
		beego.Error("Error extrayendo datos de histórico:", err)
		return nil, err
	}

	return historialFinal, nil
}
