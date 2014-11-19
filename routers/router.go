package routers

import (
	"github.com/fromruijin/gocms/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
