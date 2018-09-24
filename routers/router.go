package routers

import (
	"my_blog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//添加注解路由
	ns := beego.NewNamespace("v1",
		beego.NSNamespace("test",
			beego.NSInclude(
				&controllers.TestController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
