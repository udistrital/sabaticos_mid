package clients

import (
	"api_mid_sabaticos/enums"
	"api_mid_sabaticos/helpers"
	"api_mid_sabaticos/models"
	"errors"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

// Peticiones GET
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

// Peticiones POST
func RegistrarSolicitud(terceroId int, codigoTipoSolicitud string, sabaticoId *int) (*models.Solicitud, error) {
	var solicitudRes interface{}
	var solicitudCreada *models.Solicitud

	tipoSolicitud, err := ConsultarTipoSolicitud(codigoTipoSolicitud)
	if err != nil {
		return nil, err
	}

	solicitud := models.SolicitudCreateRequest{
		TerceroId: terceroId,
		Activo:    true,
		TipoSolicitudId: models.IdReference{
			Id: tipoSolicitud.Id,
		},
	}

	if err := request.SendJson(beego.AppConfig.String("sabaticosService")+"/solicitud/", "POST", &solicitudRes, solicitud); err != nil {
		beego.Error("Error POST solicitud:", err)
		return nil, err
	}

	if err := helpers.ExtractDataApi(solicitudRes, &solicitudCreada); err != nil {
		beego.Error("Error extrayendo datos de solicitud:", err)
		return nil, err
	}

	return solicitudCreada, nil
}

func RegistrarHistorialSolicitud(solicitudId int, terceroId int) (*models.HistorialSolicitud, error) {
	var historicoResp interface{}

	tipoSolicitud, err := ConsultarEstadoSolicitud(string(enums.ENVIADA))
	if err != nil {
		return nil, err
	}

	historial := models.HistorialSolicitudCreateRequest{
		TerceroId:     terceroId,
		Justificacion: "Nueva Solicitud Creada",
		Activo:        true,
		EstadoSolicitudId: models.IdReference{
			Id: tipoSolicitud.Id,
		},
		SolicitudId: models.IdReference{
			Id: solicitudId,
		},
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

func RegistrarFormularioSolicitud(solicitudId int, contenido string) (*models.FormularioSolicitud, error) {
	var formularioResp interface{}

	formulario := models.FormularioSolicitudCreateRequest{
		Contenido:   contenido,
		SolicitudId: models.IdReference{Id: solicitudId},
		Activo:      true,
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
