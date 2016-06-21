package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["vkapi/v2/controllers:VtwoController"] = append(beego.GlobalControllerRouter["vkapi/v2/controllers:VtwoController"],
		beego.ControllerComments{
			"VtwoApiSetonsale",
			`/vtwo/api/setonsale`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["vkapi/v2/controllers:VtwoController"] = append(beego.GlobalControllerRouter["vkapi/v2/controllers:VtwoController"],
		beego.ControllerComments{
			"VtwoApi",
			`/vtwo/api`,
			[]string{"post"},
			nil})

}
