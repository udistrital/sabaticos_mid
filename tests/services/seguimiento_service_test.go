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
