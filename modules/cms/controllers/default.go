package controllers

import (
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/cms", &IndexController{})

}
