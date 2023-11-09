package dao

import (
	"md/model/entity"
	"md/util"
	"sort"

	"github.com/jmoiron/sqlx"
)

// 添加文档
func DocumentAdd(tx *sqlx.Tx, document entity.Document) error {
	sql := `insert into t_document (id,name,content,type,published,create_time,update_time,book_id,user_id) values (:id,:name,:content,:type,:published,:create_time,:update_time,:book_id,:user_id)`
	_, err := tx.NamedExec(sql, document)
	return err
}

// 修改文档基础信息
func DocumentUpdate(tx *sqlx.Tx, document entity.Document) error {
	sql := `update t_document set name=:name,published=:published,book_id=:book_id where id=:id and user_id=:user_id`
	_, err := tx.NamedExec(sql, document)
	return err
}

// 修改文档内容
func DocumentUpdateContent(tx *sqlx.Tx, document entity.Document) error {
	sql := `update t_document set content=:content,update_time=:update_time where id=:id and user_id=:user_id`
	_, err := tx.NamedExec(sql, document)
	return err
}

// 根据id删除文档
func DocumentDeleteById(tx *sqlx.Tx, id, userId string) error {
	sql := `delete from t_document where id=? and user_id=?`
	_, err := tx.Exec(sql, id, userId)
	return err
}

// 查询文档列表
func DocumentList(db *sqlx.DB, bookId, userId string) ([]entity.Document, error) {
	sqlCompletion := util.SqlCompletion{}
	sqlCompletion.InitSql(`select id,name,type,published,create_time,update_time,book_id from t_document`, ``)
	sqlCompletion.Eq("user_id", userId, true)
	if bookId != "" {
		sqlCompletion.Eq("book_id", bookId, true)
	}
	result := []entity.Document{}
	err := db.Select(&result, sqlCompletion.GetSql(), sqlCompletion.GetParams()...)
	// 按名称升序
	sort.Slice(result, func(i, j int) bool {
		return util.StringSort(result[i].Name, result[j].Name)
	})
	return result, err
}

// 根据id查询文档
func DocumentGetById(db *sqlx.DB, id, userId string) (entity.Document, error) {
	sql := `select id,name,content,type,published,create_time,update_time,book_id from t_document where id=? and user_id=?`
	result := entity.Document{}
	err := db.Get(&result, sql, id, userId)
	return result, err
}

// 清空文档的bookId
func DocumentClearBookId(tx *sqlx.Tx, bookId string) error {
	sql := `update t_document set book_id='' where book_id=?`
	_, err := tx.Exec(sql, bookId)
	return err
}

// 根据id查询公开发布文档
func DocumentGetPublished(db *sqlx.DB, id string) (entity.Document, error) {
	sql := `select id,name,content,type,published,create_time,update_time,book_id from t_document where id=? and published=true`
	result := entity.Document{}
	err := db.Get(&result, sql, id)
	return result, err
}
