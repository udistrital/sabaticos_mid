package service

import (
	"errors"
	"fmt"

	"github.com/udistrital/sabaticos_mid/clients"
	"github.com/udistrital/sabaticos_mid/models"
)

func CrearSabatico(sabaticoReq models.CrearSabaticoRequest) (*models.CrearSabaticoResult, error) {
	fmt.Println("---------- Entra a Service Crear Sabatico -------------")
	terceroId := sabaticoReq.TerceroId
	estadoSabaticoId := sabaticoReq.EstadoSabaticoId
	observaciones := sabaticoReq.Observaciones
	fechaInicio := sabaticoReq.FechaInicio
	fechaFin := sabaticoReq.FechaFin

	if terceroId <= 0 {
		return nil, errors.New("tercero_id es obligatorio")
	}

	if estadoSabaticoId <= 0 {
		return nil, errors.New("estado_sabatico_id es obligatorio")
	}

	if fechaInicio == "" || fechaFin == "" {
		return nil, errors.New("fecha_inicio y fecha_fin son obligatorias")
	}

	sabatico, err := clients.RegistrarSabatico(terceroId, observaciones, fechaInicio, fechaFin, estadoSabaticoId)
	if err != nil {
		return nil, fmt.Errorf("error registrando sabático: %v", err)
	}

	return sabatico, nil
}
