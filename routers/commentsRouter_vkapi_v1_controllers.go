package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["vkapi/v1/controllers:VoneController"] = append(beego.GlobalControllerRouter["vkapi/v1/controllers:VoneController"],
		beego.ControllerComments{
			"VoneApiSetonsale",
			`/vone/api/setonsale`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["vkapi/v1/controllers:VoneController"] = append(beego.GlobalControllerRouter["vkapi/v1/controllers:VoneController"],
		beego.ControllerComments{
			"VoneApiInsertMongodb",
			`/vone/api/insertmongodb`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["vkapi/v1/controllers:VoneController"] = append(beego.GlobalControllerRouter["vkapi/v1/controllers:VoneController"],
		beego.ControllerComments{
			"VoneApi",
			`/vone/api`,
			[]string{"post"},
			nil})

}
