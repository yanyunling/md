package controller

import (
	"md/model/common"

	"github.com/kataras/iris/v12"
)

// 初始化iris路由
func InitRouter(app *iris.Application) {
	app.PartyFunc("/api", func(api iris.Party) {
		api.Post("/user/register", UserRegister)

		// token相关接口
		api.PartyFunc("/token", func(token iris.Party) {

		})
	})
}

// 解析参数
func resolveParam(ctx iris.Context, con interface{}) {
	err := ctx.ReadJSON(&con)
	if err != nil {
		panic(common.NewErr("参数解析失败", err))
	}
}
