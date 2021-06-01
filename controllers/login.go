package controllers

import (
	"beegoweb/constants"
	"beegoweb/models"
	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.tpl"
}

func (c *LoginController) Post() {
	username := c.GetString("username")
	user, err := models.FindUserByUserName(username)
	
	result := make(map[string]interface{})
	result["code"] = 0
	result["data"] = make(map[string]interface{})
	result["message"] = "ok"
	
	if err != nil {
		result["code"] = 1
		result["message"] = err.Error()
		c.Data["json"] = result
		c.ServeJSON() // c.Ctx.Output.JSON(result,false,false)
		return
	}
	
	// 保存 session 登录态
	c.SetSession(constants.UserSessionKey, user.Id)
	
	result["data"] = user
	c.Data["json"] = result
	c.ServeJSON()
}
