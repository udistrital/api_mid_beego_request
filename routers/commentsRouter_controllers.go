package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/api_mid_beego_request/controllers:FuncionalidadMidController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_mid_beego_request/controllers:FuncionalidadMidController"],
        beego.ControllerComments{
            Method: "GetUser",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
