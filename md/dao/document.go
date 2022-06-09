package dao

import (
	"md/model/entity"

	"github.com/jmoiron/sqlx"
)

// 添加文档
func DocumentAdd(tx *sqlx.Tx, document entity.Document) error {
	sql := `insert into t_document (id,name,content,create_time,update_time,book_id,user_id) values (:id,:name,:content,:create_time,:update_time,:book_id,:user_id)`
	_, err := tx.NamedExec(sql, document)
	return err
}

// 修改文档
func DocumentUpdate(tx *sqlx.Tx, document entity.Document) error {
	sql := `update t_document set name=:name,content=:content,update_time=:update_time,book_id=:book_id where id=:id and user_id=:user_id`
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
	sql := `select id,name,create_time,update_time,book_id from t_document where book_id=? and user_id=? order by create_time desc`
	result := []entity.Document{}
	err := db.Select(&result, sql, bookId, userId)
	return result, err
}

// 根据id查询文档
func DocumentGetById(db *sqlx.DB, id, userId string) (entity.Document, error) {
	sql := `select id,name,content,create_time,update_time,book_id from t_document where id=? and user_id=?`
	result := entity.Document{}
	err := db.Get(&result, sql, id, userId)
	return result, err
}
