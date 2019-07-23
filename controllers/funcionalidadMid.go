package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/api_beego_request/models"
	"github.com/udistrital/utils_oas/request"
)

// Operations about object
type FuncionalidadMidController struct {
	beego.Controller
}

// URLMapping ...
func (c *FuncionalidadMidController) URLMapping() {
	c.Mapping("GetUser", c.GetUser)

}

// GetUser ...
// @Title Get User
// @Description get Usuario by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Usuario
// @Failure 404 not found resource
// @router /:id [get]
func (c *FuncionalidadMidController) GetUser() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetUsuarioById(id) // funcion
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// get Información de Usuario del sistema Agora
func getUserAgora(idUser int64) (dataUser interface{}, outputError map[string]interface{}) {
	if idUser != 0 {
		if err := request.GetJson("http://localhost:8080/v1/usuario", &dataUser); err == nil {

		}
		outputError = map[string]interface{}{"Code": "E_0458", "Body": "Not enough parameter in pagoSsPorPersona", "Type": "error"}
		return outputError, nil
	} else {
		outputError = map[string]interface{}{"Code": "E_0458", "Body": "Not enough parameter in pagoSsPorPersona", "Type": "error"}
		return nil, outputError
	}
}
