package controller

import (
	"md/model/common"
	"md/model/entity"
	"md/service"

	"github.com/kataras/iris/v12"
)

func UpdatePassword(ctx iris.Context) {
	user := entity.User{}
	resolveParam(ctx, &user)
	service.SignUp(user)
	ctx.JSON(common.NewSuccess("更新成功"))
}
