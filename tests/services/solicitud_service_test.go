package services_test

import (
	"encoding/json"
	"errors"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/udistrital/sabaticos_mid/clients"
	"github.com/udistrital/sabaticos_mid/enums"
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
	t.Run("Error_RegistrarHistorialEstadoFalla", func(t *testing.T) {
		request := models.SolicitudAprobarRechazarRequest{
			SolicitudId:     10,
			TerceroId:       20,
			EstadoSolicitud: "APROBADA",
			EstadoSoporte:   "",
		}
		monkey.Patch(clients.ConsultarEstadoSolicitud, func(codigo string) (*models.EstadoSolicitud, error) {
			return &models.EstadoSolicitud{Id: 5}, nil
		})
		defer monkey.Unpatch(clients.ConsultarEstadoSolicitud)
		monkey.Patch(clients.ConsultarIdsHistorialSolicitud, func(solicitudId int) ([]int, error) {
			return []int{}, nil
		})
		defer monkey.Unpatch(clients.ConsultarIdsHistorialSolicitud)
		monkey.Patch(clients.RegistrarHistorialSolicitudEstado, func(solicitudId, terceroId int, justificacion string, estadoId int) (*models.HistorialSolicitud, error) {
			return nil, errors.New("error registrando historial estado")
		})
		defer monkey.Unpatch(clients.RegistrarHistorialSolicitudEstado)
		_, err := service.CambiarEstado(request)
		if err == nil {
			t.Fatal("se esperaba error al fallar el registro del historial de estado")
		}
	})
	t.Run("Error_ConsultarSoportesFalla", func(t *testing.T) {
		request := models.SolicitudAprobarRechazarRequest{
			SolicitudId:     10,
			EstadoSolicitud: "APROBADA",
			EstadoSoporte:   "APROBADO",
		}
		monkey.Patch(clients.ConsultarEstadoSolicitud, func(codigo string) (*models.EstadoSolicitud, error) {
			return &models.EstadoSolicitud{Id: 5}, nil
		})
		defer monkey.Unpatch(clients.ConsultarEstadoSolicitud)
		monkey.Patch(clients.ConsultarIdsHistorialSolicitud, func(solicitudId int) ([]int, error) {
			return []int{}, nil
		})
		defer monkey.Unpatch(clients.ConsultarIdsHistorialSolicitud)
		monkey.Patch(clients.ConsultarSoportesSolicitud, func(id int) ([]models.SoporteSolicitud, error) {
			return nil, errors.New("error consultando soportes")
		})
		defer monkey.Unpatch(clients.ConsultarSoportesSolicitud)
		_, err := service.CambiarEstado(request)
		if err == nil {
			t.Fatal("se esperaba error al fallar la consulta de soportes")
		}
	})
	t.Run("Ok_ActualizarSoporteFallaPeroContinua", func(t *testing.T) {
		request := models.SolicitudAprobarRechazarRequest{
			SolicitudId:     10,
			TerceroId:       20,
			EstadoSolicitud: "APROBADA",
			EstadoSoporte:   "APROBADO",
		}
		historialMock := &models.HistorialSolicitud{Id: 100}
		monkey.Patch(clients.ConsultarEstadoSolicitud, func(codigo string) (*models.EstadoSolicitud, error) {
			return &models.EstadoSolicitud{Id: 5}, nil
		})
		defer monkey.Unpatch(clients.ConsultarEstadoSolicitud)
		monkey.Patch(clients.ConsultarIdsHistorialSolicitud, func(solicitudId int) ([]int, error) {
			return []int{}, nil
		})
		defer monkey.Unpatch(clients.ConsultarIdsHistorialSolicitud)
		monkey.Patch(clients.ConsultarSoportesSolicitud, func(id int) ([]models.SoporteSolicitud, error) {
			return []models.SoporteSolicitud{
				{Id: 1, DocumentoId: 101},
				{Id: 2, DocumentoId: 102},
			}, nil
		})
		defer monkey.Unpatch(clients.ConsultarSoportesSolicitud)
		actualizaciones := 0
		monkey.Patch(clients.ActualizarSoporteSolicitud, func(soporteId, solicitudId int, estado string) (*models.SoporteSolicitud, error) {
			actualizaciones++
			if actualizaciones == 1 {
				return nil, errors.New("falla parcial al actualizar soporte")
			}
			return &models.SoporteSolicitud{Id: soporteId}, nil
		})
		defer monkey.Unpatch(clients.ActualizarSoporteSolicitud)
		monkey.Patch(clients.RegistrarHistorialSolicitudEstado, func(solicitudId, terceroId int, justificacion string, estadoId int) (*models.HistorialSolicitud, error) {
			return historialMock, nil
		})
		defer monkey.Unpatch(clients.RegistrarHistorialSolicitudEstado)
		resultado, err := service.CambiarEstado(request)
		if err != nil {
			t.Fatalf("no se esperaba error a pesar de un fallo parcial, llegó: %v", err)
		}
		if resultado.Id != historialMock.Id {
			t.Fatalf("se esperaba ID %d y llegó %d", historialMock.Id, resultado.Id)
		}
		if actualizaciones != 2 {
			t.Fatalf("se esperaban 2 intentos de actualización, hubo %d", actualizaciones)
		}
	})
	t.Run("Ok_HistorialVacio", func(t *testing.T) {
		request := models.SolicitudAprobarRechazarRequest{
			SolicitudId:     10,
			TerceroId:       20,
			EstadoSolicitud: "APROBADA",
			EstadoSoporte:   "",
		}
		historialMock := &models.HistorialSolicitud{Id: 100}
		monkey.Patch(clients.ConsultarEstadoSolicitud, func(codigo string) (*models.EstadoSolicitud, error) {
			return &models.EstadoSolicitud{Id: 5}, nil
		})
		defer monkey.Unpatch(clients.ConsultarEstadoSolicitud)
		monkey.Patch(clients.ConsultarIdsHistorialSolicitud, func(solicitudId int) ([]int, error) {
			return []int{}, nil
		})
		defer monkey.Unpatch(clients.ConsultarIdsHistorialSolicitud)
		desactivaciones := 0
		monkey.Patch(clients.DesactivarHistorialSolicitud, func(idHistorial int) (bool, error) {
			desactivaciones++
			return true, nil
		})
		defer monkey.Unpatch(clients.DesactivarHistorialSolicitud)
		monkey.Patch(clients.RegistrarHistorialSolicitudEstado, func(solicitudId, terceroId int, justificacion string, estadoId int) (*models.HistorialSolicitud, error) {
			return historialMock, nil
		})
		defer monkey.Unpatch(clients.RegistrarHistorialSolicitudEstado)
		resultado, err := service.CambiarEstado(request)
		if err != nil {
			t.Fatalf("no se esperaba error con historial vacío, llegó: %v", err)
		}
		if resultado.Id != historialMock.Id {
			t.Fatalf("se esperaba ID %d y llegó %d", historialMock.Id, resultado.Id)
		}
		if desactivaciones != 0 {
			t.Fatalf("no se esperaban desactivaciones, hubo %d", desactivaciones)
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

	t.Run("Ok_SolicitudSuspensionConSabaticoValido", func(t *testing.T) {
		idSabatico := 50
		req := models.SolicitudRequest{
			TerceroId:       1,
			TipoSolicitudId: "SS",
			SabaticoId:      &idSabatico,
			Formulario:      json.RawMessage("{}"),
		}
		fechaReciente := time.Now().AddDate(0, -1, 0).Format("2006-01-02 15:04:05 -0700 -0700")
		sabaticoMock := &models.Sabatico{Id: idSabatico, FechaCreacion: fechaReciente}
		solicitudMock := &models.Solicitud{Id: 200}

		monkey.Patch(clients.ConsultarTipoSolicitud, func(codigo string) (*models.TipoSolicitud, error) {
			return &models.TipoSolicitud{Id: 2, CodigoAbreviacion: "SS"}, nil
		})
		defer monkey.Unpatch(clients.ConsultarTipoSolicitud)
		monkey.Patch(clients.ConsultarSabatico, func(id int) (*models.Sabatico, error) {
			return sabaticoMock, nil
		})
		defer monkey.Unpatch(clients.ConsultarSabatico)
		monkey.Patch(clients.RegistrarSolicitud, func(terceroId int, tipoSolicitudId int, sabaticoId *int) (*models.Solicitud, error) {
			return solicitudMock, nil
		})
		defer monkey.Unpatch(clients.RegistrarSolicitud)
		formularioCount := 0
		monkey.Patch(clients.RegistrarFormularioSolicitud, func(solicitudId int, contenido string) (*models.FormularioSolicitud, error) {
			formularioCount++
			return &models.FormularioSolicitud{}, nil
		})
		defer monkey.Unpatch(clients.RegistrarFormularioSolicitud)
		historialEstados := make([]string, 0, 1)
		monkey.Patch(clients.RegistrarHistorialSolicitud, func(solicitudId int, terceroId int, justificacion string, codigoEstadoSolicitud string) (*models.HistorialSolicitud, error) {
			historialEstados = append(historialEstados, codigoEstadoSolicitud)
			return &models.HistorialSolicitud{}, nil
		})
		defer monkey.Unpatch(clients.RegistrarHistorialSolicitud)

		resultado, err := service.CrearSolicitud(req)
		if err != nil {
			t.Fatalf("no se esperaba error, llegó: %v", err)
		}
		if resultado.Id != solicitudMock.Id {
			t.Errorf("se esperaba ID %d, llegó %d", solicitudMock.Id, resultado.Id)
		}
		if formularioCount != 1 {
			t.Errorf("se esperaba creación de formulario para SUSPENSION, hubo %d", formularioCount)
		}
		if len(historialEstados) != 1 || historialEstados[0] != string(enums.RADICADA_ENVIADA_SA) {
			t.Errorf("se esperaba historial con estado %s, llegó %v", enums.RADICADA_ENVIADA_SA, historialEstados)
		}
	})

	t.Run("Error_SuspensionSinSabatico", func(t *testing.T) {
		req := models.SolicitudRequest{
			TipoSolicitudId: "SS",
		}
		monkey.Patch(clients.ConsultarTipoSolicitud, func(codigo string) (*models.TipoSolicitud, error) {
			return &models.TipoSolicitud{CodigoAbreviacion: "SS"}, nil
		})
		defer monkey.Unpatch(clients.ConsultarTipoSolicitud)

		_, err := service.CrearSolicitud(req)
		if err == nil || err.Error() != "a SUSPENSION request must have an associated Sabbatical" {
			t.Fatalf("se esperaba error de validación de sabático para SUSPENSION, llegó: %v", err)
		}
	})

	t.Run("Error_SuspensionFechaInvalida", func(t *testing.T) {
		idSabatico := 50
		req := models.SolicitudRequest{
			TipoSolicitudId: "SS",
			SabaticoId:      &idSabatico,
		}
		monkey.Patch(clients.ConsultarTipoSolicitud, func(codigo string) (*models.TipoSolicitud, error) {
			return &models.TipoSolicitud{CodigoAbreviacion: "SS"}, nil
		})
		defer monkey.Unpatch(clients.ConsultarTipoSolicitud)
		monkey.Patch(clients.ConsultarSabatico, func(id int) (*models.Sabatico, error) {
			return &models.Sabatico{FechaCreacion: "fecha-mal-formada"}, nil
		})
		defer monkey.Unpatch(clients.ConsultarSabatico)

		_, err := service.CrearSolicitud(req)
		if err == nil || err.Error() != "invalid FechaCreacion format for the Sabbatical" {
			t.Fatalf("se esperaba error por formato de fecha inválido, llegó: %v", err)
		}
	})

	t.Run("Error_SuspensionFueraDeVentana", func(t *testing.T) {
		idSabatico := 50
		req := models.SolicitudRequest{
			TipoSolicitudId: "SS",
			SabaticoId:      &idSabatico,
		}
		fechaVieja := time.Now().AddDate(0, -6, 0).Format("2006-01-02 15:04:05 -0700 -0700")
		monkey.Patch(clients.ConsultarTipoSolicitud, func(codigo string) (*models.TipoSolicitud, error) {
			return &models.TipoSolicitud{CodigoAbreviacion: "SS"}, nil
		})
		defer monkey.Unpatch(clients.ConsultarTipoSolicitud)
		monkey.Patch(clients.ConsultarSabatico, func(id int) (*models.Sabatico, error) {
			return &models.Sabatico{FechaCreacion: fechaVieja}, nil
		})
		defer monkey.Unpatch(clients.ConsultarSabatico)

		_, err := service.CrearSolicitud(req)
		if err == nil || err.Error() != "a SUSPENSION request cannot be created after 3 months from the Sabbatical creation date" {
			t.Fatalf("se esperaba error por ventana de 3 meses excedida, llegó: %v", err)
		}
	})

	t.Run("Ok_SolicitudModificacionConSabaticoValido", func(t *testing.T) {
		idSabatico := 70
		req := models.SolicitudRequest{
			TerceroId:       1,
			TipoSolicitudId: "MS",
			SabaticoId:      &idSabatico,
			Formulario:      json.RawMessage("{}"),
		}
		fechaReciente := time.Now().AddDate(0, -1, 0).Format("2006-01-02 15:04:05 -0700 -0700")
		sabaticoMock := &models.Sabatico{Id: idSabatico, FechaCreacion: fechaReciente}
		solicitudMock := &models.Solicitud{Id: 300}

		monkey.Patch(clients.ConsultarTipoSolicitud, func(codigo string) (*models.TipoSolicitud, error) {
			return &models.TipoSolicitud{Id: 3, CodigoAbreviacion: "MS"}, nil
		})
		defer monkey.Unpatch(clients.ConsultarTipoSolicitud)
		monkey.Patch(clients.ConsultarSabatico, func(id int) (*models.Sabatico, error) {
			return sabaticoMock, nil
		})
		defer monkey.Unpatch(clients.ConsultarSabatico)
		monkey.Patch(clients.RegistrarSolicitud, func(terceroId int, tipoSolicitudId int, sabaticoId *int) (*models.Solicitud, error) {
			return solicitudMock, nil
		})
		defer monkey.Unpatch(clients.RegistrarSolicitud)
		formularioCount := 0
		monkey.Patch(clients.RegistrarFormularioSolicitud, func(solicitudId int, contenido string) (*models.FormularioSolicitud, error) {
			formularioCount++
			return &models.FormularioSolicitud{}, nil
		})
		defer monkey.Unpatch(clients.RegistrarFormularioSolicitud)
		historialEstados := make([]string, 0, 1)
		monkey.Patch(clients.RegistrarHistorialSolicitud, func(solicitudId int, terceroId int, justificacion string, codigoEstadoSolicitud string) (*models.HistorialSolicitud, error) {
			historialEstados = append(historialEstados, codigoEstadoSolicitud)
			return &models.HistorialSolicitud{}, nil
		})
		defer monkey.Unpatch(clients.RegistrarHistorialSolicitud)

		resultado, err := service.CrearSolicitud(req)
		if err != nil {
			t.Fatalf("no se esperaba error, llegó: %v", err)
		}
		if resultado.Id != solicitudMock.Id {
			t.Errorf("se esperaba ID %d, llegó %d", solicitudMock.Id, resultado.Id)
		}
		if formularioCount != 1 {
			t.Errorf("se esperaba creación de formulario para MODIFICACION, hubo %d", formularioCount)
		}
		if len(historialEstados) != 1 || historialEstados[0] != string(enums.RADICADA_ENVIADA_SA) {
			t.Errorf("se esperaba historial con estado %s, llegó %v", enums.RADICADA_ENVIADA_SA, historialEstados)
		}
	})

	t.Run("Error_ModificacionSinSabatico", func(t *testing.T) {
		req := models.SolicitudRequest{
			TipoSolicitudId: "MS",
		}
		monkey.Patch(clients.ConsultarTipoSolicitud, func(codigo string) (*models.TipoSolicitud, error) {
			return &models.TipoSolicitud{CodigoAbreviacion: "MS"}, nil
		})
		defer monkey.Unpatch(clients.ConsultarTipoSolicitud)

		_, err := service.CrearSolicitud(req)
		if err == nil || err.Error() != "a MODIFICATION request must have an associated Sabbatical" {
			t.Fatalf("se esperaba error de validación de sabático para MODIFICACION, llegó: %v", err)
		}
	})

	t.Run("Error_ModificacionFechaInvalida", func(t *testing.T) {
		idSabatico := 70
		req := models.SolicitudRequest{
			TipoSolicitudId: "MS",
			SabaticoId:      &idSabatico,
		}
		monkey.Patch(clients.ConsultarTipoSolicitud, func(codigo string) (*models.TipoSolicitud, error) {
			return &models.TipoSolicitud{CodigoAbreviacion: "MS"}, nil
		})
		defer monkey.Unpatch(clients.ConsultarTipoSolicitud)
		monkey.Patch(clients.ConsultarSabatico, func(id int) (*models.Sabatico, error) {
			return &models.Sabatico{FechaCreacion: "fecha-mal-formada"}, nil
		})
		defer monkey.Unpatch(clients.ConsultarSabatico)

		_, err := service.CrearSolicitud(req)
		if err == nil || err.Error() != "invalid FechaCreacion format for the Sabbatical" {
			t.Fatalf("se esperaba error por formato de fecha inválido, llegó: %v", err)
		}
	})

	t.Run("Error_ModificacionFueraDeVentana", func(t *testing.T) {
		idSabatico := 70
		req := models.SolicitudRequest{
			TipoSolicitudId: "MS",
			SabaticoId:      &idSabatico,
		}
		fechaVieja := time.Now().AddDate(0, -6, 0).Format("2006-01-02 15:04:05 -0700 -0700")
		monkey.Patch(clients.ConsultarTipoSolicitud, func(codigo string) (*models.TipoSolicitud, error) {
			return &models.TipoSolicitud{CodigoAbreviacion: "MS"}, nil
		})
		defer monkey.Unpatch(clients.ConsultarTipoSolicitud)
		monkey.Patch(clients.ConsultarSabatico, func(id int) (*models.Sabatico, error) {
			return &models.Sabatico{FechaCreacion: fechaVieja}, nil
		})
		defer monkey.Unpatch(clients.ConsultarSabatico)

		_, err := service.CrearSolicitud(req)
		if err == nil || err.Error() != "a MODIFICATION request cannot be created after 3 months from the Sabbatical creation date" {
			t.Fatalf("se esperaba error por ventana de 3 meses excedida, llegó: %v", err)
		}
	})

	t.Run("Error_ConsultarTipoSolicitudFalla", func(t *testing.T) {
		req := models.SolicitudRequest{TipoSolicitudId: "DESCONOCIDO"}
		monkey.Patch(clients.ConsultarTipoSolicitud, func(codigo string) (*models.TipoSolicitud, error) {
			return nil, errors.New("tipo solicitud no encontrado")
		})
		defer monkey.Unpatch(clients.ConsultarTipoSolicitud)

		_, err := service.CrearSolicitud(req)
		if err == nil {
			t.Fatal("se esperaba error al fallar la consulta de tipo de solicitud")
		}
	})

	t.Run("Error_RegistrarSolicitudFalla", func(t *testing.T) {
		req := models.SolicitudRequest{
			TerceroId:       1,
			TipoSolicitudId: "NS",
		}
		monkey.Patch(clients.ConsultarTipoSolicitud, func(codigo string) (*models.TipoSolicitud, error) {
			return &models.TipoSolicitud{Id: 1, CodigoAbreviacion: "NS"}, nil
		})
		defer monkey.Unpatch(clients.ConsultarTipoSolicitud)
		monkey.Patch(clients.RegistrarSolicitud, func(terceroId int, tipoSolicitudId int, sabaticoId *int) (*models.Solicitud, error) {
			return nil, errors.New("error registrando solicitud")
		})
		defer monkey.Unpatch(clients.RegistrarSolicitud)

		_, err := service.CrearSolicitud(req)
		if err == nil {
			t.Fatal("se esperaba error al fallar el registro de la solicitud")
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
				SolicitudId: models.SolicitudMinima{Id: 10},
			},
			{
				Id:          2,
				Contenido:   `{"docente":{"facultad":"OTRA_FACULTAD"}}`,
				SolicitudId: models.SolicitudMinima{Id: 11},
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

	t.Run("Error_DependenciaVacia", func(t *testing.T) {
		monkey.Patch(clients.ConsultarSecretariaAcademicaDocumentoUserId, func(id string) (*models.Persona, error) {
			return &models.Persona{Dependencia: "   "}, nil
		})
		defer monkey.Unpatch(clients.ConsultarSecretariaAcademicaDocumentoUserId)
		monkey.Patch(clients.ConsultarTodosFormulariosSolicitud, func() ([]models.FormularioSolicitud, error) {
			return []models.FormularioSolicitud{}, nil
		})
		defer monkey.Unpatch(clients.ConsultarTodosFormulariosSolicitud)

		_, err := service.GetFormulariosByDocumentoId("123", nil)
		if err == nil || err.Error() != "dependencia de secretaria academica vacía" {
			t.Fatalf("se esperaba error por dependencia vacía, llegó: %v", err)
		}
	})

	t.Run("Ok_SinFormulariosCoincidentes", func(t *testing.T) {
		monkey.Patch(clients.ConsultarSecretariaAcademicaDocumentoUserId, func(id string) (*models.Persona, error) {
			return &models.Persona{Dependencia: "FACULTAD_INGENIERIA"}, nil
		})
		defer monkey.Unpatch(clients.ConsultarSecretariaAcademicaDocumentoUserId)
		monkey.Patch(clients.ConsultarTodosFormulariosSolicitud, func() ([]models.FormularioSolicitud, error) {
			return []models.FormularioSolicitud{
				{Id: 1, Contenido: `{"docente":{"facultad":"OTRA"}}`, SolicitudId: models.SolicitudMinima{Id: 10}},
			}, nil
		})
		defer monkey.Unpatch(clients.ConsultarTodosFormulariosSolicitud)

		resultado, err := service.GetFormulariosByDocumentoId("123", nil)
		if err != nil {
			t.Fatalf("no se esperaba error, llegó: %v", err)
		}
		if len(resultado) != 0 {
			t.Errorf("se esperaba slice vacío, llegaron %d", len(resultado))
		}
	})

	t.Run("Ok_SolicitudesDuplicadasFiltradas", func(t *testing.T) {
		monkey.Patch(clients.ConsultarSecretariaAcademicaDocumentoUserId, func(id string) (*models.Persona, error) {
			return &models.Persona{Dependencia: "INGENIERIA"}, nil
		})
		defer monkey.Unpatch(clients.ConsultarSecretariaAcademicaDocumentoUserId)
		monkey.Patch(clients.ConsultarTodosFormulariosSolicitud, func() ([]models.FormularioSolicitud, error) {
			return []models.FormularioSolicitud{
				{Id: 1, Contenido: `{"docente":{"facultad":"INGENIERIA"}}`, SolicitudId: models.SolicitudMinima{Id: 10}},
				{Id: 2, Contenido: `{"docente":{"facultad":"INGENIERIA"}}`, SolicitudId: models.SolicitudMinima{Id: 10}},
			}, nil
		})
		defer monkey.Unpatch(clients.ConsultarTodosFormulariosSolicitud)
		consultas := 0
		monkey.Patch(clients.ConsultarIdsHistorialSolicitud, func(id int) ([]int, error) {
			consultas++
			return []int{100}, nil
		})
		defer monkey.Unpatch(clients.ConsultarIdsHistorialSolicitud)
		monkey.Patch(clients.ConsultarHistorialSolicitudIdEstadoId, func(id int, est []string) ([]models.HistorialSolicitud, error) {
			return []models.HistorialSolicitud{{Id: 100, SolicitudId: models.IdReference{Id: 10}}}, nil
		})
		defer monkey.Unpatch(clients.ConsultarHistorialSolicitudIdEstadoId)

		_, err := service.GetFormulariosByDocumentoId("123", nil)
		if err != nil {
			t.Fatalf("no se esperaba error, llegó: %v", err)
		}
		if consultas != 1 {
			t.Errorf("se esperaba 1 consulta de historial (deduplicación por solicitud), hubo %d", consultas)
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

	t.Run("Ok_SinDocumentos", func(t *testing.T) {
		req := models.RadicarSolicitudRequest{
			SolicitudId:  1,
			FormularioId: 2,
			Formulario:   json.RawMessage("{}"),
			DocumentosId: []int{},
		}
		monkey.Patch(clients.ConsultarSolicitud, func(id int) (*models.Solicitud, error) {
			return &models.Solicitud{Id: 1, TerceroId: 10}, nil
		})
		defer monkey.Unpatch(clients.ConsultarSolicitud)
		monkey.Patch(clients.ConsultarIdsHistorialSolicitud, func(id int) ([]int, error) { return []int{}, nil })
		defer monkey.Unpatch(clients.ConsultarIdsHistorialSolicitud)
		monkey.Patch(clients.RegistrarHistorialSolicitud, func(s, t int, j, c string) (*models.HistorialSolicitud, error) {
			return &models.HistorialSolicitud{Id: 50}, nil
		})
		defer monkey.Unpatch(clients.RegistrarHistorialSolicitud)
		monkey.Patch(clients.ActualizarFormularioSolicitud, func(s, f int, c string) (*models.FormularioSolicitud, error) {
			return &models.FormularioSolicitud{Id: 2}, nil
		})
		defer monkey.Unpatch(clients.ActualizarFormularioSolicitud)
		soporteCalls := 0
		monkey.Patch(clients.ActualizarSoporteSolicitud, func(doc, sol int, est string) (*models.SoporteSolicitud, error) {
			soporteCalls++
			return &models.SoporteSolicitud{Id: doc}, nil
		})
		defer monkey.Unpatch(clients.ActualizarSoporteSolicitud)

		res, err := service.RadicarSolicitud(req)
		if err != nil {
			t.Fatalf("no se esperaba error: %v", err)
		}
		if soporteCalls != 0 {
			t.Errorf("no se esperaban llamadas a ActualizarSoporteSolicitud, hubo %d", soporteCalls)
		}
		if soportes, ok := res["soportes"].([]*models.SoporteSolicitud); ok && len(soportes) != 0 {
			t.Errorf("se esperaba lista de soportes vacía, llegaron %d", len(soportes))
		}
	})

	t.Run("Ok_MultiplesDocumentos", func(t *testing.T) {
		req := models.RadicarSolicitudRequest{
			SolicitudId:  1,
			FormularioId: 2,
			Formulario:   json.RawMessage("{}"),
			DocumentosId: []int{301, 302, 303},
		}
		monkey.Patch(clients.ConsultarSolicitud, func(id int) (*models.Solicitud, error) {
			return &models.Solicitud{Id: 1, TerceroId: 10}, nil
		})
		defer monkey.Unpatch(clients.ConsultarSolicitud)
		monkey.Patch(clients.ConsultarIdsHistorialSolicitud, func(id int) ([]int, error) { return []int{}, nil })
		defer monkey.Unpatch(clients.ConsultarIdsHistorialSolicitud)
		monkey.Patch(clients.RegistrarHistorialSolicitud, func(s, t int, j, c string) (*models.HistorialSolicitud, error) {
			return &models.HistorialSolicitud{Id: 50}, nil
		})
		defer monkey.Unpatch(clients.RegistrarHistorialSolicitud)
		monkey.Patch(clients.ActualizarFormularioSolicitud, func(s, f int, c string) (*models.FormularioSolicitud, error) {
			return &models.FormularioSolicitud{Id: 2}, nil
		})
		defer monkey.Unpatch(clients.ActualizarFormularioSolicitud)
		soporteCalls := 0
		monkey.Patch(clients.ActualizarSoporteSolicitud, func(doc, sol int, est string) (*models.SoporteSolicitud, error) {
			soporteCalls++
			return &models.SoporteSolicitud{Id: doc}, nil
		})
		defer monkey.Unpatch(clients.ActualizarSoporteSolicitud)

		res, err := service.RadicarSolicitud(req)
		if err != nil {
			t.Fatalf("no se esperaba error: %v", err)
		}
		if soporteCalls != 3 {
			t.Errorf("se esperaban 3 actualizaciones de soporte, hubo %d", soporteCalls)
		}
		soportes, ok := res["soportes"].([]*models.SoporteSolicitud)
		if !ok || len(soportes) != 3 {
			t.Errorf("se esperaban 3 soportes en la respuesta, llegaron %d", len(soportes))
		}
	})

	t.Run("Error_ActualizarSoporteFalla", func(t *testing.T) {
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
		monkey.Patch(clients.ConsultarIdsHistorialSolicitud, func(id int) ([]int, error) { return []int{}, nil })
		defer monkey.Unpatch(clients.ConsultarIdsHistorialSolicitud)
		monkey.Patch(clients.RegistrarHistorialSolicitud, func(s, t int, j, c string) (*models.HistorialSolicitud, error) {
			return &models.HistorialSolicitud{Id: 50}, nil
		})
		defer monkey.Unpatch(clients.RegistrarHistorialSolicitud)
		monkey.Patch(clients.ActualizarFormularioSolicitud, func(s, f int, c string) (*models.FormularioSolicitud, error) {
			return &models.FormularioSolicitud{Id: 2}, nil
		})
		defer monkey.Unpatch(clients.ActualizarFormularioSolicitud)
		monkey.Patch(clients.ActualizarSoporteSolicitud, func(doc, sol int, est string) (*models.SoporteSolicitud, error) {
			return nil, errors.New("falla actualizando soporte")
		})
		defer monkey.Unpatch(clients.ActualizarSoporteSolicitud)

		_, err := service.RadicarSolicitud(req)
		if err == nil {
			t.Fatal("se esperaba error al fallar la actualización del soporte; RadicarSolicitud aborta a diferencia de CambiarEstado")
		}
	})
}
