//create: 2017/10/23 19:13:17 change: 2017/10/24 10:38:22 author:lijiao
package routers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/lijiaocn/study-beego/hello/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	nsv1 := beego.NewNamespace(
		"/v1",
		beego.NSCond(func(ctx *context.Context) bool {
			fmt.Println("/v1  NSCond")
			return true
		}),
		beego.NSBefore(func(ctx *context.Context) {
			fmt.Println("/v1  Before1")
		}),
		beego.NSBefore(func(ctx *context.Context) {
			fmt.Println("/v1  Before2")
		}),
		beego.NSGet("/version", func(ctx *context.Context) {
			//if call this function, NSAfter will not work in beego 1.9.0
			//ctx.Output.Body([]byte("v1/version"))
			fmt.Println("/v1  version")
		}),
		beego.NSAfter(func(ctx *context.Context) {
			fmt.Println("/v1  After1")
		}),
		beego.NSAfter(func(ctx *context.Context) {
			fmt.Println("/v1  After2")
		}),
		beego.NSNamespace("user",
			beego.NSGet("/:id", func(ctx *context.Context) {
				ctx.Output.JSON(ctx.Input.Param(":id"), true, false)
			}),
		),
	)
	nsv2 := beego.NewNamespace(
		"/v2",
		beego.NSCond(func(ctx *context.Context) bool {
			fmt.Println("/v2  NSCond")
			return true
		}),
		beego.NSBefore(func(ctx *context.Context) {
			fmt.Println("/v2  Before1")
			return
		}),
		beego.NSBefore(func(ctx *context.Context) {
			fmt.Println("/v2  Before2")
		}),
		beego.NSAfter(func(ctx *context.Context) {
			fmt.Println("/v2  After1")
		}),
		beego.NSGet("/version", func(ctx *context.Context) {
			ctx.Output.Body([]byte("v2/version"))
			return
		}),
		beego.NSAfter(func(ctx *context.Context) {
			fmt.Println("/v2  After2")
		}),
		beego.NSNamespace("user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(nsv1, nsv2)
}
