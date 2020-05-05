package controllers

import (
	"fmt"
	"ziyoubiancheng/mbook/models"

	"github.com/astaxie/beego"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Index() {
	if cates, err := new(models.Category).GetCates(-1, 1); err == nil {
		fmt.Println(cates)
		c.Data["Cates"] = cates
	} else {
		beego.Error(err.Error())
	}

	c.TplName = "home/list.html"
}

func (c *HomeController) Index2() {
	c.TplName = "home/list.html"
}
