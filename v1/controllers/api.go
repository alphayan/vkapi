package v1

import (
	"github.com/astaxie/beego"
	"vkapi/v1/controllers/setonsale"
	"fmt"
	"vkapi/models/result"
	"vkapi/v1/controllers/insertmongodb"
	"vkapi/v1/controllers/insertpassword"
	"vkapi/v1/controllers/tinyurl"
)

type VoneController struct {
	beego.Controller
}
//@router /vone/api/setonsale [get]
func (c *VoneController)VoneApiSetonsale() {
	c.Data["Website"] = "www.vkaifu.com"
	c.Data["Email"] = "vk@ivknet.com"
	c.TplName = "v1/setonsale.html"
}
//@router /vone/api/insertmongodb [get]
func (c *VoneController)VoneApiInsertMongodb() {
	c.Data["Website"] = "www.vkaifu.com"
	c.Data["Email"] = "vk@ivknet.com"
	c.TplName = "v1/insertmongodb.html"
}
//@router /vone/api [post]
func (c *VoneController)VoneApi(){
	fmt.Println("收到请求:"+c.Ctx.Request.RequestURI)
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin","*")
	service:=c.GetString("service")
	if service=="vone.setonsale"{
		result:=setonsale.Setonsale(c.GetString("device_info"),c.GetString("max"),c.GetString("mch_id"),c.GetString("min"),c.GetString("money"),c.GetString("number"),c.GetString("timestart"),c.GetString("timeend"))
		c.Data["json"]=&result
		c.ServeJSON()
	}else if service=="vone.insertmongodb"{
		result:=insertmongodb.InsertMongodb(c.GetString("mch_id"))
		c.Data["json"]=&result
		c.ServeJSON()
	}else if service=="vone.insertpassword"{
		result:=insertpassword.InsertPassword(c.GetString("password"),c.GetString("address"),c.GetString("device_info"),c.GetString("name"),c.GetString("prin"),c.GetString("mch_id"))
		c.Data["json"]=&result
		c.ServeJSON()
	}else if service=="vone.tinyurl"{
		result:=tinyurl.Tinyurl(c.GetString("device_info"),c.GetString("body"),c.GetString("mch_id"),c.GetString("tips"))
		c.Data["json"]=&result
		c.ServeJSON()
	}else{
		result:=result.Result{}
		result.Result_code="FAIL"
		result.Result_des="未知的服务名"
		c.Data["json"]=&result
		c.ServeJSON()
	}
}