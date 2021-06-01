package main

import (
	"beegoweb/constants"
	_ "beegoweb/routers"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func main() {
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser) // 在 beego.Run() 之前
	
	beego.Run()
}

var FilterUser = func(ctx *context.Context) {
	s, ok := ctx.Input.Session(constants.UserSessionKey).(int)
	fmt.Println("FilterUser>", s, ok)
	
	if !ok && ctx.Request.RequestURI != "/login" {
		ctx.Redirect(302, "/login")
	}
}
