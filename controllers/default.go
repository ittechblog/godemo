package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

// @Title home
// @Description home
// @Success 200 {object}
// @Failure 404 User not found
// @router /
func (c *MainController) Home() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}


func (c *MainController) Json() {
	var jsonMap map[string]interface{} = map[string]interface{}{}
	jsonMap["success"] = "0"
	c.Data["json"] = jsonMap
	c.ServeJSON()
	return
}