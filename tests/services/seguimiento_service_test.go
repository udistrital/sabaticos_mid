package services_test

import (
	"errors"
	"testing"

	"bou.ke/monkey"
	"github.com/udistrital/sabaticos_mid/clients"
	"github.com/udistrital/sabaticos_mid/models"
	"github.com/udistrital/sabaticos_mid/service"
)

func TestConsultarSoportesSabaticos(t *testing.T) {
	t.Run("Ok_ConsultarSoportes", func(t *testing.T) {
		soportesMock := []models.SoporteSabatico{
			{Id: 1, DocumentoId: 101, EstadoSoporteSabaticoId: models.EstadoSabatico{Id: 1}},
			{Id: 2, DocumentoId: 102, EstadoSoporteSabaticoId: models.EstadoSabatico{Id: 1}},
		}
		gestorDocumentalMock := &models.GestorDocumental{Id: 201, Nombre: "documento1.pdf"}

		monkey.Patch(clients.ConsultarSoportesSabaticos, func(sabaticoId int) ([]models.SoporteSabatico, error) {
			return soportesMock, nil
		})
		defer monkey.Unpatch(clients.ConsultarSoportesSabaticos)

		monkey.Patch(clients.ConsultarGestorDocumental, func(documentoId int) (*models.GestorDocumental, error) {
			return gestorDocumentalMock, nil
		})
		defer monkey.Unpatch(clients.ConsultarGestorDocumental)

		resultado, err := service.ConsultarSoportesSabaticos(100)
		if err != nil {
			t.Fatalf("no se esperaba error, llegó: %v", err)
		}
		if len(resultado) != 2 {
			t.Fatalf("se esperaban 2 soportes, llegó %d", len(resultado))
		}
		if resultado[0]["Id"] != soportesMock[0].Id {
			t.Fatalf("se esperaba Id %d y llegó %d", soportesMock[0].Id, resultado[0]["Id"])
		}
	})

	t.Run("Error_SoporteIdInvalido", func(t *testing.T) {
		_, err := service.ConsultarSoportesSabaticos(0)
		if err == nil {
			t.Fatal("se esperaba error por sabaticoId inválido")
		}
	})

	t.Run("Error_ConsultarSoportesFallido", func(t *testing.T) {
		monkey.Patch(clients.ConsultarSoportesSabaticos, func(sabaticoId int) ([]models.SoporteSabatico, error) {
			return nil, errors.New("error consultando soportes")
		})
		defer monkey.Unpatch(clients.ConsultarSoportesSabaticos)

		_, err := service.ConsultarSoportesSabaticos(100)
		if err == nil {
			t.Fatal("se esperaba error al fallar la consulta de soportes")
		}
	})

	t.Run("Error_ConsultarGestorDocumentalFallido", func(t *testing.T) {
		soportesMock := []models.SoporteSabatico{
			{Id: 1, DocumentoId: 101, EstadoSoporteSabaticoId: models.EstadoSabatico{Id: 1}},
		}

		monkey.Patch(clients.ConsultarSoportesSabaticos, func(sabaticoId int) ([]models.SoporteSabatico, error) {
			return soportesMock, nil
		})
		defer monkey.Unpatch(clients.ConsultarSoportesSabaticos)

		monkey.Patch(clients.ConsultarGestorDocumental, func(documentoId int) (*models.GestorDocumental, error) {
			return nil, errors.New("error consultando gestor documental")
		})
		defer monkey.Unpatch(clients.ConsultarGestorDocumental)

		_, err := service.ConsultarSoportesSabaticos(100)
		if err == nil {
			t.Fatal("se esperaba error al fallar la consulta de gestor documental")
		}
	})
}

func TestCrearSabatico(t *testing.T) {
	t.Run("Ok_CrearSabatico", func(t *testing.T) {
		sabaticoMock := &models.CrearSabaticoResult{Id: 100}

		monkey.Patch(clients.ConsultarEstadoSolicitud, func(codigo string) (*models.EstadoSolicitud, error) {
			return &models.EstadoSolicitud{Id: 12}, nil
		})
		defer monkey.Unpatch(clients.ConsultarEstadoSolicitud)

		monkey.Patch(clients.ConsultarIdsHistorialSolicitud, func(solicitudId int) ([]int, error) {
			return []int{}, nil
		})
		defer monkey.Unpatch(clients.ConsultarIdsHistorialSolicitud)

		monkey.Patch(clients.ConsultarSoportesSolicitud, func(id int) ([]models.SoporteSolicitud, error) {
			return []models.SoporteSolicitud{{Id: 1, DocumentoId: 101}}, nil
		})
		defer monkey.Unpatch(clients.ConsultarSoportesSolicitud)

		monkey.Patch(clients.ActualizarSoporteSolicitud, func(soporteId, solicitudId int, estado string) (*models.SoporteSolicitud, error) {
			return &models.SoporteSolicitud{Id: soporteId}, nil
		})
		defer monkey.Unpatch(clients.ActualizarSoporteSolicitud)

		monkey.Patch(clients.RegistrarHistorialSolicitudEstado, func(solicitudId, terceroId int, justificacion string, estadoId int) (*models.HistorialSolicitud, error) {
			return &models.HistorialSolicitud{Id: 50}, nil
		})
		defer monkey.Unpatch(clients.RegistrarHistorialSolicitudEstado)

		monkey.Patch(clients.RegistrarSabatico, func(solicitudId int, terceroId int, observaciones string, fechaInicio string, fechaFin string, estadoSabatico string) (*models.CrearSabaticoResult, error) {
			return sabaticoMock, nil
		})
		defer monkey.Unpatch(clients.RegistrarSabatico)

		monkey.Patch(clients.AsociarSabaticoSolicitud, func(solicitudId int, sabaticoId int) error {
			return nil
		})
		defer monkey.Unpatch(clients.AsociarSabaticoSolicitud)

		resultado, err := service.CrearSabatico(10, 20, "obs test", "2024-01-01", "2025-01-01")
		if err != nil {
			t.Fatalf("no se esperaba error, llegó: %v", err)
		}
		if resultado.Id != sabaticoMock.Id {
			t.Fatalf("se esperaba ID %d y llegó %d", sabaticoMock.Id, resultado.Id)
		}
	})

	t.Run("Error_CambiarEstadoFalla", func(t *testing.T) {
		monkey.Patch(clients.ConsultarEstadoSolicitud, func(codigo string) (*models.EstadoSolicitud, error) {
			return nil, errors.New("error consultando estado solicitud")
		})
		defer monkey.Unpatch(clients.ConsultarEstadoSolicitud)

		_, err := service.CrearSabatico(10, 20, "obs", "2024-01-01", "2025-01-01")
		if err == nil {
			t.Fatal("se esperaba error al fallar cambio de estado")
		}
	})

	t.Run("Error_RegistrarSabaticoFalla", func(t *testing.T) {
		monkey.Patch(clients.ConsultarEstadoSolicitud, func(codigo string) (*models.EstadoSolicitud, error) {
			return &models.EstadoSolicitud{Id: 12}, nil
		})
		defer monkey.Unpatch(clients.ConsultarEstadoSolicitud)

		monkey.Patch(clients.ConsultarIdsHistorialSolicitud, func(solicitudId int) ([]int, error) {
			return []int{}, nil
		})
		defer monkey.Unpatch(clients.ConsultarIdsHistorialSolicitud)

		monkey.Patch(clients.ConsultarSoportesSolicitud, func(id int) ([]models.SoporteSolicitud, error) {
			return []models.SoporteSolicitud{{Id: 1, DocumentoId: 101}}, nil
		})
		defer monkey.Unpatch(clients.ConsultarSoportesSolicitud)

		monkey.Patch(clients.ActualizarSoporteSolicitud, func(soporteId, solicitudId int, estado string) (*models.SoporteSolicitud, error) {
			return &models.SoporteSolicitud{Id: soporteId}, nil
		})
		defer monkey.Unpatch(clients.ActualizarSoporteSolicitud)

		monkey.Patch(clients.RegistrarHistorialSolicitudEstado, func(solicitudId, terceroId int, justificacion string, estadoId int) (*models.HistorialSolicitud, error) {
			return &models.HistorialSolicitud{Id: 50}, nil
		})
		defer monkey.Unpatch(clients.RegistrarHistorialSolicitudEstado)

		monkey.Patch(clients.RegistrarSabatico, func(solicitudId int, terceroId int, observaciones string, fechaInicio string, fechaFin string, estadoSabatico string) (*models.CrearSabaticoResult, error) {
			return nil, errors.New("error registrando sabático")
		})
		defer monkey.Unpatch(clients.RegistrarSabatico)

		_, err := service.CrearSabatico(10, 20, "obs", "2024-01-01", "2025-01-01")
		if err == nil {
			t.Fatal("se esperaba error al fallar el registro del sabático")
		}
	})

	t.Run("Error_AsociarSabaticoSolicitudFalla", func(t *testing.T) {
		sabaticoMock := &models.CrearSabaticoResult{Id: 100}

		monkey.Patch(clients.ConsultarEstadoSolicitud, func(codigo string) (*models.EstadoSolicitud, error) {
			return &models.EstadoSolicitud{Id: 12}, nil
		})
		defer monkey.Unpatch(clients.ConsultarEstadoSolicitud)

		monkey.Patch(clients.ConsultarIdsHistorialSolicitud, func(solicitudId int) ([]int, error) {
			return []int{}, nil
		})
		defer monkey.Unpatch(clients.ConsultarIdsHistorialSolicitud)

		monkey.Patch(clients.ConsultarSoportesSolicitud, func(id int) ([]models.SoporteSolicitud, error) {
			return []models.SoporteSolicitud{{Id: 1, DocumentoId: 101}}, nil
		})
		defer monkey.Unpatch(clients.ConsultarSoportesSolicitud)

		monkey.Patch(clients.ActualizarSoporteSolicitud, func(soporteId, solicitudId int, estado string) (*models.SoporteSolicitud, error) {
			return &models.SoporteSolicitud{Id: soporteId}, nil
		})
		defer monkey.Unpatch(clients.ActualizarSoporteSolicitud)

		monkey.Patch(clients.RegistrarHistorialSolicitudEstado, func(solicitudId, terceroId int, justificacion string, estadoId int) (*models.HistorialSolicitud, error) {
			return &models.HistorialSolicitud{Id: 50}, nil
		})
		defer monkey.Unpatch(clients.RegistrarHistorialSolicitudEstado)

		monkey.Patch(clients.RegistrarSabatico, func(solicitudId int, terceroId int, observaciones string, fechaInicio string, fechaFin string, estadoSabatico string) (*models.CrearSabaticoResult, error) {
			return sabaticoMock, nil
		})
		defer monkey.Unpatch(clients.RegistrarSabatico)

		monkey.Patch(clients.AsociarSabaticoSolicitud, func(solicitudId int, sabaticoId int) error {
			return errors.New("error asociando sabático a solicitud")
		})
		defer monkey.Unpatch(clients.AsociarSabaticoSolicitud)

		_, err := service.CrearSabatico(10, 20, "obs", "2024-01-01", "2025-01-01")
		if err == nil {
			t.Fatal("se esperaba error al fallar la asociación del sabático")
		}
	})
}

func TestGuardarPlanTrabajoSabatico(t *testing.T) {

	t.Run("Ok_ActualizarPlanTrabajoExistente", func(t *testing.T) {

		historial := models.HistorialEstadoSabatico{
			Id: 1,
			EstadoSabaticoId: models.EstadoSabatico{
				CodigoAbreviacion: "ES1",
			},
			Justificacion: "Anterior",
		}

		monkey.Patch(
			clients.ConsultarTodosHistorialEstadoSabatico,
			func(id int) ([]models.HistorialEstadoSabatico, error) {
				return []models.HistorialEstadoSabatico{historial}, nil
			},
		)
		defer monkey.Unpatch(clients.ConsultarTodosHistorialEstadoSabatico)

		monkey.Patch(
			clients.ActualizarHistorialEstadoSabatico,
			func(historial models.HistorialEstadoSabatico) error {
				return nil
			},
		)
		defer monkey.Unpatch(clients.ActualizarHistorialEstadoSabatico)

		resp, err := service.GuardarPlanTrabajoSabatico(
			models.PlanTrabajoSabaticoRequest{
				SabaticoId:    1,
				Justificacion: "Nueva Justificación",
			},
		)

		if err != nil {
			t.Fatalf("no se esperaba error: %v", err)
		}

		if resp.Justificacion != "Nueva Justificación" {
			t.Fatalf("se esperaba justificación actualizada")
		}
	})

	t.Run("Error_ActualizarHistorial", func(t *testing.T) {

		historial := models.HistorialEstadoSabatico{
			Id: 1,
			EstadoSabaticoId: models.EstadoSabatico{
				CodigoAbreviacion: "ES1",
			},
		}

		monkey.Patch(
			clients.ConsultarTodosHistorialEstadoSabatico,
			func(id int) ([]models.HistorialEstadoSabatico, error) {
				return []models.HistorialEstadoSabatico{historial}, nil
			},
		)
		defer monkey.Unpatch(clients.ConsultarTodosHistorialEstadoSabatico)

		monkey.Patch(
			clients.ActualizarHistorialEstadoSabatico,
			func(historial models.HistorialEstadoSabatico) error {
				return errors.New("error actualizando historial")
			},
		)
		defer monkey.Unpatch(clients.ActualizarHistorialEstadoSabatico)

		_, err := service.GuardarPlanTrabajoSabatico(
			models.PlanTrabajoSabaticoRequest{
				SabaticoId: 1,
			},
		)

		if err == nil {
			t.Fatal("se esperaba error")
		}
	})

	t.Run("Ok_CrearNuevoPlanTrabajo", func(t *testing.T) {

		historiales := []models.HistorialEstadoSabatico{
			{
				Id:        1,
				TerceroId: 100,
				Activo:    true,
				EstadoSabaticoId: models.EstadoSabatico{
					Id:                1,
					CodigoAbreviacion: "ES0",
				},
			},
		}

		monkey.Patch(
			clients.ConsultarTodosHistorialEstadoSabatico,
			func(id int) ([]models.HistorialEstadoSabatico, error) {
				return historiales, nil
			},
		)
		defer monkey.Unpatch(clients.ConsultarTodosHistorialEstadoSabatico)

		monkey.Patch(
			clients.DesactivarHistorialEstadoSabatico,
			func(historial models.HistorialEstadoSabatico) error {
				return nil
			},
		)
		defer monkey.Unpatch(clients.DesactivarHistorialEstadoSabatico)

		monkey.Patch(
			clients.ConsultarIdEstadoSabatico,
			func(codigo string) (int, error) {
				if codigo != "ES1" {
					t.Fatalf("se esperaba ES1 y llegó %s", codigo)
				}
				return 5, nil
			},
		)
		defer monkey.Unpatch(clients.ConsultarIdEstadoSabatico)

		monkey.Patch(
			clients.CrearHistorialEstadoSabatico,
			func(
				terceroId int,
				justificacion string,
				estadoId int,
				sabaticoId int,
			) (*models.HistorialEstadoSabatico, error) {

				if terceroId != 100 {
					t.Fatalf("se esperaba terceroId 100 y llegó %d", terceroId)
				}

				if estadoId != 5 {
					t.Fatalf("se esperaba estadoId 5 y llegó %d", estadoId)
				}

				if sabaticoId != 1 {
					t.Fatalf("se esperaba sabaticoId 1 y llegó %d", sabaticoId)
				}

				return &models.HistorialEstadoSabatico{
					Id:            99,
					TerceroId:     terceroId,
					Justificacion: justificacion,
				}, nil
			},
		)
		defer monkey.Unpatch(clients.CrearHistorialEstadoSabatico)

		resp, err := service.GuardarPlanTrabajoSabatico(
			models.PlanTrabajoSabaticoRequest{
				SabaticoId:    1,
				Justificacion: "Plan nuevo",
			},
		)

		if err != nil {
			t.Fatalf("no se esperaba error: %v", err)
		}

		if resp == nil {
			t.Fatal("respuesta nil")
		}

		if resp.Id != 99 {
			t.Fatalf("se esperaba id 99 y llegó %d", resp.Id)
		}
	})

	t.Run("Error_ConsultarHistoriales", func(t *testing.T) {

		monkey.Patch(
			clients.ConsultarTodosHistorialEstadoSabatico,
			func(id int) ([]models.HistorialEstadoSabatico, error) {
				return nil, errors.New("error")
			},
		)
		defer monkey.Unpatch(clients.ConsultarTodosHistorialEstadoSabatico)

		_, err := service.GuardarPlanTrabajoSabatico(
			models.PlanTrabajoSabaticoRequest{
				SabaticoId: 1,
			},
		)

		if err == nil {
			t.Fatal("se esperaba error")
		}
	})

	t.Run("Error_DesactivarHistorial", func(t *testing.T) {

		historiales := []models.HistorialEstadoSabatico{
			{
				Id:        1,
				TerceroId: 100,
				EstadoSabaticoId: models.EstadoSabatico{
					CodigoAbreviacion: "ES0",
				},
			},
		}

		monkey.Patch(
			clients.ConsultarTodosHistorialEstadoSabatico,
			func(id int) ([]models.HistorialEstadoSabatico, error) {
				return historiales, nil
			},
		)
		defer monkey.Unpatch(clients.ConsultarTodosHistorialEstadoSabatico)

		monkey.Patch(
			clients.DesactivarHistorialEstadoSabatico,
			func(historial models.HistorialEstadoSabatico) error {
				return errors.New("error")
			},
		)
		defer monkey.Unpatch(clients.DesactivarHistorialEstadoSabatico)

		_, err := service.GuardarPlanTrabajoSabatico(
			models.PlanTrabajoSabaticoRequest{
				SabaticoId: 1,
			},
		)

		if err == nil {
			t.Fatal("se esperaba error")
		}
	})
}

func TestCambiarEstadoPlanTrabajoSabatico(t *testing.T) {

	t.Run("Ok_CambiarEstadoSinSoportes", func(t *testing.T) {

		historiales := []models.HistorialEstadoSabatico{
			{
				Id:        1,
				TerceroId: 123,
			},
		}

		monkey.Patch(
			clients.ConsultarHistorialEstadoSabatico,
			func(id int) ([]models.HistorialEstadoSabatico, error) {
				return historiales, nil
			},
		)
		defer monkey.Unpatch(clients.ConsultarHistorialEstadoSabatico)

		monkey.Patch(
			clients.ConsultarIdEstadoSabatico,
			func(codigo string) (int, error) {
				return 5, nil
			},
		)
		defer monkey.Unpatch(clients.ConsultarIdEstadoSabatico)

		monkey.Patch(
			clients.DesactivarHistorialEstadoSabatico,
			func(historial models.HistorialEstadoSabatico) error {
				return nil
			},
		)
		defer monkey.Unpatch(clients.DesactivarHistorialEstadoSabatico)

		monkey.Patch(
			clients.CrearHistorialEstadoSabatico,
			func(terceroId int, justificacion string, estadoId int, sabaticoId int) (*models.HistorialEstadoSabatico, error) {
				return &models.HistorialEstadoSabatico{
					Id: 10,
				}, nil
			},
		)
		defer monkey.Unpatch(clients.CrearHistorialEstadoSabatico)

		resp, err := service.CambiarEstadoPlanTrabajoSabatico(
			models.AprobarRechazarPlanTRabajoSabaticoequest{
				SabaticoId:     1,
				EstadoSabatico: "ES2",
				Justificacion:  "Aprobado",
			},
		)

		if err != nil {
			t.Fatalf("no se esperaba error: %v", err)
		}

		if resp.Id != 10 {
			t.Fatalf("resultado incorrecto")
		}
	})

	t.Run("Error_SinHistoriales", func(t *testing.T) {

		monkey.Patch(
			clients.ConsultarHistorialEstadoSabatico,
			func(id int) ([]models.HistorialEstadoSabatico, error) {
				return []models.HistorialEstadoSabatico{}, nil
			},
		)
		defer monkey.Unpatch(clients.ConsultarHistorialEstadoSabatico)

		monkey.Patch(
			clients.ConsultarIdEstadoSabatico,
			func(codigo string) (int, error) {
				return 5, nil
			},
		)
		defer monkey.Unpatch(clients.ConsultarIdEstadoSabatico)

		_, err := service.CambiarEstadoPlanTrabajoSabatico(
			models.AprobarRechazarPlanTRabajoSabaticoequest{
				SabaticoId:     1,
				EstadoSabatico: "ES2",
			},
		)

		if err == nil {
			t.Fatal("se esperaba error")
		}
	})

	t.Run("Ok_CambiarEstadoConActualizacionSoportes", func(t *testing.T) {

		historiales := []models.HistorialEstadoSabatico{
			{
				Id:        1,
				TerceroId: 123,
			},
		}

		soportes := []models.SoporteSabatico{
			{
				Id: 1,
			},
		}

		monkey.Patch(clients.ConsultarHistorialEstadoSabatico,
			func(id int) ([]models.HistorialEstadoSabatico, error) {
				return historiales, nil
			})
		defer monkey.Unpatch(clients.ConsultarHistorialEstadoSabatico)

		monkey.Patch(clients.ConsultarIdEstadoSabatico,
			func(codigo string) (int, error) {
				return 5, nil
			})
		defer monkey.Unpatch(clients.ConsultarIdEstadoSabatico)

		monkey.Patch(clients.DesactivarHistorialEstadoSabatico,
			func(historial models.HistorialEstadoSabatico) error {
				return nil
			})
		defer monkey.Unpatch(clients.DesactivarHistorialEstadoSabatico)

		monkey.Patch(
			clients.ConsultarEstadoSoporteSabatico,
			func(codigo string) (*models.EstadoSoporteSabatico, error) {
				return &models.EstadoSoporteSabatico{
					Id: 2,
				}, nil
			},
		)
		defer monkey.Unpatch(clients.ConsultarEstadoSoporteSabatico)

		monkey.Patch(clients.ConsultarSoportesSabaticos,
			func(id int) ([]models.SoporteSabatico, error) {
				return soportes, nil
			})
		defer monkey.Unpatch(clients.ConsultarSoportesSabaticos)

		monkey.Patch(clients.ActualizarSoporteSabatico,
			func(soporte models.SoporteSabatico) (*models.SoporteSabatico, error) {
				return &soporte, nil
			})
		defer monkey.Unpatch(clients.ActualizarSoporteSabatico)

		monkey.Patch(clients.CrearHistorialEstadoSabatico,
			func(terceroId int, justificacion string, estadoId int, sabaticoId int) (*models.HistorialEstadoSabatico, error) {
				return &models.HistorialEstadoSabatico{
					Id: 20,
				}, nil
			})
		defer monkey.Unpatch(clients.CrearHistorialEstadoSabatico)

		_, err := service.CambiarEstadoPlanTrabajoSabatico(
			models.AprobarRechazarPlanTRabajoSabaticoequest{
				SabaticoId:            1,
				EstadoSabatico:        "ES2",
				EstadoSoporteSabatico: "SSA",
			},
		)

		if err != nil {
			t.Fatalf("no se esperaba error: %v", err)
		}
	})

	t.Run("Error_ActualizarSoporteSabatico", func(t *testing.T) {

		historiales := []models.HistorialEstadoSabatico{
			{
				Id:        1,
				TerceroId: 123,
			},
		}

		monkey.Patch(clients.ConsultarHistorialEstadoSabatico,
			func(id int) ([]models.HistorialEstadoSabatico, error) {
				return historiales, nil
			})
		defer monkey.Unpatch(clients.ConsultarHistorialEstadoSabatico)

		monkey.Patch(clients.ConsultarIdEstadoSabatico,
			func(codigo string) (int, error) {
				return 5, nil
			})
		defer monkey.Unpatch(clients.ConsultarIdEstadoSabatico)

		monkey.Patch(clients.DesactivarHistorialEstadoSabatico,
			func(historial models.HistorialEstadoSabatico) error {
				return nil
			})
		defer monkey.Unpatch(clients.DesactivarHistorialEstadoSabatico)

		monkey.Patch(
			clients.ConsultarEstadoSoporteSabatico,
			func(codigo string) (*models.EstadoSoporteSabatico, error) {
				return &models.EstadoSoporteSabatico{
					Id: 2,
				}, nil
			},
		)
		defer monkey.Unpatch(clients.ConsultarEstadoSoporteSabatico)

		monkey.Patch(clients.ConsultarSoportesSabaticos,
			func(id int) ([]models.SoporteSabatico, error) {
				return []models.SoporteSabatico{{Id: 1}}, nil
			})
		defer monkey.Unpatch(clients.ConsultarSoportesSabaticos)

		monkey.Patch(clients.ActualizarSoporteSabatico,
			func(soporte models.SoporteSabatico) (*models.SoporteSabatico, error) {
				return nil, errors.New("error")
			})
		defer monkey.Unpatch(clients.ActualizarSoporteSabatico)

		_, err := service.CambiarEstadoPlanTrabajoSabatico(
			models.AprobarRechazarPlanTRabajoSabaticoequest{
				SabaticoId:            1,
				EstadoSabatico:        "ES2",
				EstadoSoporteSabatico: "SSA",
			},
		)

		if err == nil {
			t.Fatal("se esperaba error")
		}
	})
}
