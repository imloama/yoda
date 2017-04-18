package main

import (
	_ "yoda/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.Listen.EnableAdmin = true
	beego.BConfig.Listen.AdminAddr = "localhost"
	beego.BConfig.Listen.AdminPort = 8088
	beego.Run()
}
