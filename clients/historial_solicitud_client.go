package clients

import (
	"api_mid_sabaticos/enums"
	"api_mid_sabaticos/models"

	"api_mid_sabaticos/helpers"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

func RegistrarHistorialSolicitud(historialReq models.HistorialSolicitud, solicitudCreada *models.Solicitud, solicitudReq models.SolicitudRequest) (*models.HistorialSolicitud, error) {
	var historicoResp interface{}

	historial := models.HistorialSolicitud{
		TerceroId:         solicitudReq.TerceroId,
		Justificacion:     "Nueva Solicitud Creada",
		Activo:            true,
		EstadoSolicitudId: int(enums.ENVIADA),
		SolicitudId:       solicitudCreada.Id,
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
