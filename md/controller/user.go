package controller

import (
	"md/model/common"
	"md/model/entity"
	"md/service"

	"github.com/kataras/iris/v12"
)

// 重置密码
func UserRegister(ctx iris.Context) {
	user := entity.User{}
	resolveParam(ctx, &user)
	service.UserRegister(user)
	ctx.JSON(common.NewSuccess("注册成功"))
}
