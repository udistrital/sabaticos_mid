package controllers

import (
	"api_mid_sabaticos/helpers"
	"api_mid_sabaticos/models"
	"api_mid_sabaticos/service"
	"net/http"

	"github.com/astaxie/beego"
)

// SoportesolicitudController operations for Soportesolicitud
type SoporteSolicitudController struct {
	beego.Controller
}

// URLMapping ...
func (c *SoporteSolicitudController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Soportesolicitud con múltiples documentos
// @router / [post]
func (c *SoporteSolicitudController) Post() {
	var soporteSolicitudRequest models.SoporteSolicitudRequest

	// Obtener datos de form-data
	soporteSolicitudRequest.TerceroId, _ = c.GetInt("tercero_id")
	soporteSolicitudRequest.SolicitudId, _ = c.GetInt("solicitud_id")
	soporteSolicitudRequest.RolUsuario = c.GetString("rol_usuario")
	soporteSolicitudRequest.EstadoSoporteSolicitud = c.GetString("estado_soporte_solicitud")

	// Validar campos requeridos
	if soporteSolicitudRequest.TerceroId == 0 || soporteSolicitudRequest.SolicitudId == 0 ||
		soporteSolicitudRequest.RolUsuario == "" || soporteSolicitudRequest.EstadoSoporteSolicitud == "" {
		helpers.JSONResponse(&c.Controller, false, http.StatusBadRequest, nil, "los campos terceroId, solicitudId, estadoSoporteSolicitud y rolUsuario son requeridos")
		return
	}

	// Obtener archivos
	files, _ := c.GetFiles("documentos")

	if len(files) == 0 {
		helpers.JSONResponse(&c.Controller, false, http.StatusBadRequest, nil, "se debe enviar al menos un archivo")
		return
	}

	// Llamar al servicio para procesar la solicitud
	respuesta, err := service.CrearSoporteSolicitud(soporteSolicitudRequest, files)

	if err != nil {
		helpers.JSONResponse(&c.Controller, false, http.StatusBadRequest, nil, err.Error())
		return
	}

	helpers.JSONResponse(&c.Controller, true, http.StatusOK, respuesta, "")
}

// GetOne ...
// @Title GetOne
// @Description get Soportesolicitud by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Soportesolicitud
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SoporteSolicitudController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Soportesolicitud
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Soportesolicitud
// @Failure 403
// @router / [get]
func (c *SoporteSolicitudController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Soportesolicitud
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Soportesolicitud	true		"body for Soportesolicitud content"
// @Success 200 {object} models.Soportesolicitud
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SoporteSolicitudController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Soportesolicitud
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SoporteSolicitudController) Delete() {

}
