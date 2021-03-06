package v2

import (
	"github.com/astaxie/beego"
	"vkapi/v2/controllers/setonsale"

	"fmt"
	"vkapi/models/result"
)

type VtwoController struct {
	beego.Controller
}
//@router /vtwo/api/setonsale [get]
func (c *VtwoController)VtwoApiSetonsale() {
	c.Data["Website"] = "www.vkaifu.com"
	c.Data["Email"] = "vk@ivknet.com"
	c.TplName = "v2/setonsale.html"
}

//@router /vtwo/api [post]
func (c *VtwoController)VtwoApi() {
	fmt.Println("收到请求:" + c.Ctx.Request.RequestURI)
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	service := c.GetString("service")
	if service == "vtwo.setonsale" {
		fmt.Println("开始配置优惠")
		result := setonsale.Setonsale(c.GetString("cashiername"), c.GetString("max"), c.GetString("mch_id"), c.GetString("min"), c.GetString("money"), c.GetString("number"), c.GetString("timestart"), c.GetString("timeend"), c.GetString("limamount"))
		c.Data["json"] = &result
		c.ServeJSON()
	} else {
		result := result.Result{}
		result.Result_code = "FAIL"
		result.Result_des = "未知的服务名"
		c.Data["json"] = &result
		c.ServeJSON()
	}
}