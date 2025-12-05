package controller

import (
	"md/middleware"
	"md/model/common"
	"md/service"

	"github.com/kataras/iris/v12"
)

// SSE处理器
func SSEHandler(ctx iris.Context) {
	// 设置SSE请求头
	ctx.ContentType("text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")

	service.SSEHandler(ctx)
}

// 获取SSE临时token
func GetSSEToken(ctx iris.Context) {
	userId := middleware.CurrentUserId(ctx)
	token := service.GetSSEToken(userId)
	ctx.JSON(common.NewSuccessData("获取成功", token))
}
