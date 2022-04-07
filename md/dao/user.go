package dao

import (
	"md/model/entity"

	"github.com/jmoiron/sqlx"
)

// 添加用户
func UserAdd(tx *sqlx.Tx, user entity.User) error {
	_, err := tx.NamedExec(`insert into t_user (id,name,password,create_time)
	values (:id,:name,:password,:create_time)`, user)
	return err
}

// 修改用户
func UserUpdate(tx *sqlx.Tx, user entity.User) error {
	_, err := tx.NamedExec(`update t_user set name=:name where id=:id`, user)
	return err
}

// 修改用户密码
func UserResetPassword(tx *sqlx.Tx, user entity.User) error {
	_, err := tx.NamedExec(`update t_user set password=:password where id=:id`, user)
	return err
}

// 根据id删除用户
func UserDeleteById(tx *sqlx.Tx, id string) error {
	_, err := tx.Exec("delete from t_user where id=?", id)
	return err
}

// 根据id查询用户
func UserGetById(tx *sqlx.Tx, id string) (entity.User, error) {
	result := entity.User{}
	err := tx.Get(&result, "select * from t_user where id=?", id)
	return result, err
}
