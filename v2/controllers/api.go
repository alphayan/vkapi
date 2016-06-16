package v2


import "github.com/astaxie/beego"

type VtwoController struct {
beego.Controller
}
//@route /vtwo/api [get]
func (c *VtwoController)VoneApi(){
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin","*")
}