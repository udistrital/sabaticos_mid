package service

import (
	"api_mid_sabaticos/clients"
	"api_mid_sabaticos/models"

	"github.com/astaxie/beego"
)

func CrearSolicitud(solicitudReq models.SolicitudRequest) (*models.Solicitud, *models.HistorialSolicitud, *models.FormularioSolicitud, error) {
	terceroId := solicitudReq.TerceroId
	codigoTipoSolicitud := solicitudReq.TipoSolicitudId
	sabatico := solicitudReq.SabaticoId

	// Validar tercero
	if err := clients.ValidarTercero(terceroId); err != nil {
		return nil, nil, nil, err
	}

	// Crear solicitud en CRUD y obtener ID
	solicitud, err := clients.RegistrarSolicitud(terceroId, codigoTipoSolicitud, sabatico)
	if err != nil {
		return nil, nil, nil, err
	}

	// Crear historial y formulario en paralelo
	historial, formulario, err := registrarHistorialYFormulario(solicitud.Id, terceroId)
	if err != nil {
		return nil, nil, nil, err
	}

	return solicitud, historial, formulario, nil
}

func registrarHistorialYFormulario(solicitudId int, terceroId int) (*models.HistorialSolicitud, *models.FormularioSolicitud, error) {
	type historialResult struct {
		historial *models.HistorialSolicitud
		err       error
	}
	type formularioResult struct {
		formulario *models.FormularioSolicitud
		err        error
	}

	historialChan := make(chan historialResult, 1)
	formularioChan := make(chan formularioResult, 1)

	// Goroutine para crear histórico
	go func() {
		historial, err := clients.RegistrarHistorialSolicitud(solicitudId, terceroId)
		if err != nil {
			beego.Error("Error POST histórico:", err)
		}
		historialChan <- historialResult{historial: historial, err: err}
	}()

	// Goroutine para crear formulario
	go func() {
		formulario, err := clients.RegistrarFormularioSolicitud(solicitudId)
		if err != nil {
			beego.Error("Error POST formulario:", err)
		}
		formularioChan <- formularioResult{formulario: formulario, err: err}
	}()

	historialRes := <-historialChan
	if historialRes.err != nil {
		beego.Error("Error en proceso paralelo:", historialRes.err)
		return nil, nil, historialRes.err
	}

	formularioRes := <-formularioChan
	if formularioRes.err != nil {
		beego.Error("Error en proceso paralelo:", formularioRes.err)
		return nil, nil, formularioRes.err
	}

	return historialRes.historial, formularioRes.formulario, nil
}
