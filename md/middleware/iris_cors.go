package middleware

import "github.com/kataras/iris/v12"

// CORS跨域
func CORS(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "x-requested-with, Authorization, Content-Type, Accept")
	ctx.Header("Access-Control-Max-Age", "3600")
	ctx.Next()
}
