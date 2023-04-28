package service

import (
	"md/dao"
	"md/middleware"
	"md/model/common"
	"md/model/entity"
	"md/util"
	"strings"
	"time"
)

// 添加文档
func DocumentAdd(document entity.Document) entity.Document {
	tx := middleware.Db.MustBegin()
	defer tx.Rollback()

	document.Name = strings.TrimSpace(document.Name)
	if document.Name == "" {
		panic(common.NewError("文档名称不可为空"))
	}
	if util.StringLength(document.Name) > 1000 {
		panic(common.NewError("文档名称过长，请小于1000个字符"))
	}
	if util.StringLength(document.Content) > 1000000 {
		panic(common.NewError("文档内容过多，请小于100万个字符"))
	}
	document.Id = util.SnowflakeString()
	document.CreateTime = time.Now().UnixMilli()
	document.UpdateTime = time.Now().UnixMilli()
	err := dao.DocumentAdd(tx, document)
	if err != nil {
		panic(common.NewErr("添加失败", err))
	}

	err = tx.Commit()
	if err != nil {
		panic(common.NewErr("添加失败", err))
	}

	return document
}

// 修改文档基础信息
func DocumentUpdate(document entity.Document) {
	tx := middleware.Db.MustBegin()
	defer tx.Rollback()

	document.Name = strings.TrimSpace(document.Name)
	if document.Name == "" {
		panic(common.NewError("文档名称不可为空"))
	}
	if util.StringLength(document.Name) > 1000 {
		panic(common.NewError("文档名称过长，请小于1000个字符"))
	}
	err := dao.DocumentUpdate(tx, document)
	if err != nil {
		panic(common.NewErr("更新失败", err))
	}

	err = tx.Commit()
	if err != nil {
		panic(common.NewErr("更新失败", err))
	}
}

// 修改文档内容
func DocumentUpdateContent(document entity.Document) entity.Document {
	tx := middleware.Db.MustBegin()
	defer tx.Rollback()

	if util.StringLength(document.Content) > 1000000 {
		panic(common.NewError("文档内容过多，请小于100万个字符"))
	}
	document.UpdateTime = time.Now().UnixMilli()
	err := dao.DocumentUpdateContent(tx, document)
	if err != nil {
		panic(common.NewErr("更新失败", err))
	}

	err = tx.Commit()
	if err != nil {
		panic(common.NewErr("更新失败", err))
	}

	return document
}

// 删除文档
func DocumentDelete(id, userId string) {
	tx := middleware.Db.MustBegin()
	defer tx.Rollback()

	err := dao.DocumentDeleteById(tx, id, userId)
	if err != nil {
		panic(common.NewErr("删除失败", err))
	}

	err = tx.Commit()
	if err != nil {
		panic(common.NewErr("删除失败", err))
	}
}

// 查询文档列表
func DocumentList(bookId, userId string) []entity.Document {
	documents, err := dao.DocumentList(middleware.Db, bookId, userId)
	if err != nil {
		panic(common.NewErr("查询失败", err))
	}
	return documents
}

// 查询文档
func DocumentGet(id, userId string) entity.Document {
	document, err := dao.DocumentGetById(middleware.Db, id, userId)
	if err != nil {
		panic(common.NewErr("查询失败", err))
	}
	return document
}
