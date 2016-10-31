package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["vkapi/v1/controllers:VoneController"] = append(beego.GlobalControllerRouter["vkapi/v1/controllers:VoneController"],
		beego.ControllerComments{
			Method: "VoneApiSetonsale",
			Router: `/vone/api/setonsale`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["vkapi/v1/controllers:VoneController"] = append(beego.GlobalControllerRouter["vkapi/v1/controllers:VoneController"],
		beego.ControllerComments{
			Method: "VoneApiInsertMongodb",
			Router: `/vone/api/insertmongodb`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["vkapi/v1/controllers:VoneController"] = append(beego.GlobalControllerRouter["vkapi/v1/controllers:VoneController"],
		beego.ControllerComments{
			Method: "VoneApi",
			Router: `/vone/api`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
