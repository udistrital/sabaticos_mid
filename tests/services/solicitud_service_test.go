package services_test

import (
	"errors"
	"testing"

	"bou.ke/monkey"
	"github.com/udistrital/sabaticos_mid/clients"
	"github.com/udistrital/sabaticos_mid/models"
	"github.com/udistrital/sabaticos_mid/service"
)

func TestCambiarEstado(t *testing.T) {
	t.Run("Ok_SinEstadoSoporte", func(t *testing.T) {
		request := models.SolicitudAprobarRechazarRequest{
			SolicitudId:     10,
			TerceroId:       20,
			EstadoSolicitud: "APROBADA",
			Justificacion:   "Cambio de estado de prueba",
			EstadoSoporte:   "",
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
	})
	t.Run("Ok_ConEstadoSoporte", func(t *testing.T) {
		request := models.SolicitudAprobarRechazarRequest{
			SolicitudId:     10,
			TerceroId:       20,
			EstadoSolicitud: "APROBADA",
			Justificacion:   "Prueba",
			EstadoSoporte:   "APROBADO",
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
		monkey.Patch(clients.ConsultarSoportesSolicitud, func(id int) ([]models.SoporteSolicitud, error) {
			return []models.SoporteSolicitud{{Id: 1, DocumentoId: 101}}, nil
		})
		defer monkey.Unpatch(clients.ConsultarSoportesSolicitud)
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
	})
	t.Run("Error_EstadoInvalido", func(t *testing.T) {
		request := models.SolicitudAprobarRechazarRequest{
			EstadoSolicitud: "INVALIDO",
			EstadoSoporte:   "",
		}
		monkey.Patch(clients.ConsultarEstadoSolicitud, func(codigo string) (*models.EstadoSolicitud, error) {
			return nil, errors.New("request status not found: " + codigo)
		})
		defer monkey.Unpatch(clients.ConsultarEstadoSolicitud)
		_, err := service.CambiarEstado(request)
		if err == nil {
			t.Fatal("se esperaba error pero se obtuvo nil")
		}
	})
	t.Run("Error_SinSoportes", func(t *testing.T) {
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
		monkey.Patch(clients.ConsultarSoportesSolicitud, func(id int) ([]models.SoporteSolicitud, error) {
			return []models.SoporteSolicitud{}, nil
		})
		defer monkey.Unpatch(clients.ConsultarSoportesSolicitud)
		_, err := service.CambiarEstado(request)
		if err == nil {
			t.Fatal("se esperaba error por falta de soportes")
		}
	})
	t.Run("Error_ConsultarIdsHistorial", func(t *testing.T) {
		request := models.SolicitudAprobarRechazarRequest{
			SolicitudId:     10,
			EstadoSolicitud: "APROBADA",
			EstadoSoporte:   "",
		}
		estadoMock := &models.EstadoSolicitud{Id: 5}
		monkey.Patch(clients.ConsultarEstadoSolicitud, func(codigo string) (*models.EstadoSolicitud, error) {
			return estadoMock, nil
		})
		defer monkey.Unpatch(clients.ConsultarEstadoSolicitud)
		monkey.Patch(clients.ConsultarIdsHistorialSolicitud, func(solicitudId int) ([]int, error) {
			return nil, errors.New("error consultando historiales")
		})
		defer monkey.Unpatch(clients.ConsultarIdsHistorialSolicitud)
		_, err := service.CambiarEstado(request)
		if err == nil {
			t.Fatal("se esperaba error pero se obtuvo nil")
		}
	})
	t.Run("Error_DesactivarHistorial", func(t *testing.T) {
		request := models.SolicitudAprobarRechazarRequest{
			SolicitudId:     10,
			EstadoSolicitud: "APROBADA",
			EstadoSoporte:   "",
		}
		estadoMock := &models.EstadoSolicitud{Id: 5}
		monkey.Patch(clients.ConsultarEstadoSolicitud, func(codigo string) (*models.EstadoSolicitud, error) {
			return estadoMock, nil
		})
		defer monkey.Unpatch(clients.ConsultarEstadoSolicitud)
		monkey.Patch(clients.ConsultarIdsHistorialSolicitud, func(solicitudId int) ([]int, error) {
			return []int{1, 2}, nil
		})
		defer monkey.Unpatch(clients.ConsultarIdsHistorialSolicitud)
		desactivaciones := 0
		monkey.Patch(clients.DesactivarHistorialSolicitud, func(idHistorial int) (bool, error) {
			desactivaciones++
			if desactivaciones == 2 {
				return false, errors.New("error desactivando historial")
			}
			return true, nil
		})
		defer monkey.Unpatch(clients.DesactivarHistorialSolicitud)
		_, err := service.CambiarEstado(request)
		if err == nil {
			t.Fatal("se esperaba error pero se obtuvo nil")
		}
		if desactivaciones != 2 {
			t.Fatalf("se esperaban 2 intentos de desactivación, hubo %d", desactivaciones)
		}
	})
}
