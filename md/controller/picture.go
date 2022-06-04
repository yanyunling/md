package controller

import (
	"md/middleware"
	"md/model/common"
	"md/model/entity"
	"md/service"

	"github.com/kataras/iris/v12"
)

// 分页查询图片记录
func PicturePage(ctx iris.Context) {
	pageCondition := common.PageCondition[interface{}]{}
	resolveParam(ctx, &pageCondition)
	userId := middleware.CurrentUserId(ctx)
	ctx.JSON(common.NewSuccessData("查询成功", service.PicturePage(pageCondition, userId)))
}

// 删除图片
func PictureDelete(ctx iris.Context) {
	picture := entity.Picture{}
	resolveParam(ctx, &picture)
	userId := middleware.CurrentUserId(ctx)
	service.PictureDelete(picture.Id, userId)
	ctx.JSON(common.NewSuccess("删除成功"))
}

// 上传图片
func PictureUpload(ctx iris.Context) {
	//userId := middleware.CurrentUserId(ctx)
	ctx.JSON(common.NewSuccess("上传成功"))
}
