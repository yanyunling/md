package controller

import (
	"md/middleware"
	"md/model/common"

	"github.com/kataras/iris/v12"
)

// 初始化iris路由
func InitRouter(app *iris.Application) {
	app.PartyFunc("/api", func(api iris.Party) {
		// token相关接口
		api.PartyFunc("/token", func(token iris.Party) {
			token.Use(middleware.TokenAuth)
			token.Post("/register", SignUp)
			token.Post("/get", SignIn)
			token.Post("/remove", SignOut)
			token.Post("/refresh", TokenRefresh)
		})

		// 数据接口
		api.PartyFunc("/data", func(data iris.Party) {
			data.Use(middleware.DataAuth)
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
