package routers

import (

	"github.com/astaxie/beego"
	"vkapi/controllers"
	"vkapi/v1/controllers"
	"vkapi/v2/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Include(&v1.VoneController{})
	beego.Include(&v2.VtwoController{})
}
