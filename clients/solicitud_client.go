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
		return nil, errors.New("request not found: " + fmt.Sprint(id))
	}

	return &solicitud, nil
}

/*
ConsultarEstadoActualSolicitud tiene como objetivo consultar el estado actual de una solicitud;
consume el crud para traer el historial más reciente activo de esa solicitud.
*/
func ConsultarIdsHistorialSolicitud(idSolicitud int) ([]int, error) {
	var historialSolicitudRes interface{}
	var historialSolicitud []models.HistorialSolicitud

	url := beego.AppConfig.String("sabaticosService") +
		"/historial_solicitud?query=Activo:true,SolicitudId:" +
		fmt.Sprint(idSolicitud) +
		"&sortby=FechaCreacion&order=desc&limit=0"

	if err := request.GetJson(url, &historialSolicitudRes); err != nil {
		return nil, err
	}

	if err := helpers.ExtractDataApi(historialSolicitudRes, &historialSolicitud); err != nil {
		return nil, err
	}

	ids := make([]int, 0, len(historialSolicitud))
	for _, hs := range historialSolicitud {
		ids = append(ids, hs.Id)
	}

	return ids, nil
}

func ConsultarEstadoSolicitud(codigo string) (*models.EstadoSolicitud, error) {
	var estadoSolicitudRes interface{}
	var estadoSolicitud []models.EstadoSolicitud

	codigoAbreviacion, ok := enums.ObtenerCodigoEstadoSolicitud(codigo)
	if !ok {
		return nil, errors.New("request status not valid: " + codigo)
	}

	baseURL := beego.AppConfig.String("sabaticosService")
	if baseURL == "" {
		return nil, errors.New("config sabaticosService is empty")
	}

	url := baseURL + "/estado_solicitud?query=CodigoAbreviacion:" + codigoAbreviacion

	if err := request.GetJson(url, &estadoSolicitudRes); err != nil {
		return nil, err
	}

	if err := helpers.ExtractDataApi(estadoSolicitudRes, &estadoSolicitud); err != nil {
		return nil, err
	}

	if len(estadoSolicitud) == 0 {
		return nil, errors.New("request status not found: " + codigo)
	}

	return &estadoSolicitud[0], nil
}

func ConsultarHistorialSolicitud(idHistorialSolicitud int) (*models.HistorialSolicitud, error) {
	var historialRes interface{}
	var historial models.HistorialSolicitud

	url := beego.AppConfig.String("sabaticosService") +
		"/historial_solicitud/" + fmt.Sprint(idHistorialSolicitud)

	if err := request.GetJson(url, &historialRes); err != nil {
		return nil, err
	}

	if err := helpers.ExtractDataApi(historialRes, &historial); err != nil {
		return nil, err
	}

	if historial.Id == 0 {
		return nil, errors.New("historial_solicitud not found: " + fmt.Sprint(idHistorialSolicitud))
	}

	return &historial, nil
}

func ConsultarTipoSolicitud(codigo string) (*models.TipoSolicitud, error) {
	var tipoSolicitudRes interface{}
	var tipoSolicitud []models.TipoSolicitud

	codigoAbreviacion, ok := enums.ObtenerCodigoTipoSolicitud(codigo)
	if !ok {
		return nil, errors.New("request type not valid: " + codigo)
	}

	if err := request.GetJson(beego.AppConfig.String("sabaticosService")+"/tipo_solicitud?query=CodigoAbreviacion:"+codigoAbreviacion, &tipoSolicitudRes); err != nil {
		return nil, err
	}
	if err := helpers.ExtractDataApi(tipoSolicitudRes, &tipoSolicitud); err != nil {
		return nil, err

	}
	if len(tipoSolicitud) == 0 {
		return nil, errors.New("request type not found: " + codigo)
	}

	return &tipoSolicitud[0], nil
}

func ConsultarEstadoSoporteSolicitud(codigo string) (*models.EstadoSoporteSolicitud, error) {
	var estadoSoporteSolicitudRes interface{}
	var estadoSoporteSolicitud []models.EstadoSoporteSolicitud

	codigoAbreviacion, ok := enums.ObtenerCodigoEstadoSoporteSolicitud(codigo)

	if !ok {
		return nil, errors.New("request support status not valid: " + codigo)
	}

	if err := request.GetJson(beego.AppConfig.String("sabaticosService")+"/estado_soporte_solicitud?query=CodigoAbreviacion:"+codigoAbreviacion, &estadoSoporteSolicitudRes); err != nil {
		return nil, err
	}
	if err := helpers.ExtractDataApi(estadoSoporteSolicitudRes, &estadoSoporteSolicitud); err != nil {
		return nil, err

	}
	if len(estadoSoporteSolicitud) == 0 {
		return nil, errors.New("request support status not found: " + codigo)
	}

	return &estadoSoporteSolicitud[0], nil
}

func ConsultarHistorialSolicitudIdEstadoId(historialId int, estadosSolicitud []string) ([]models.HistorialSolicitud, error) {
	var historial []models.HistorialSolicitud

	decodeHistorial := func(res interface{}) ([]models.HistorialSolicitud, error) {
		var items []models.HistorialSolicitud
		if err := helpers.ExtractDataApi(res, &items); err == nil {
			return items, nil
		}

		var item models.HistorialSolicitud
		if err := helpers.ExtractDataApi(res, &item); err != nil {
			return nil, err
		}

		if item.Id == 0 {
			return []models.HistorialSolicitud{}, nil
		}

		return []models.HistorialSolicitud{item}, nil
	}

	baseURL := beego.AppConfig.String("sabaticosService") +
		"/historial_solicitud?query=Activo:true,Id:" + fmt.Sprint(historialId)

	if len(estadosSolicitud) == 0 {
		var historialRes interface{}
		if err := request.GetJson(baseURL, &historialRes); err != nil {
			return nil, err
		}
		return decodeHistorial(historialRes)
	}

	vistos := make(map[int]bool)

	for _, estado := range estadosSolicitud {
		codigoAbreviacion, ok := enums.ObtenerCodigoEstadoSolicitud(estado)
		if !ok || codigoAbreviacion == "" {
			return nil, fmt.Errorf("invalid request status code: %s", estado)
		}

		estadoSolicitud, errConsult := ConsultarEstadoSolicitud(codigoAbreviacion)
		if errConsult != nil {
			return nil, fmt.Errorf("error consulting request status for code %s: %w", estado, errConsult)
		}

		url := baseURL + ",EstadoSolicitudId.Id:" + fmt.Sprint(estadoSolicitud.Id)
		var historialRes interface{}
		if err := request.GetJson(url, &historialRes); err != nil {
			return nil, err
		}

		items, err := decodeHistorial(historialRes)
		if err != nil {
			return nil, err
		}

		for _, item := range items {
			if !vistos[item.Id] {
				historial = append(historial, item)
				vistos[item.Id] = true
			}
		}
	}

	return historial, nil
}

func ConsultarTodosFormulariosSolicitud() ([]models.FormularioSolicitud, error) {
	var formulariosRes interface{}
	var formularios []models.FormularioSolicitud

	url := beego.AppConfig.String("sabaticosService") + "/formulario_solicitud?query=Activo:true&sortby=FechaCreacion&order=desc&limit=0"

	if err := request.GetJson(url, &formulariosRes); err != nil {
		return nil, err
	}

	if err := helpers.ExtractDataApi(formulariosRes, &formularios); err != nil {
		return nil, err
	}

	return formularios, nil
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
		return nil, errors.New("request form not found: " + fmt.Sprint(id))
	}

	return &formulario, nil
}

func ConsultarSoporteSolicitud(documetoId int) (*models.SoporteSolicitud, error) {
	var soporteSolicitudRes interface{}
	var soporteSolicitud []models.SoporteSolicitud

	if err := request.GetJson(beego.AppConfig.String("sabaticosService")+"soporte_solicitud?query=DocumentoId:"+fmt.Sprint(documetoId), &soporteSolicitudRes); err != nil {
		return nil, err
	}

	if err := helpers.ExtractDataApi(soporteSolicitudRes, &soporteSolicitud); err != nil {
		return nil, err

	}

	if soporteSolicitud[0].Id == 0 {
		return nil, errors.New("request support not found: " + fmt.Sprint(documetoId))
	}

	return &soporteSolicitud[0], nil
}

/*
	 Consultar Soporte Solicitud Tiene el Objetivo dar un arreglo
		con todos los soporte asociados a una solicitud
*/
func ConsultarSoportesSolicitud(id int) ([]models.SoporteSolicitud, error) {

	var res interface{}
	var soportes []models.SoporteSolicitud

	// Construcción de la URL usando variable de entorno
	url := fmt.Sprintf(
		"%s/soporte_solicitud?query=Activo:true,SolicitudId:%d&limit=0",
		beego.AppConfig.String("sabaticosService"),
		id,
	)

	// Consumo del servicio
	if err := request.GetJson(url, &res); err != nil {
		return nil, err
	}

	// Extracción del data (formato estándar OAS)
	if err := helpers.ExtractDataApi(res, &soportes); err != nil {
		return nil, err
	}

	// Validación: si no hay resultados, retornar slice vacío (no error)
	if len(soportes) == 0 {
		return []models.SoporteSolicitud{}, nil
	}

	return soportes, nil
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
		beego.Error("error registering request for third party:", err)
		return nil, err
	}

	if err := helpers.ExtractDataApi(solicitudRes, &solicitudCreada); err != nil {
		beego.Error("error extracting request data for third party:", err)
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
		beego.Error("error registering history:", err)
	}

	var historialFinal *models.HistorialSolicitud

	if err := helpers.ExtractDataApi(historicoResp, &historialFinal); err != nil {
		beego.Error("error extracting history data:", err)
		return nil, err
	}

	return historialFinal, nil
}

func RegistrarHistorialSolicitudEstado(solicitudId int, terceroId int, justificacion string, estadoSolicitudIdEntrante int) (*models.HistorialSolicitud, error) {

	var historicoResp interface{}

	historial := models.HistorialSolicitudCreateRequest{
		TerceroId:     terceroId,
		Justificacion: justificacion,
		Activo:        true,
		EstadoSolicitudId: models.IdReference{
			Id: estadoSolicitudIdEntrante,
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
		beego.Error("error registering form:", err)
	}

	var formularioFinal *models.FormularioSolicitud

	if err := helpers.ExtractDataApi(formularioResp, &formularioFinal); err != nil {
		beego.Error("error extracting form data:", err)
		return nil, err
	}

	return formularioFinal, nil

}

func RegistrarSoporteSolicitud(documentoId int, terceroId int, solicitudId int, estadoSoporteSolicitudId int, rolUsuario string, tipoDocumentoId int) (*models.SoporteSolicitud, error) {
	var soporteSolicitudRes interface{}
	var soporteSolicitudFinal *models.SoporteSolicitud

	soporteSolicitud := models.SoporteSolicitudCreateRequest{
		DocumentoId:              documentoId,
		TerceroId:                terceroId,
		TipoDocumentoId:          tipoDocumentoId,
		Activo:                   true,
		SolicitudId:              models.IdReference{Id: solicitudId},
		EstadoSoporteSolicitudId: models.IdReference{Id: estadoSoporteSolicitudId},
		RolUsuario:               rolUsuario,
	}

	if err := request.SendJson(beego.AppConfig.String("sabaticosService")+"/soporte_solicitud", "POST", &soporteSolicitudRes, soporteSolicitud); err != nil {
		return nil, fmt.Errorf("creation failed in sabaticosService /soporte_solicitud: %w", err)
	}

	if err := helpers.ValidateServiceResponse(soporteSolicitudRes); err != nil {
		return nil, fmt.Errorf("sabaticosService /soporte_solicitud returned error: %w", err)
	}

	if err := helpers.ExtractDataApi(soporteSolicitudRes, &soporteSolicitudFinal); err != nil {
		beego.Error("error extracting support request data:", err)
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
		beego.Error("error updating form:", err)
		return nil, err
	}

	if err := helpers.ValidateServiceResponse(formularioResp); err != nil {
		return nil, fmt.Errorf("sabaticosService /formulario_solicitud/%d returned error: %w", formularioId, err)
	}

	if err := helpers.ExtractDataApi(formularioResp, &formularioFinal); err != nil {
		beego.Error("error extracting form data:", err)
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
		return nil, fmt.Errorf("could not query soporte_solicitud/%d: %w", soporteId, err)
	}

	if soporteExistente.DocumentoId <= 0 {
		return nil, fmt.Errorf("soporte_solicitud/%d has no valid DocumentoId", soporteId)
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

	if err := request.SendJson(beego.AppConfig.String("sabaticosService")+"soporte_solicitud/"+fmt.Sprint(soporteExistente.Id), "PUT", &soporteSolicitudRes, soporteSolicitud); err != nil {
		return nil, fmt.Errorf("update failed in sabaticosService /soporte_solicitud/%d: %w", soporteId, err)
	}

	if err := helpers.ValidateServiceResponse(soporteSolicitudRes); err != nil {
		return nil, fmt.Errorf("sabaticosService /soporte_solicitud/%d returned error: %w", soporteId, err)
	}

	if err := helpers.ExtractDataApi(soporteSolicitudRes, &soporteSolicitudFinal); err != nil {
		beego.Error("error extracting support request data:", err)
		return nil, err
	}

	return soporteSolicitudFinal, nil
}

/*
Esta función tiene el Objetivo de dado un id de historial solicitud,
Actualizar el registro de historial solicitud en su valor de activo a false
*/
func DesactivarHistorialSolicitud(idHistorialSolicitud int) (bool, error) {
	var historialResp interface{}
	var historialActual *models.HistorialSolicitud
	var historialFinal *models.HistorialSolicitud

	historialActual, err := ConsultarHistorialSolicitud(idHistorialSolicitud)
	if err != nil {
		return false, err
	}

	historial := models.HistorialSolicitud{
		Id:                historialActual.Id,
		TerceroId:         historialActual.TerceroId,
		Justificacion:     historialActual.Justificacion,
		Activo:            false,
		FechaCreacion:     historialActual.FechaCreacion,
		FechaModificacion: historialActual.FechaModificacion,
		EstadoSolicitudId: historialActual.EstadoSolicitudId,
		SolicitudId:       historialActual.SolicitudId,
	}

	url := beego.AppConfig.String("sabaticosService") +
		"/historial_solicitud/" + fmt.Sprint(idHistorialSolicitud)

	if err := request.SendJson(url, "PUT", &historialResp, historial); err != nil {
		beego.Error("error updating historial_solicitud:", err)
		return false, err
	}

	if err := helpers.ValidateServiceResponse(historialResp); err != nil {
		return false, fmt.Errorf("sabaticosCrudService /historial_solicitud/%d returned error: %w", idHistorialSolicitud, err)
	}

	if err := helpers.ExtractDataApi(historialResp, &historialFinal); err != nil {
		beego.Error("error extracting historial_solicitud data:", err)
		return false, err
	}

	if historialFinal == nil || historialFinal.Id == 0 || historialFinal.Activo {
		return false, errors.New("historial_solicitud was not deactivated: " + fmt.Sprint(idHistorialSolicitud))
	}

	return true, nil
}
