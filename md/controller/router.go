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

			token.Post("/sign-up", SignUp)
			token.Post("/sign-in", SignIn)
			token.Post("/sign-out", SignOut)
			token.Post("/refresh", TokenRefresh)
		})

		// 数据接口
		api.PartyFunc("/data", func(data iris.Party) {
			data.Use(middleware.DataAuth)

			data.PartyFunc("/user", func(user iris.Party) {
				user.Post("/update-password", UserUpdatePassword)
			})

			data.PartyFunc("/doc", func(doc iris.Party) {

			})

			data.PartyFunc("/pic", func(pic iris.Party) {
				pic.Post("/page", PicturePage)
				pic.Post("/delete", PictureDelete)
				pic.Post("/upload", PictureUpload)
			})
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
