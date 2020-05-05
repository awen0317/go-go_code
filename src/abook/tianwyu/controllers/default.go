package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	beego.BConfig.AppName ="beepkg"
	c.Data["Website"] = "beego.me"+beego.BConfig.AppName
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
