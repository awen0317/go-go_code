package main

import (
	"github.com/astaxie/beego"
	_ "ziyoubiancheng/mbook/routers"
	_ "ziyoubiancheng/mbook/sysinit"
)

func main() {
	beego.Run()
}
