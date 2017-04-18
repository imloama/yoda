package routers

import (
	"yoda/controllers"
	_ "yoda/modules/cms/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

}
