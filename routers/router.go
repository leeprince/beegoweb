package routers

import (
	"beegoweb/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/hello", &controllers.HelloController{})
    beego.Router("/login", &controllers.LoginController{})
}
