package controllers

import (
	"github.com/astaxie/beego"
)

// SabaticoController operations for Sabatico
// @Tag Sabatico
type SabaticoController struct {
	beego.Controller
}

// URLMapping ...
func (c *SabaticoController) URLMapping() {
	//Se desarrollara mas adelante, por ahora solo se consulta el sabatico desde el servicio de sabaticos
}
