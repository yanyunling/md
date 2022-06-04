package service

import (
	"md/dao"
	"md/middleware"
	"md/model/common"
	"md/model/entity"
	"os"
)

// 分页查询图片记录
func PicturePage(pageCondition common.PageCondition[interface{}], userId string) common.PageResult[entity.PicturePageResult] {
	pictures, total, err := dao.PicturePage(middleware.Db, pageCondition.Page, userId)
	if err != nil {
		panic(common.NewErr("查询失败", err))
	}
	picturePageResults := []entity.PicturePageResult{}
	for _, v := range pictures {
		picturePageResults = append(picturePageResults, entity.PicturePageResult{Picture: v, PicturePrefix: "/" + common.ResourceName + "/" + common.PictureName + "/", ThumbnailPrefix: "/" + common.ResourceName + "/" + common.ThumbnailName + "/"})
	}
	pageResult := common.PageResult[entity.PicturePageResult]{Records: picturePageResults, Total: total}
	return pageResult
}

// 删除图片
func PictureDelete(id, userId string) {
	tx := middleware.Db.MustBegin()
	defer tx.Rollback()

	// 查询图片
	picture, err := dao.PictureGetById(tx, id, userId)
	if err != nil {
		panic(common.NewErr("删除失败", err))
	}

	// 查询相同大小、相同hash的图片数量
	countResult, err := dao.PictureCountBySizeHash(tx, picture.Size, picture.Hash)
	if err != nil {
		panic(common.NewErr("删除失败", err))
	}

	// 删除记录
	err = dao.PictureDeleteById(tx, id, userId)
	if err != nil {
		panic(common.NewErr("删除失败", err))
	}

	// 如果相同的图片只有一条记录，删除文件
	if countResult.Count == 1 {
		os.Remove(common.DataPath + common.ResourceName + "/" + common.PictureName + "/" + picture.Path)
		os.Remove(common.DataPath + common.ResourceName + "/" + common.ThumbnailName + "/" + picture.Path)
	}

	err = tx.Commit()
	if err != nil {
		panic(common.NewErr("删除失败", err))
	}
}
