package routers

import (

	"github.com/astaxie/beego"
	"vkapi/controllers"
	"vkapi/v1/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Include(&v1.VoneController{})
}
