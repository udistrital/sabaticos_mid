package clients

import (
	"api_mid_sabaticos/enums"
	"api_mid_sabaticos/helpers"
	"api_mid_sabaticos/models"
	"errors"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

func ConsultarTipoSolicitud(codigo string) (*models.TipoSolicitud, error) {
	var tipoSolicitudRes interface{}
	var tipoSolicitud []models.TipoSolicitud

	codigoAbreviacion, ok := enums.ObtenerCodigoTipoSolicitud(codigo)
	if !ok {
		return nil, errors.New("tipo de solicitud no válido: " + codigo)
	}

	if err := request.GetJson(beego.AppConfig.String("sabaticosService")+"/tipo_solicitud?query=CodigoAbreviacion:"+codigoAbreviacion, &tipoSolicitudRes); err != nil {
		return nil, err
	}
	if err := helpers.ExtractDataApi(tipoSolicitudRes, &tipoSolicitud); err != nil {
		return nil, err

	}
	if len(tipoSolicitud) == 0 {
		return nil, errors.New("tipo de solicitud no encontrado: " + codigo)
	}

	return &tipoSolicitud[0], nil
}
