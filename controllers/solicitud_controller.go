package controllers

import (
	"fmt"
	"net/http"

	"github.com/udistrital/sabaticos_mid/helpers"
	"github.com/udistrital/sabaticos_mid/models"
	"github.com/udistrital/sabaticos_mid/service"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/errorhandler"
	"github.com/udistrital/utils_oas/requestmanager"
)

// SolicitudController operations for Solicitud
// @Tag Solicitud
type SolicitudController struct {
	beego.Controller
}

// URLMapping ...
func (c *SolicitudController) URLMapping() {
	c.Mapping("aprobar", c.Aprobar)
	c.Mapping("rechazar", c.Rechazar)
}

// Post ...
// @Title Create
// @Description create Solicitud
// @Param	body		body 	models.Solicitud	true		"body for Solicitud content"
// @Success 201 {object} models.Solicitud
// @Failure 403 body is empty
// @router / [post]
func (c *SolicitudController) Post() {
	defer errorhandler.HandlePanic(&c.Controller)

	var solicitudRequest models.SolicitudRequest

	requestmanager.FillRequestWithPanic(&c.Controller, &solicitudRequest)

	if solicitudRequest.TerceroId <= 0 || solicitudRequest.TipoSolicitudId == "" {
		helpers.JSONResponse(&c.Controller, false, http.StatusBadRequest, nil, "Los campos terceroId y tipoSolicitudId son requeridos")
	}

	solicitud, err := service.CrearSolicitud(solicitudRequest)

	if err != nil {
		helpers.JSONResponse(&c.Controller, false, http.StatusNotFound, nil, "Recurso no encontrado: "+err.Error())
		return
	}

	respuesta := models.SolicitudResponse{
		Solicitud: solicitud,
	}

	helpers.JSONResponse(&c.Controller, true, http.StatusCreated, respuesta, "Solicitud creada exitosamente")
}

// Aprobar ...
// @Title Aprobar
// @Description Aprueba una solicitud creando un registro en el historial
// @Param   body  body  interface{}  true  "body para aprobar solicitud"
// @Success 200 {object} interface{}
// @Failure 400 the request contains incorrect syntax
// @router /aprobar [post]
func (c *SolicitudController) Aprobar() {
	defer errorhandler.HandlePanic(&c.Controller)
	fmt.Println("ENTRA A AL AFUNCION DE APROBAR")
	var SolicitudAprobarRechazarRequest models.SolicitudAprobarRechazarRequest

	requestmanager.FillRequestWithPanic(&c.Controller, &SolicitudAprobarRechazarRequest)

	if SolicitudAprobarRechazarRequest.TerceroId <= 0 || SolicitudAprobarRechazarRequest.SolicitudId <= 0 {
		helpers.JSONResponse(&c.Controller, false, http.StatusBadRequest, nil, "Los campos terceroId, solicitudId")
	}

	HistorialSolicitud, err := service.Aprobar(SolicitudAprobarRechazarRequest)

	if err != nil {
		helpers.JSONResponse(&c.Controller, false, http.StatusNotFound, nil, "Recurso no encontrado: "+err.Error())
		return
	}

	respuesta := models.SolicitudAprobarRechazarResponse{
		SolicitudId:     SolicitudAprobarRechazarRequest.SolicitudId,
		EstadoSolicitud: 3,
	}

	fmt.Printf("HistorialSolicitud: %+v\n", HistorialSolicitud)

	helpers.JSONResponse(&c.Controller, true, http.StatusOK, respuesta, "Solicitud procesada exitosamente")
}

func (c *SolicitudController) Rechazar() {
	fmt.Println("Rechazar solicitud - Endpoint en desarrollo")
}
