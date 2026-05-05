package services_test

import (
	"errors"
	"testing"
	"bou.ke/monkey"
	"github.com/udistrital/sabaticos_mid/clients"
	"github.com/udistrital/sabaticos_mid/models"
	"github.com/udistrital/sabaticos_mid/service"
)

// TestCambiarEstado_Ok_SinEstadoSoporte prueba el caso exitoso cuando NO se proporciona EstadoSoporte.
// Flujo: consultar estado → desactivar historiales existentes → registrar nuevo historial.
func TestCambiarEstado_Ok_SinEstadoSoporte(t *testing.T) {
	request := models.SolicitudAprobarRechazarRequest{
		SolicitudId:     10,
		TerceroId:       20,
		EstadoSolicitud: "APROBADA",
		Justificacion:   "Cambio de estado de prueba",
		EstadoSoporte:   "", // vacío = no procesa soportes
	}
	estadoMock := &models.EstadoSolicitud{Id: 5}
	historialMock := &models.HistorialSolicitud{Id: 100}
	monkey.Patch(clients.ConsultarEstadoSolicitud, func(codigo string) (*models.EstadoSolicitud, error) {
		return estadoMock, nil
	})
	defer monkey.Unpatch(clients.ConsultarEstadoSolicitud)
	monkey.Patch(clients.ConsultarIdsHistorialSolicitud, func(solicitudId int) ([]int, error) {
		return []int{1, 2}, nil
	})
	defer monkey.Unpatch(clients.ConsultarIdsHistorialSolicitud)
	// CORREGIDO: retorna (bool, error) como la función real
	monkey.Patch(clients.DesactivarHistorialSolicitud, func(idHistorial int) (bool, error) {
		return true, nil
	})
	defer monkey.Unpatch(clients.DesactivarHistorialSolicitud)
	monkey.Patch(clients.RegistrarHistorialSolicitudEstado, func(solicitudId, terceroId int, justificacion string, estadoId int) (*models.HistorialSolicitud, error) {
		return historialMock, nil
	})
	defer monkey.Unpatch(clients.RegistrarHistorialSolicitudEstado)
	resultado, err := service.CambiarEstado(request)
	if err != nil {
		t.Fatalf("no se esperaba error, llegó: %v", err)
	}
	if resultado.Id != historialMock.Id {
		t.Fatalf("se esperaba ID %d y llegó %d", historialMock.Id, resultado.Id)
	}
}

// TestCambiarEstado_Ok_ConEstadoSoporte prueba el flujo completo CON procesamiento de soportes.
func TestCambiarEstado_Ok_ConEstadoSoporte(t *testing.T) {
	request := models.SolicitudAprobarRechazarRequest{
		SolicitudId:     10,
		TerceroId:       20,
		EstadoSolicitud: "APROBADA",
		Justificacion:   "Prueba",
		EstadoSoporte:   "APROBADO", // con valor = sí procesa soportes
	}
	estadoMock := &models.EstadoSolicitud{Id: 5}
	historialMock := &models.HistorialSolicitud{Id: 100}
	monkey.Patch(clients.ConsultarEstadoSolicitud, func(codigo string) (*models.EstadoSolicitud, error) {
		return estadoMock, nil
	})
	defer monkey.Unpatch(clients.ConsultarEstadoSolicitud)
	monkey.Patch(clients.ConsultarIdsHistorialSolicitud, func(solicitudId int) ([]int, error) {
		return []int{1}, nil
	})
	defer monkey.Unpatch(clients.ConsultarIdsHistorialSolicitud)
	monkey.Patch(clients.DesactivarHistorialSolicitud, func(idHistorial int) (bool, error) {
		return true, nil
	})
	defer monkey.Unpatch(clients.DesactivarHistorialSolicitud)
	// Mock adicional: consultar soportes asociados
	monkey.Patch(clients.ConsultarSoportesSolicitud, func(id int) ([]models.SoporteSolicitud, error) {
		return []models.SoporteSolicitud{{Id: 1, DocumentoId: 101}}, nil
	})
	defer monkey.Unpatch(clients.ConsultarSoportesSolicitud)
	// Mock adicional: actualizar cada soporte
	monkey.Patch(clients.ActualizarSoporteSolicitud, func(soporteId, solicitudId int, estado string) (*models.SoporteSolicitud, error) {
		return &models.SoporteSolicitud{Id: soporteId}, nil
	})
	defer monkey.Unpatch(clients.ActualizarSoporteSolicitud)
	monkey.Patch(clients.RegistrarHistorialSolicitudEstado, func(solicitudId, terceroId int, justificacion string, estadoId int) (*models.HistorialSolicitud, error) {
		return historialMock, nil
	})
	defer monkey.Unpatch(clients.RegistrarHistorialSolicitudEstado)
	resultado, err := service.CambiarEstado(request)
	if err != nil {
		t.Fatalf("no se esperaba error, llegó: %v", err)
	}
	if resultado.Id != historialMock.Id {
		t.Fatalf("se esperaba ID %d y llegó %d", historialMock.Id, resultado.Id)
	}
}

// TestCambiarEstado_Error_EstadoInvalido verifica que retorne error con estado inválido.
func TestCambiarEstado_Error_EstadoInvalido(t *testing.T) {
	request := models.SolicitudAprobarRechazarRequest{
		EstadoSolicitud: "INVALIDO",
		EstadoSoporte:   "",
	}
	monkey.Patch(clients.ConsultarEstadoSolicitud, func(codigo string) (*models.EstadoSolicitud, error) {
		return nil, errors.New("request status not found: " + codigo)
	})
	defer monkey.Unpatch(clients.ConsultarEstadoSolicitud)
	resultado, err := service.CambiarEstado(request)
	if err == nil {
		t.Fatal("se esperaba error pero se obtuvo nil")
	}
	if resultado != nil {
		t.Fatal("se esperaba nil")
	}
}

// TestCambiarEstado_Error_SinSoportes verifica que falle cuando EstadoSoporte tiene valor
// pero no hay soportes asociados (la función retorna error en ese caso).
func TestCambiarEstado_Error_SinSoportes(t *testing.T) {
	request := models.SolicitudAprobarRechazarRequest{
		SolicitudId:     10,
		EstadoSolicitud: "APROBADA",
		EstadoSoporte:   "APROBADO",
	}
	estadoMock := &models.EstadoSolicitud{Id: 5}
	monkey.Patch(clients.ConsultarEstadoSolicitud, func(codigo string) (*models.EstadoSolicitud, error) {
		return estadoMock, nil
	})
	defer monkey.Unpatch(clients.ConsultarEstadoSolicitud)
	monkey.Patch(clients.ConsultarIdsHistorialSolicitud, func(solicitudId int) ([]int, error) {
		return []int{1}, nil
	})
	defer monkey.Unpatch(clients.ConsultarIdsHistorialSolicitud)
	monkey.Patch(clients.DesactivarHistorialSolicitud, func(idHistorial int) (bool, error) {
		return true, nil
	})
	defer monkey.Unpatch(clients.DesactivarHistorialSolicitud)
	// Retorna slice vacío — la función debe detectar que no hay soportes
	monkey.Patch(clients.ConsultarSoportesSolicitud, func(id int) ([]models.SoporteSolicitud, error) {
		return []models.SoporteSolicitud{}, nil
	})
	defer monkey.Unpatch(clients.ConsultarSoportesSolicitud)
	_, err := service.CambiarEstado(request)
	if err == nil {
    	t.Fatal("se esperaba error por falta de soportes")
	}
}