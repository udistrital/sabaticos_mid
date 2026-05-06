package services_test

import (
	"encoding/json"
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

func TestCrearSolicitud(t *testing.T) {
	t.Run("Ok_NuevaSolicitud", func(t *testing.T) {
		req := models.SolicitudRequest{
			TerceroId:       1,
			TipoSolicitudId: "NS",
			Formulario:      json.RawMessage("{}"),
		}
		tipoSolicitudMock := &models.TipoSolicitud{Id: 1, CodigoAbreviacion: "NS"}
		solicitudMock := &models.Solicitud{Id: 100}

		monkey.Patch(clients.ConsultarTipoSolicitud, func(codigo string) (*models.TipoSolicitud, error) {
			return tipoSolicitudMock, nil
		})
		defer monkey.Unpatch(clients.ConsultarTipoSolicitud)

		monkey.Patch(clients.RegistrarSolicitud, func(terceroId int, tipoSolicitudId int, sabaticoId *int) (*models.Solicitud, error) {
			return solicitudMock, nil
		})
		defer monkey.Unpatch(clients.RegistrarSolicitud)

		monkey.Patch(clients.RegistrarHistorialSolicitud, func(solicitudId int, terceroId int, justificacion string, codigoEstadoSolicitud string) (*models.HistorialSolicitud, error) {
			return &models.HistorialSolicitud{}, nil
		})
		defer monkey.Unpatch(clients.RegistrarHistorialSolicitud)

		monkey.Patch(clients.RegistrarFormularioSolicitud, func(solicitudId int, contenido string) (*models.FormularioSolicitud, error) {
			return &models.FormularioSolicitud{}, nil
		})
		defer monkey.Unpatch(clients.RegistrarFormularioSolicitud)

		resultado, err := service.CrearSolicitud(req)
		if err != nil {
			t.Fatalf("no se esperaba error, llegó: %v", err)
		}
		if resultado.Id != solicitudMock.Id {
			t.Errorf("se esperaba ID %d, llegó %d", solicitudMock.Id, resultado.Id)
		}
	})

	t.Run("Error_ValidacionSabaticoEnNueva", func(t *testing.T) {
		idSabatico := 50
		req := models.SolicitudRequest{
			TipoSolicitudId: "NS",
			SabaticoId:      &idSabatico,
		}
		monkey.Patch(clients.ConsultarTipoSolicitud, func(codigo string) (*models.TipoSolicitud, error) {
			return &models.TipoSolicitud{CodigoAbreviacion: "NS"}, nil
		})
		defer monkey.Unpatch(clients.ConsultarTipoSolicitud)

		_, err := service.CrearSolicitud(req)
		if err == nil || err.Error() != "a NEW request cannot be created with an associated Sabbatical" {
			t.Fatal("se esperaba error de validación para solicitud NUEVA con sabático")
		}
	})

}

func TestGetFormulariosByDocumentoId(t *testing.T) {
	t.Run("Ok_FiltradoPorFacultad", func(t *testing.T) {
		docId := "12345"
		estados := []string{"S1", "S2"}
		personaMock := &models.Persona{Dependencia: "FACULTAD_INGENIERIA"}
		formulariosMock := []models.FormularioSolicitud{
			{
				Id:          1,
				Contenido:   `{"docente":{"facultad":"FACULTAD_INGENIERIA"}}`,
				SolicitudId: models.IdReference{Id: 10},
			},
			{
				Id:          2,
				Contenido:   `{"docente":{"facultad":"OTRA_FACULTAD"}}`,
				SolicitudId: models.IdReference{Id: 11},
			},
		}

		monkey.Patch(clients.ConsultarSecretariaAcademicaDocumentoUserId, func(id string) (*models.Persona, error) {
			return personaMock, nil
		})
		defer monkey.Unpatch(clients.ConsultarSecretariaAcademicaDocumentoUserId)

		monkey.Patch(clients.ConsultarTodosFormulariosSolicitud, func() ([]models.FormularioSolicitud, error) {
			return formulariosMock, nil
		})
		defer monkey.Unpatch(clients.ConsultarTodosFormulariosSolicitud)

		monkey.Patch(clients.ConsultarIdsHistorialSolicitud, func(id int) ([]int, error) {
			return []int{100}, nil
		})
		defer monkey.Unpatch(clients.ConsultarIdsHistorialSolicitud)

		monkey.Patch(clients.ConsultarHistorialSolicitudIdEstadoId, func(id int, est []string) ([]models.HistorialSolicitud, error) {
			return []models.HistorialSolicitud{{Id: 100, SolicitudId: models.IdReference{Id: 10}}}, nil
		})
		defer monkey.Unpatch(clients.ConsultarHistorialSolicitudIdEstadoId)

		resultados, err := service.GetFormulariosByDocumentoId(docId, estados)
		if err != nil {
			t.Fatalf("error inesperado: %v", err)
		}
		if len(resultados) != 1 {
			t.Errorf("se esperaba 1 resultado filtrado, llegaron %d", len(resultados))
		}
	})

	t.Run("Error_ConsultaSecretariaFalla", func(t *testing.T) {
		monkey.Patch(clients.ConsultarSecretariaAcademicaDocumentoUserId, func(id string) (*models.Persona, error) {
			return nil, errors.New("error de red")
		})
		defer monkey.Unpatch(clients.ConsultarSecretariaAcademicaDocumentoUserId)

		_, err := service.GetFormulariosByDocumentoId("123", nil)
		if err == nil {
			t.Fatal("se esperaba error al fallar consulta de secretaria")
		}
	})

	t.Run("Error_JsonInvalidoEnFormulario", func(t *testing.T) {
		personaMock := &models.Persona{Dependencia: "INGENIERIA"}
		formulariosMock := []models.FormularioSolicitud{
			{Id: 1, Contenido: "no-es-json"},
		}

		monkey.Patch(clients.ConsultarSecretariaAcademicaDocumentoUserId, func(id string) (*models.Persona, error) {
			return personaMock, nil
		})
		defer monkey.Unpatch(clients.ConsultarSecretariaAcademicaDocumentoUserId)
		monkey.Patch(clients.ConsultarTodosFormulariosSolicitud, func() ([]models.FormularioSolicitud, error) {
			return formulariosMock, nil
		})
		defer monkey.Unpatch(clients.ConsultarTodosFormulariosSolicitud)

		_, err := service.GetFormulariosByDocumentoId("123", nil)
		if err == nil {
			t.Fatal("se esperaba error por JSON inválido")
		}
	})

	t.Run("Error_ConsultaFormulariosFalla", func(t *testing.T) {
		monkey.Patch(clients.ConsultarSecretariaAcademicaDocumentoUserId, func(id string) (*models.Persona, error) {
			return &models.Persona{Dependencia: "OK"}, nil
		})
		defer monkey.Unpatch(clients.ConsultarSecretariaAcademicaDocumentoUserId)
		monkey.Patch(clients.ConsultarTodosFormulariosSolicitud, func() ([]models.FormularioSolicitud, error) { return nil, errors.New("fail") })
		defer monkey.Unpatch(clients.ConsultarTodosFormulariosSolicitud)
		if _, err := service.GetFormulariosByDocumentoId("123", nil); err == nil {
			t.Fatal("fail")
		}
	})
}

func TestRadicarSolicitud(t *testing.T) {
	t.Run("Ok_RadicacionExitosa", func(t *testing.T) {
		req := models.RadicarSolicitudRequest{
			SolicitudId:  1,
			FormularioId: 2,
			Formulario:   json.RawMessage("{}"),
			DocumentosId: []int{301},
		}

		monkey.Patch(clients.ConsultarSolicitud, func(id int) (*models.Solicitud, error) {
			return &models.Solicitud{Id: 1, TerceroId: 10}, nil
		})
		defer monkey.Unpatch(clients.ConsultarSolicitud)

		monkey.Patch(clients.ConsultarIdsHistorialSolicitud, func(id int) ([]int, error) { return []int{5}, nil })
		defer monkey.Unpatch(clients.ConsultarIdsHistorialSolicitud)
		monkey.Patch(clients.DesactivarHistorialSolicitud, func(id int) (bool, error) { return true, nil })
		defer monkey.Unpatch(clients.DesactivarHistorialSolicitud)
		monkey.Patch(clients.RegistrarHistorialSolicitud, func(s, t int, j, c string) (*models.HistorialSolicitud, error) {
			return &models.HistorialSolicitud{Id: 50}, nil
		})
		defer monkey.Unpatch(clients.RegistrarHistorialSolicitud)
		monkey.Patch(clients.ActualizarFormularioSolicitud, func(s, f int, c string) (*models.FormularioSolicitud, error) {
			return &models.FormularioSolicitud{Id: 2}, nil
		})
		defer monkey.Unpatch(clients.ActualizarFormularioSolicitud)
		monkey.Patch(clients.ActualizarSoporteSolicitud, func(doc, sol int, est string) (*models.SoporteSolicitud, error) {
			return &models.SoporteSolicitud{Id: 301}, nil
		})
		defer monkey.Unpatch(clients.ActualizarSoporteSolicitud)

		res, err := service.RadicarSolicitud(req)
		if err != nil {
			t.Fatalf("error no esperado: %v", err)
		}
		if res["solicitud"] == nil || res["historial"] == nil {
			t.Fatal("faltan datos en la respuesta de radicación")
		}
	})

	t.Run("Error_ConsultaSolicitudNoExiste", func(t *testing.T) {
		req := models.RadicarSolicitudRequest{SolicitudId: 99}
		monkey.Patch(clients.ConsultarSolicitud, func(id int) (*models.Solicitud, error) {
			return nil, errors.New("not found")
		})
		defer monkey.Unpatch(clients.ConsultarSolicitud)

		_, err := service.RadicarSolicitud(req)
		if err == nil {
			t.Fatal("se esperaba error al no encontrar la solicitud")
		}
	})

	t.Run("Error_DesactivarHistorialFalla", func(t *testing.T) {
		req := models.RadicarSolicitudRequest{SolicitudId: 1}
		monkey.Patch(clients.ConsultarSolicitud, func(id int) (*models.Solicitud, error) {
			return &models.Solicitud{Id: 1}, nil
		})
		defer monkey.Unpatch(clients.ConsultarSolicitud)
		monkey.Patch(clients.ConsultarIdsHistorialSolicitud, func(id int) ([]int, error) {
			return []int{10}, nil
		})
		defer monkey.Unpatch(clients.ConsultarIdsHistorialSolicitud)
		monkey.Patch(clients.DesactivarHistorialSolicitud, func(id int) (bool, error) {
			return false, errors.New("db error")
		})
		defer monkey.Unpatch(clients.DesactivarHistorialSolicitud)

		_, err := service.RadicarSolicitud(req)
		if err == nil {
			t.Fatal("se esperaba error al fallar desactivación de historial")
		}
	})

	t.Run("Error_ActualizarFormularioFalla", func(t *testing.T) {
		req := models.RadicarSolicitudRequest{SolicitudId: 1, FormularioId: 2}
		monkey.Patch(clients.ConsultarSolicitud, func(id int) (*models.Solicitud, error) { return &models.Solicitud{Id: 1}, nil })
		defer monkey.Unpatch(clients.ConsultarSolicitud)
		monkey.Patch(clients.ConsultarIdsHistorialSolicitud, func(id int) ([]int, error) { return []int{}, nil })
		defer monkey.Unpatch(clients.ConsultarIdsHistorialSolicitud)
		monkey.Patch(clients.RegistrarHistorialSolicitud, func(s, t int, j, c string) (*models.HistorialSolicitud, error) {
			return &models.HistorialSolicitud{}, nil
		})
		defer monkey.Unpatch(clients.RegistrarHistorialSolicitud)

		monkey.Patch(clients.ActualizarFormularioSolicitud, func(s, f int, c string) (*models.FormularioSolicitud, error) {
			return nil, errors.New("update fail")
		})
		defer monkey.Unpatch(clients.ActualizarFormularioSolicitud)

		_, err := service.RadicarSolicitud(req)
		if err == nil {
			t.Fatal("se esperaba error al fallar actualización de formulario")
		}
	})
}
