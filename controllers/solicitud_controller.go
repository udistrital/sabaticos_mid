package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
	"github.com/udistrital/sabaticos_mid/clients"
	"github.com/udistrital/sabaticos_mid/helpers"
	"github.com/udistrital/sabaticos_mid/models"
	"github.com/udistrital/sabaticos_mid/service"
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
	c.Mapping("aprobar-rechazar", c.Aprobar_Rechazar_solicitud)
	c.Mapping("Post", c.Post)
	c.Mapping("GetFormulariosByDocumentoId", c.GetFormulariosByDocumentoId)
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
		helpers.JSONResponse(&c.Controller, false, http.StatusBadRequest, nil, "fields terceroId and tipoSolicitudId are required")
	}

	solicitud, err := service.CrearSolicitud(solicitudRequest)

	if err != nil {
		helpers.JSONResponse(&c.Controller, false, http.StatusNotFound, nil, "resource not found: "+err.Error())
		return
	}

	respuesta := models.SolicitudResponse{
		Solicitud: solicitud,
	}

	helpers.JSONResponse(&c.Controller, true, http.StatusCreated, respuesta, "request created successfully")
}

// Aprobar_Rechazar ...
// @Title Aprobar_Rechazar
// @Description Aprueba o rechaza una solicitud creando un registro en el historial
// @Param   body  body  interface{}  true  "body para aprobar o rechazar solicitud"
// @Success 200 {object} interface{}
// @Failure 400 the request contains incorrect syntax
// @router /aprobar-rechazar [post]
func (c *SolicitudController) Aprobar_Rechazar_solicitud() {
	defer errorhandler.HandlePanic(&c.Controller)

	var ValidarRequest models.SolicitudAprobarRechazarRequest
	requestmanager.FillRequestWithPanic(&c.Controller, &ValidarRequest)

	if ValidarRequest.TerceroId <= 0 {
		helpers.JSONResponse(
			&c.Controller,
			false,
			http.StatusBadRequest,
			nil,
			"El campo terceroId es necesario",
		)
		return
	}

	estado_actualizar, err_estado := clients.ConsultarEstadoSolicitud(ValidarRequest.EstadoSolicitud)
	if err_estado != nil {
		helpers.JSONResponse(
			&c.Controller,
			false,
			http.StatusInternalServerError,
			nil,
			"Error consultando estado de solicitud: "+err_estado.Error(),
		)
		return
	}

	// 2. Construir respuesta
	respuesta := models.SolicitudAprobarRechazarResponse{
		SolicitudId:     ValidarRequest.SolicitudId,
		EstadoSolicitud: estado_actualizar.Id,
	}

	helpers.JSONResponse(
		&c.Controller,
		true,
		http.StatusOK,
		respuesta,
		"Solicitud procesada exitosamente",
	)
}

// GetFormulariosByDocumentoId ...
// @Title GetFormulariosByDocumentoId
// @Description Obtener formularios por documentoId
// @Param	documentoId	path 	string	true		"Identificador del usuario en front"
// @Param	estadoSolicitud	query 	[]string	false		"Estados de la solicitud (repetible)"
// @Success 200 {object} interface{}
// @Failure 400 bad request
// @router /formularios/:documentoId [get]
func (c *SolicitudController) GetFormulariosByDocumentoId() {
	defer errorhandler.HandlePanic(&c.Controller)

	documentoId := c.GetString(":documentoId")
	estadosSolicitud := c.Ctx.Request.URL.Query()["estadoSolicitud"]
	if documentoId == "" {
		helpers.JSONResponse(&c.Controller, false, http.StatusBadRequest, nil, "documentoId is required")
		return
	}

	formularios, err := service.GetFormulariosByDocumentoId(documentoId, estadosSolicitud)
	if err != nil {
		helpers.JSONResponse(&c.Controller, false, http.StatusNotFound, nil, "error getting formularios: "+err.Error())
		return
	}

	helpers.JSONResponse(&c.Controller, true, http.StatusOK, formularios, "request completed successfully")
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
		helpers.JSONResponse(&c.Controller, false, http.StatusBadRequest, nil, "request id is required")
		return
	}

	//Validar RadicarSolicitudRequest si es necesario, por ejemplo:
	if RadicarSolicitudRequest.SolicitudId == 0 || RadicarSolicitudRequest.FormularioId == 0 || RadicarSolicitudRequest.Formulario == nil || RadicarSolicitudRequest.DocumentosId == nil || len(RadicarSolicitudRequest.DocumentosId) == 0 {
		helpers.JSONResponse(&c.Controller, false, http.StatusBadRequest, nil, "fields SolicitudId, FormularioId, Formulario and DocumentosId are required in the request body")
		return
	}

	// Llamar al servicio para radicar
	solicitud, err := service.RadicarSolicitud(RadicarSolicitudRequest)

	if err != nil {
		helpers.JSONResponse(&c.Controller, false, http.StatusNotFound, nil, "error filing request: "+err.Error())
		return
	}

	respuesta := models.SolicitudResponse{
		Solicitud: solicitud,
	}

	helpers.JSONResponse(&c.Controller, true, http.StatusOK, respuesta, "request filed successfully")
}
