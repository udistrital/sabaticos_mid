package service

import (
	"api_mid_sabaticos/clients"
	"api_mid_sabaticos/models"
	"errors"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func CrearSolicitud(solicitudReq models.SolicitudRequest) (*models.Solicitud, *models.HistorialSolicitud, *models.FormularioSolicitud, error) {
	// Validar tercero
	if err := clients.ValidarTercero(solicitudReq.TerceroId); err != nil {
		return nil, nil, nil, err
	}

	// Crear solicitud en CRUD y obtener ID
	solicitud, err := clients.RegistrarSolicitud(solicitudReq)
	if err != nil {
		return nil, nil, nil, err
	}

	// Crear historial y formulario en paralelo
	historial, formulario, err := crearHistorialYFormulario(solicitudReq, solicitud)
	if err != nil {
		return nil, nil, nil, err
	}

	return solicitud, historial, formulario, nil
}

func crearHistorialYFormulario(solicitudReq models.SolicitudRequest, solicitudCreada *models.Solicitud) (*models.HistorialSolicitud, *models.FormularioSolicitud, error) {
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
		historial, err := clients.RegistrarHistorialSolicitud(models.HistorialSolicitud{}, solicitudCreada, solicitudReq)
		if err != nil {
			beego.Error("Error POST histórico:", err)
		}
		historialChan <- historialResult{historial: historial, err: err}
	}()

	// Goroutine para crear formulario
	go func() {
		formulario, err := clients.RegistrarFormularioSolicitud(models.FormularioSolicitud{}, solicitudCreada)
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

func ValidarSolicitud(nuevo models.HistorialSolicitud) (*models.HistorialSolicitud, bool, error) {
	if nuevo.SolicitudId <= 0 {
		return nil, false, errors.New("solicitud_id requerido")
	}
	if nuevo.EstadoSolicitudId <= 0 {
		return nil, false, errors.New("estado_solicitud_id requerido")
	}

	actual, err := models.GetHistorialSolicitudActual(nuevo.SolicitudId)
	if err != nil && err != orm.ErrNoRows {
		return nil, false, err
	}

	if err != orm.ErrNoRows && actual.EstadoSolicitudId == nuevo.EstadoSolicitudId {
		return actual, false, nil
	}

	now := time.Now()
	nuevo.Id = 0
	nuevo.Activo = true
	nuevo.FechaModificacion = &now

	o := orm.NewOrm()
	if _, err := o.Insert(&nuevo); err != nil {
		return nil, false, err
	}

	return &nuevo, true, nil
}
