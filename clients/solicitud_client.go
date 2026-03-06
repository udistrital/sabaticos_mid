package clients

import (
	"errors"
	"fmt"
	"github.com/udistrital/sabaticos_mid/enums"
	"github.com/udistrital/sabaticos_mid/helpers"
	"github.com/udistrital/sabaticos_mid/models"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

// Peticiones GET
func ConsultarSolicitud(id int) (*models.Solicitud, error) {
	var solicitudRes interface{}
	var solicitud models.Solicitud

	if err := request.GetJson(beego.AppConfig.String("sabaticosService")+"/solicitud/"+fmt.Sprint(id), &solicitudRes); err != nil {
		return nil, err
	}

	if err := helpers.ExtractDataApi(solicitudRes, &solicitud); err != nil {
		return nil, err

	}

	if solicitud.Id == 0 {
		return nil, errors.New("Solicitud no Encontrada: " + fmt.Sprint(id))
	}

	return &solicitud, nil
}

func ConsultarEstadoSolicitud(codigo string) (*models.EstadoSolicitud, error) {
	var estadoSolicitudRes interface{}
	var estadoSolicitud []models.EstadoSolicitud

	codigoAbreviacion, ok := enums.ObtenerCodigoEstadoSolicitud(codigo)
	if !ok {
		return nil, errors.New("Estado de Solicitud no Válido: " + codigo)
	}

	if err := request.GetJson(beego.AppConfig.String("sabaticosService")+"/estado_solicitud?query=CodigoAbreviacion:"+codigoAbreviacion, &estadoSolicitudRes); err != nil {
		return nil, err
	}

	if err := helpers.ExtractDataApi(estadoSolicitudRes, &estadoSolicitud); err != nil {
		return nil, err

	}

	if len(estadoSolicitud) == 0 {
		return nil, errors.New("Estado de Solicitud no Encontrado: " + codigo)
	}

	return &estadoSolicitud[0], nil
}

func ConsultarTipoSolicitud(codigo string) (*models.TipoSolicitud, error) {
	var tipoSolicitudRes interface{}
	var tipoSolicitud []models.TipoSolicitud

	codigoAbreviacion, ok := enums.ObtenerCodigoTipoSolicitud(codigo)
	if !ok {
		return nil, errors.New("Tipo de Solicitud no Válido: " + codigo)
	}

	if err := request.GetJson(beego.AppConfig.String("sabaticosService")+"/tipo_solicitud?query=CodigoAbreviacion:"+codigoAbreviacion, &tipoSolicitudRes); err != nil {
		return nil, err
	}
	if err := helpers.ExtractDataApi(tipoSolicitudRes, &tipoSolicitud); err != nil {
		return nil, err

	}
	if len(tipoSolicitud) == 0 {
		return nil, errors.New("Tipo de Solicitud no Encontrado: " + codigo)
	}

	return &tipoSolicitud[0], nil
}

func ConsultarEstadoSoporteSolicitud(codigo string) (*models.EstadoSoporteSolicitud, error) {
	var estadoSoporteSolicitudRes interface{}
	var estadoSoporteSolicitud []models.EstadoSoporteSolicitud

	codigoAbreviacion, ok := enums.ObtenerCodigoEstadoSoporteSolicitud(codigo)

	if !ok {
		return nil, errors.New("Estado de Soporte de Solicitud no Válido: " + codigo)
	}

	if err := request.GetJson(beego.AppConfig.String("sabaticosService")+"/estado_soporte_solicitud?query=CodigoAbreviacion:"+codigoAbreviacion, &estadoSoporteSolicitudRes); err != nil {
		return nil, err
	}
	if err := helpers.ExtractDataApi(estadoSoporteSolicitudRes, &estadoSoporteSolicitud); err != nil {
		return nil, err

	}
	if len(estadoSoporteSolicitud) == 0 {
		return nil, errors.New("Estado de Soporte de Solicitud no Encontrado: " + codigo)
	}

	return &estadoSoporteSolicitud[0], nil
}

func ConsultarFormulario(id int) (*models.FormularioSolicitud, error) {
	var formularioRes interface{}
	var formulario models.FormularioSolicitud

	if err := request.GetJson(beego.AppConfig.String("sabaticosService")+"/formulario_solicitud/"+fmt.Sprint(id), &formularioRes); err != nil {
		return nil, err
	}

	if err := helpers.ExtractDataApi(formularioRes, &formulario); err != nil {
		return nil, err

	}

	if formulario.Id == 0 {
		return nil, errors.New("Formulario de Solicitud no Encontrado: " + fmt.Sprint(id))
	}

	return &formulario, nil
}

func ConsultarSoporteSolicitud(id int) (*models.SoporteSolicitud, error) {
	var soporteSolicitudRes interface{}
	var soporteSolicitud models.SoporteSolicitud

	if err := request.GetJson(beego.AppConfig.String("sabaticosService")+"/soporte_solicitud/"+fmt.Sprint(id), &soporteSolicitudRes); err != nil {
		return nil, err
	}

	if err := helpers.ExtractDataApi(soporteSolicitudRes, &soporteSolicitud); err != nil {
		return nil, err

	}

	if soporteSolicitud.Id == 0 {
		return nil, errors.New("Soporte de Solicitud no Encontrado: " + fmt.Sprint(id))
	}

	return &soporteSolicitud, nil
}

// Peticiones POST
func RegistrarSolicitud(terceroId int, tipoSolicitudId int, sabaticoId *int) (*models.Solicitud, error) {
	var solicitudRes interface{}
	var solicitudCreada *models.Solicitud

	solicitud := models.SolicitudCreateRequest{
		TerceroId: terceroId,
		Activo:    true,
		TipoSolicitudId: models.IdReference{
			Id: tipoSolicitudId,
		},
	}

	if err := request.SendJson(beego.AppConfig.String("sabaticosService")+"/solicitud/", "POST", &solicitudRes, solicitud); err != nil {
		beego.Error("Error Solicitud Tercero:", err)
		return nil, err
	}

	if err := helpers.ExtractDataApi(solicitudRes, &solicitudCreada); err != nil {
		beego.Error("Error Extrayendo Datos de Solicitud Tercero:", err)
		return nil, err
	}

	return solicitudCreada, nil
}

func RegistrarHistorialSolicitud(solicitudId int, terceroId int, justificacion string, codigoEstadoSolicitud string) (*models.HistorialSolicitud, error) {
	var historicoResp interface{}

	estadoSolicitud, err := ConsultarEstadoSolicitud(codigoEstadoSolicitud)
	if err != nil {
		return nil, err
	}

	historial := models.HistorialSolicitudCreateRequest{
		TerceroId:     terceroId,
		Justificacion: justificacion,
		Activo:        true,
		EstadoSolicitudId: models.IdReference{
			Id: estadoSolicitud.Id,
		},
		SolicitudId: models.IdReference{
			Id: solicitudId,
		},
	}

	if err := request.SendJson(beego.AppConfig.String("sabaticosService")+"/historial_solicitud/", "POST", &historicoResp, historial); err != nil {
		beego.Error("Error Histórico:", err)
	}

	var historialFinal *models.HistorialSolicitud

	if err := helpers.ExtractDataApi(historicoResp, &historialFinal); err != nil {
		beego.Error("Error extrayendo Datos de Histórico:", err)
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
		beego.Error("Error formulario:", err)
	}

	var formularioFinal *models.FormularioSolicitud

	if err := helpers.ExtractDataApi(formularioResp, &formularioFinal); err != nil {
		beego.Error("Error Extrayendo Datos de Formulario:", err)
		return nil, err
	}

	return formularioFinal, nil

}

func RegistrarSoporteSolicitud(documentoId int, terceroId int, solicitudId int, estadoSoporteSolicitudId int, rolUsuario string) (*models.SoporteSolicitud, error) {
	var soporteSolicitudRes interface{}
	var soporteSolicitudFinal *models.SoporteSolicitud

	soporteSolicitud := models.SoporteSolicitudCreateRequest{
		DocumentoId:              documentoId,
		TerceroId:                terceroId,
		Activo:                   true,
		SolicitudId:              models.IdReference{Id: solicitudId},
		EstadoSoporteSolicitudId: models.IdReference{Id: estadoSoporteSolicitudId},
		RolUsuario:               rolUsuario,
	}

	if err := request.SendJson(beego.AppConfig.String("sabaticosService")+"/soporte_solicitud", "POST", &soporteSolicitudRes, soporteSolicitud); err != nil {
		return nil, fmt.Errorf("falló creación en sabaticosService /soporte_solicitud: %w", err)
	}

	if err := helpers.ValidateServiceResponse(soporteSolicitudRes); err != nil {
		return nil, fmt.Errorf("sabaticosService /soporte_solicitud devolvió error: %w", err)
	}

	if err := helpers.ExtractDataApi(soporteSolicitudRes, &soporteSolicitudFinal); err != nil {
		beego.Error("Error extrayendo Datos de Soporte Solicitud:", err)
		return nil, err
	}

	return soporteSolicitudFinal, nil
}

// peticiones PUT
func ActualizarFormularioSolicitud(solicitudId int, formularioId int, contenido string) (*models.FormularioSolicitud, error) {
	var formularioResp interface{}
	var formularioFinal *models.FormularioSolicitud

	formularioExistente, err := ConsultarFormulario(formularioId)
	if err != nil {
		return nil, err
	}

	formulario := models.FormularioSolicitudCreateRequest{
		Contenido:     contenido,
		Activo:        true,
		FechaCreacion: formularioExistente.FechaCreacion,
		SolicitudId:   models.IdReference{Id: solicitudId},
	}

	if err := request.SendJson(beego.AppConfig.String("sabaticosService")+"/formulario_solicitud/"+fmt.Sprint(formularioId), "PUT", &formularioResp, formulario); err != nil {
		beego.Error("Error Actualizando Formulario:", err)
		return nil, err
	}

	if err := helpers.ValidateServiceResponse(formularioResp); err != nil {
		return nil, fmt.Errorf("sabaticosService /formulario_solicitud/%d devolvió error: %w", formularioId, err)
	}

	if err := helpers.ExtractDataApi(formularioResp, &formularioFinal); err != nil {
		beego.Error("Error Extrayendo Datos de Formulario:", err)
		return nil, err
	}

	return formularioFinal, nil
}

func ActualizarSoporteSolicitud(soporteId int, solicitudId int, ObtenerCodigoEstadoSoporteSolicitud string) (*models.SoporteSolicitud, error) {
	var soporteSolicitudRes interface{}
	var soporteSolicitudFinal *models.SoporteSolicitud

	estadoSoporteSolicitud, err := ConsultarEstadoSoporteSolicitud(ObtenerCodigoEstadoSoporteSolicitud)
	if err != nil {
		return nil, err
	}

	// Primero obtener el soporte existente
	soporteExistente, err := ConsultarSoporteSolicitud(soporteId)
	if err != nil {
		return nil, fmt.Errorf("no se pudo consultar soporte_solicitud/%d: %w", soporteId, err)
	}

	if soporteExistente.DocumentoId <= 0 {
		return nil, fmt.Errorf("soporte_solicitud/%d no tiene DocumentoId válido", soporteId)
	}

	// Actualizar con los nuevos valores
	soporteSolicitud := models.SoporteSolicitudCreateRequest{
		DocumentoId:              soporteExistente.DocumentoId,
		TerceroId:                soporteExistente.TerceroId,
		Activo:                   true,
		FechaCreacion:            soporteExistente.FechaCreacion,
		SolicitudId:              models.IdReference{Id: solicitudId},
		EstadoSoporteSolicitudId: models.IdReference{Id: estadoSoporteSolicitud.Id},
		RolUsuario:               soporteExistente.RolUsuario,
	}

	if err := request.SendJson(beego.AppConfig.String("sabaticosService")+"/soporte_solicitud/"+fmt.Sprint(soporteId), "PUT", &soporteSolicitudRes, soporteSolicitud); err != nil {
		return nil, fmt.Errorf("falló actualización en sabaticosService /soporte_solicitud/%d: %w", soporteId, err)
	}

	if err := helpers.ValidateServiceResponse(soporteSolicitudRes); err != nil {
		return nil, fmt.Errorf("sabaticosService /soporte_solicitud/%d devolvió error: %w", soporteId, err)
	}

	if err := helpers.ExtractDataApi(soporteSolicitudRes, &soporteSolicitudFinal); err != nil {
		beego.Error("Error extrayendo Datos de Soporte Solicitud:", err)
		return nil, err
	}

	return soporteSolicitudFinal, nil
}
