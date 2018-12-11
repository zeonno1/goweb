package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["123654
"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Name"] = "a new app"
	c.Data["Server"] = "https"
	c.TplName = "index.tpl"
}
