// @APIVersion 1.0.0
// @Title mobile API
// @Description mobile has every tool to get any job done, so codename for the new mobile APIs.
// @Contact astaxie@gmail.com
package routers

import (
	"igen.com/bee/controllers"
	"github.com/astaxie/beego"
)

func init() {
    	beego.Router("/", &controllers.MainController{},"get:Home")
	beego.Router("/Json", &controllers.MainController{},"get:Json")
	beego.Router("/write", &controllers.WriteController{})
}
