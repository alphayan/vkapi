package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.vkaifu.com"
	c.Data["Email"] = "vk@ivknet.com"
	c.TplName = "index.tpl"
}
