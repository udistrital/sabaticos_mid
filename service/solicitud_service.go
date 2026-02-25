package service

import (
	"api_mid_sabaticos/clients"
	"api_mid_sabaticos/enums"
	"api_mid_sabaticos/models"

	"github.com/astaxie/beego"
)

func CrearSolicitud(solicitudReq models.SolicitudRequest) (*models.Solicitud, error) {
	terceroId := solicitudReq.TerceroId
	codigoTipoSolicitud := solicitudReq.TipoSolicitudId
	sabatico := solicitudReq.SabaticoId
	formulario := solicitudReq.Formulario

	// Validar tercero
	if err := clients.ValidarTercero(terceroId); err != nil {
		return nil, err
	}

	// Crear solicitud en CRUD y obtener ID
	solicitud, err := clients.RegistrarSolicitud(terceroId, codigoTipoSolicitud, sabatico)
	if err != nil {
		return nil, err
	}

	// Crear historial y formulario en paralelo
	_, _, err = registrarHistorialYFormulario(solicitud.Id, terceroId, sabatico, codigoTipoSolicitud, string(formulario))
	if err != nil {
		beego.Error("Error en proceso posterior de solicitud:", err)
		return nil, err
	}

	// Solo retornar la solicitud si TODO fue exitoso
	return solicitud, nil
}

func registrarHistorialYFormulario(solicitudId int, terceroId int, sabaticoId *int, tipoSolicitud string, formularioRequest string) (*models.HistorialSolicitud, *models.FormularioSolicitud, error) {
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

	// Validar si es solicitud NUEVA usando el enum
	codigoTipoSolicitud, esValido := enums.ObtenerCodigoTipoSolicitud(tipoSolicitud)

	// Solo crear formulario si es solicitud NUEVA (NS) y no existe sabáticoId
	if esValido && codigoTipoSolicitud == string(enums.NUEVA) && sabaticoId == nil {
		go func() {
			formulario, formularioErr = clients.RegistrarFormularioSolicitud(solicitudId, formularioRequest)
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
