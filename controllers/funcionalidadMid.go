package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
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
	v, err := getUserAgora(id) // funcion getUserAgora
	logs.Info("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF")
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
func getUserAgora(idUser int) (dataUser interface{}, outputError map[string]interface{}) {
	if idUser != 0 { // (1) error parametro
		if response, err := request.GetJsonTest("http://localhost:8080/v1/usuario/"+strconv.Itoa(idUser), &dataUser); err == nil { // (2) error servicio caido
			if response.StatusCode == 200 { // (3) error estado de la solicitud
				return dataUser, nil
			} else {
				logs.Info("Error (3) estado de la solicitud")
				outputError = map[string]interface{}{"Function": "FuncionalidadMidController:getUserAgora", "Error": response.Status}
				return nil, outputError
			}
		} else {
			logs.Info("Error (2) servicio caido")
			outputError = map[string]interface{}{"Function": "FuncionalidadMidController:getUserAgora", "Error": err}
			return nil, outputError
		}
	} else {
		logs.Info("Error (1) Parametro")
		outputError = map[string]interface{}{"Function": "FuncionalidadMidController:getUserAgora", "Error": "null parameter"}
		return nil, outputError
	}
}
