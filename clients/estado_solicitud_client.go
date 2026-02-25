package clients

import (
	"api_mid_sabaticos/enums"
	"api_mid_sabaticos/helpers"
	"api_mid_sabaticos/models"
	"errors"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

func ConsultarEstadoSolicitud(codigo string) (*models.EstadoSolicitud, error) {
	var estadoSolicitudRes interface{}
	var estadoSolicitud []models.EstadoSolicitud

	codigoAbreviacion, ok := enums.ObtenerCodigoEstadoSolicitud(codigo)
	if !ok {
		return nil, errors.New("estado de solicitud no válido: " + codigo)
	}

	if err := request.GetJson(beego.AppConfig.String("sabaticosService")+"/estado_solicitud?query=CodigoAbreviacion:"+codigoAbreviacion, &estadoSolicitudRes); err != nil {
		return nil, err
	}
	if err := helpers.ExtractDataApi(estadoSolicitudRes, &estadoSolicitud); err != nil {
		return nil, err

	}
	if len(estadoSolicitud) == 0 {
		return nil, errors.New("estado de solicitud no encontrado: " + codigo)
	}

	return &estadoSolicitud[0], nil
}
