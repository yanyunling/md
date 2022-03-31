package dao

import (
	"md/model/common"
	"md/model/entity"
	"md/util"

	"github.com/jmoiron/sqlx"
)

// 添加用户
func UserAdd(tx *sqlx.Tx, user entity.User) error {
	_, err := tx.NamedExec(`insert into t_user (id,name,password,phone,email,portrait,remark,status,create_time,update_time)
	values (:id,:name,:password,:phone,:email,:portrait,:remark,:status,:create_time,:update_time)`, user)
	return err
}

// 修改用户
func UserUpdate(tx *sqlx.Tx, user entity.User) error {
	_, err := tx.NamedExec(`update t_user set name=:name,phone=:phone,email=:email,portrait=:portrait,
	remark=:remark,status=:status,update_time=:update_time where id=:id`, user)
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

// 修改用户密码
func UserResetPassword(tx *sqlx.Tx, user entity.User) error {
	_, err := tx.NamedExec(`update t_user set password=:password,update_time=:update_time where id=:id`, user)
	return err
}

// 分页查询用户
func UserListByPage(db *sqlx.DB, userCon entity.UserCon) (common.Page, error) {
	sc := util.SqlCompletion{}
	sc.InitSql(`select * from t_user`, `select count(*) as total from t_user`)
	if userCon.Name != "" {
		sc.Like("name", userCon.Name, true)
	}
	sc.Order("name", true)
	sc.Limit(userCon.Page, userCon.Size)

	// 查询总记录数
	result := common.Page{}
	err := db.Get(&result, sc.GetCountSql(), sc.GetCountParams()...)

	return result, err
}
