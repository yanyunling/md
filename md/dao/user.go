package dao

import (
	"md/model/common"
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

// 根据名字查询用户
func UserGetByName(tx *sqlx.Tx, name string) (entity.User, error) {
	result := entity.User{}
	err := tx.Get(&result, "select * from t_user where name=?", name)
	return result, err
}

// 根据用户名查询用户数量
func UserCountByName(tx *sqlx.Tx, name string) (common.CommonResult, error) {
	result := common.CommonResult{}
	err := tx.Get(&result, "select count(*) as count from t_user where name=?", name)
	return result, err
}

// 查询用户数量
func UserCount(tx *sqlx.Tx) (common.CommonResult, error) {
	result := common.CommonResult{}
	err := tx.Get(&result, "select count(*) as count from t_user")
	return result, err
}
