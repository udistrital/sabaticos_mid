package service

import (
	"api_mid_sabaticos/clients"
	"api_mid_sabaticos/enums"
	"api_mid_sabaticos/models"
	"errors"
	"time"

	"github.com/astaxie/beego"
)

func CrearSolicitud(solicitudReq models.SolicitudRequest) (*models.Solicitud, error) {
	terceroId := solicitudReq.TerceroId
	codigoTipoSolicitud := solicitudReq.TipoSolicitudId
	sabaticoId := solicitudReq.SabaticoId
	formulario := solicitudReq.Formulario

	// Validar tercero
	//Se comenta por que por el momento no hay usuario de prueba en el servicio de terceros
	// if err := clients.ValidarTercero(terceroId); err != nil {
	// 	return nil, err
	// }

	tipoSolicitud, err := clients.ConsultarTipoSolicitud(codigoTipoSolicitud)
	if err != nil {
		return nil, err
	}

	if err := validarSolicitudPorTipo(tipoSolicitud, sabaticoId); err != nil {
		return nil, err
	}

	// Crear solicitud en CRUD y obtener ID
	solicitud, err := clients.RegistrarSolicitud(terceroId, tipoSolicitud.Id, sabaticoId)
	if err != nil {
		return nil, err
	}

	// Crear historial y formulario en paralelo
	_, _, err = registrarHistorialYFormulario(solicitud.Id, terceroId, sabaticoId, codigoTipoSolicitud, string(formulario))
	if err != nil {
		return nil, err
	}

	// Solo retornar la solicitud si TODO fue exitoso
	return solicitud, nil
}

func validarSolicitudPorTipo(tipoSolicitud *models.TipoSolicitud, sabaticoId *int) error {
	if tipoSolicitud.CodigoAbreviacion == string(enums.NUEVA) {
		if sabaticoId != nil {
			return errors.New("No se Puede Crear Una Solicitud NUEVA Con Un Sabático Asociado")
		}
		return nil
	}

	if tipoSolicitud.CodigoAbreviacion != string(enums.SUSPENSION) {
		return nil
	}

	if sabaticoId == nil {
		return errors.New("Una Solicitud de SUSPENSIÓN debe Tener un Sabático Asociado")
	}

	// Consultar si el sabático existe
	sabatico, err := clients.ConsultarSabatico(*sabaticoId)
	if err != nil {
		return err
	}

	// Validar que el sabático tenga máximo 3 meses desde su creación
	fechaCreacion, err := time.Parse("2006-01-02 15:04:05 -0700 -0700", sabatico.FechaCreacion)
	if err != nil {
		return errors.New("Formato Inválido de FechaCreacion del Sabático")
	}
	fechaLimite := fechaCreacion.AddDate(0, 3, 0)
	if time.Now().After(fechaLimite) {
		return errors.New("No se Puede Crear Una Solicitud de SUSPENSIÓN Sespués de 3 Meses Desde la Creación del Sabático")
	}

	return nil
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
			beego.Error("Error Registrando Historial de Solicitud:", historialErr)
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
				beego.Error("Error Registrando Formulario de Solicitud:", formularioErr)
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
