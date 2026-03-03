package clients

import (
	"errors"
	"strings"

	"api_mid_sabaticos/enums"
	"api_mid_sabaticos/helpers"
	"api_mid_sabaticos/models"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

type relId struct {
	Id int `json:"Id"`
}

type historialSolicitudCRUD struct {
	Activo            bool   `json:"Activo"`
	Justificacion     string `json:"Justificacion"`
	TerceroId         int    `json:"TerceroId"`
	EstadoSolicitudId relId  `json:"EstadoSolicitudId"`
	SolicitudId       relId  `json:"SolicitudId"`
}

func RegistrarHistorialSolicitud(historialReq models.HistorialSolicitud, solicitudCreada *models.Solicitud, solicitudReq models.SolicitudRequest) (*models.HistorialSolicitud, error) {
	var historicoResp interface{}

	estado := historialReq.EstadoSolicitudId
	if estado <= 0 {
		estado = int(enums.ENVIADA)
	}

	justificacion := historialReq.Justificacion
	if justificacion == "" {
		justificacion = "Nueva Solicitud Creada"
	}

	payload := historialSolicitudCRUD{
		TerceroId:         solicitudReq.TerceroId,
		Justificacion:     justificacion,
		Activo:            true,
		EstadoSolicitudId: relId{Id: estado},
		SolicitudId:       relId{Id: solicitudCreada.Id},
	}

	base := strings.TrimRight(beego.AppConfig.String("sabaticosService"), "/")
	if base == "" {
		return nil, errors.New("sabaticosService está vacío")
	}

	if err := request.SendJson(base+"/v1/historial_solicitud/", "POST", &historicoResp, payload); err != nil {
		beego.Error("Error POST histórico:", err)
		return nil, err
	}

	var historialFinal *models.HistorialSolicitud
	if err := helpers.ExtractDataApi(historicoResp, &historialFinal); err != nil {
		beego.Error("Error extrayendo datos de histórico:", err)
		return nil, err
	}
	return historialFinal, nil
}
