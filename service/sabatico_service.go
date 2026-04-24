package service

import (
	"fmt"

	"github.com/udistrital/sabaticos_mid/clients"
	"github.com/udistrital/sabaticos_mid/models"
)

func CrearSabatico(
	solicitudId int,
	terceroId int,
	observaciones string,
	fechaInicio string,
	fechaFin string,
	estado string,
) (*models.CrearSabaticoResult, error) {

	estadoRequest := models.SolicitudAprobarRechazarRequest{
		TerceroId:       terceroId,
		SolicitudId:     solicitudId,
		Justificacion:   "Cambio automático al crear año sabático",
		EstadoSolicitud: "S12",
		EstadoSoporte:   "SGOK",
	}

	_, err := CambiarEstado(estadoRequest)
	if err != nil {
		return nil, fmt.Errorf(
			"error cambiando estado: %v",
			err,
		)
	}

	sabaticoCreado, err := clients.RegistrarSabatico(
		solicitudId,
		terceroId,
		observaciones,
		fechaInicio,
		fechaFin,
		estado,
	)

	if err != nil {
		return nil, err
	}

	err = clients.AsociarSabaticoSolicitud(
		solicitudId,
		sabaticoCreado.Id,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"error asociando sabático a solicitud: %v",
			err,
		)
	}

	return sabaticoCreado, nil
}
