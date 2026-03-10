package controllers

import (
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
	c.Mapping("Post", c.Post)
	c.Mapping("Radicar", c.Radicar)
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

// Radicar ...
// @Title Radicar
// @Description Radicar una solicitud (cambiar estado y crear registros)
// @Param	id		path 	int	true		"The id of the Solicitud to radicar"
// @Param	body		body 	models.SolicitudRequest	false	"body for additional data if needed"
// @Success 200 {object} models.SolicitudResponse
// @Failure 400 Bad request
// @Failure 404 Solicitud not found
// @router /radicar/:id [post]
func (c *SolicitudController) Radicar() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.GetString(":id")
	var RadicarSolicitudRequest models.RadicarSolicitudRequest

	requestmanager.FillRequestWithPanic(&c.Controller, &RadicarSolicitudRequest)

	// Validar que el ID de la solicitud esté presente
	if id == "" {
		helpers.JSONResponse(&c.Controller, false, http.StatusBadRequest, nil, "El id de la solicitud es requerido")
		return
	}

	//Validar RadicarSolicitudRequest si es necesario, por ejemplo:
	if RadicarSolicitudRequest.SolicitudId == 0 || RadicarSolicitudRequest.FormularioId == 0 || RadicarSolicitudRequest.Formulario == nil || RadicarSolicitudRequest.DocumentosId == nil || len(RadicarSolicitudRequest.DocumentosId) == 0 {
		helpers.JSONResponse(&c.Controller, false, http.StatusBadRequest, nil, "Los campos SolicitudId, FormularioId, Formulario y DocumentosId son requeridos en el cuerpo de la solicitud")
		return
	}

	// Llamar al servicio para radicar
	solicitud, err := service.RadicarSolicitud(RadicarSolicitudRequest)

	if err != nil {
		helpers.JSONResponse(&c.Controller, false, http.StatusNotFound, nil, "Error al radicar solicitud: "+err.Error())
		return
	}

	respuesta := models.SolicitudResponse{
		Solicitud: solicitud,
	}

	helpers.JSONResponse(&c.Controller, true, http.StatusOK, respuesta, "Solicitud radicada exitosamente")
}
