package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["vkapi/v2/controllers:VtwoController"] = append(beego.GlobalControllerRouter["vkapi/v2/controllers:VtwoController"],
		beego.ControllerComments{
			Method: "VtwoApiSetonsale",
			Router: `/vtwo/api/setonsale`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["vkapi/v2/controllers:VtwoController"] = append(beego.GlobalControllerRouter["vkapi/v2/controllers:VtwoController"],
		beego.ControllerComments{
			Method: "VtwoApi",
			Router: `/vtwo/api`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
