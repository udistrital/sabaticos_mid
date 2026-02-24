package service

import (
	"api_mid_sabaticos/enums"
	"api_mid_sabaticos/models"
	"fmt"

	"api_mid_sabaticos/helpers"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

func CrearSolicitud(solicitudReq models.SolicitudRequest) (*models.Solicitud, *models.HistorialSolicitud, *models.FormularioSolicitud, error) {
	// Validar tercero
	if err := validarTercero(solicitudReq.TerceroId); err != nil {
		return nil, nil, nil, err
	}

	// Crear solicitud en CRUD y obtener ID
	solicitudCreada, err := crearSolicitud(solicitudReq)
	if err != nil {
		return nil, nil, nil, err
	}

	// Crear historial y formulario en paralelo
	historicoResp, formularioResp, err := crearHistorialYFormulario(solicitudReq, solicitudCreada)
	if err != nil {
		return nil, nil, nil, err
	}

	// Extraer y convertir respuestas
	solicitud, historial, formulario := extraerRespuestas(solicitudCreada, historicoResp, formularioResp)

	return solicitud, historial, formulario, nil
}

func validarTercero(terceroId int) error {
	var tercero interface{}
	if err := request.GetJson(beego.AppConfig.String("terceroService")+"tercero/"+fmt.Sprintf("%d", terceroId), &tercero); err != nil {
		beego.Error("Error GET tercero:", err)
		return err
	}
	return nil
}

func crearSolicitud(solicitudReq models.SolicitudRequest) (*models.Solicitud, error) {
	tipoSolicitudId := helpers.InterfaceToInt(solicitudReq.TipoSolicitudId, int(enums.NUEVA))
	sabaticoId := helpers.InterfaceToIntPtr(solicitudReq.SabaticoId)

	solicitud := models.Solicitud{
		TerceroId:       solicitudReq.TerceroId,
		TipoSolicitudId: tipoSolicitudId,
		SabaticoId:      sabaticoId,
		Activo:          true,
	}

	var solicitudRes interface{}
	if err := request.SendJson(beego.AppConfig.String("sabaticosService")+"/solicitud/", "POST", &solicitudRes, solicitud); err != nil {
		beego.Error("Error POST solicitud:", err)
		return nil, err
	}

	var solicitudCreada models.Solicitud
	helpers.ExtractDataApi(solicitudRes, &solicitudCreada)
	beego.Info("ID de solicitud creada:", solicitudCreada.Id)

	return &solicitudCreada, nil
}

func crearHistorialYFormulario(solicitudReq models.SolicitudRequest, solicitudCreada *models.Solicitud) (interface{}, interface{}, error) {
	historial := models.HistorialSolicitud{
		TerceroId:         solicitudReq.TerceroId,
		Justificacion:     "Nueva Solicitud Creada",
		Activo:            true,
		EstadoSolicitudId: int(enums.ENVIADA),
		SolicitudId:       solicitudCreada.Id,
	}

	formulario := models.FormularioSolicitud{
		Contenido:   "{}",
		SolicitudId: solicitudCreada.Id,
		Activo:      true,
	}

	errChan := make(chan error, 2)
	var historicoResp interface{}
	var formularioResp interface{}

	// Goroutine para crear histórico
	go func() {
		if err := request.SendJson(beego.AppConfig.String("sabaticosService")+"/historial_solicitud/", "POST", &historicoResp, historial); err != nil {
			beego.Error("Error POST histórico:", err)
			errChan <- err
			return
		}
		errChan <- nil
	}()

	// Goroutine para crear formulario
	go func() {
		if err := request.SendJson(beego.AppConfig.String("sabaticosService")+"/formulario_solicitud/", "POST", &formularioResp, formulario); err != nil {
			beego.Error("Error POST formulario:", err)
			errChan <- err
			return
		}
		errChan <- nil
	}()

	// Esperar resultados
	for i := 0; i < 2; i++ {
		if err := <-errChan; err != nil {
			beego.Error("Error en proceso paralelo:", err)
			return nil, nil, err
		}
	}

	return historicoResp, formularioResp, nil
}

func extraerRespuestas(solicitudCreada *models.Solicitud, historicoResp interface{}, formularioResp interface{}) (*models.Solicitud, *models.HistorialSolicitud, *models.FormularioSolicitud) {
	var historial *models.HistorialSolicitud
	var formulario *models.FormularioSolicitud

	helpers.ExtractDataApi(historicoResp, &historial)
	helpers.ExtractDataApi(formularioResp, &formulario)

	return solicitudCreada, historial, formulario
}
