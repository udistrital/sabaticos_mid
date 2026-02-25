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
	historial, formulario, err := registrarHistorialYFormulario(solicitud.Id, terceroId, sabatico)
	if err != nil {
		return nil, nil, nil, err
	}

	return solicitud, historial, formulario, nil
}

func registrarHistorialYFormulario(solicitudId int, terceroId int, sabaticoId *int) (*models.HistorialSolicitud, *models.FormularioSolicitud, error) {
	var historial *models.HistorialSolicitud
	var formulario *models.FormularioSolicitud
	var historialErr, formularioErr error

	// Canal para sincronizar goroutines
	done := make(chan bool, 2)

	// Crear historial en goroutine
	go func() {
		historial, historialErr = clients.RegistrarHistorialSolicitud(solicitudId, terceroId)
		if historialErr != nil {
			beego.Error("Error registrando historial de solicitud:", historialErr)
		}
		done <- true
	}()

	// Solo crear formulario si no existe sabáticoId
	if sabaticoId == nil {
		go func() {
			formulario, formularioErr = clients.RegistrarFormularioSolicitud(solicitudId)
			if formularioErr != nil {
				beego.Error("Error registrando formulario de solicitud:", formularioErr)
			}
			done <- true
		}()
	} else {
		// Si no se crea formulario, marcar como completado
		done <- true
	}

	// Esperar a que ambas goroutines terminen
	<-done
	<-done

	// Verificar errores
	if historialErr != nil {
		return nil, nil, historialErr
	}
	if formularioErr != nil {
		return nil, nil, formularioErr
	}

	return historial, formulario, nil
}
