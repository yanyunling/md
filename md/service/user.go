package service

import (
	"md/dao"
	"md/middleware"
	"md/model/common"
	"md/model/entity"
	"md/util"
	"time"
)

// 注册
func UserRegister(user entity.User) {
	tx := middleware.Db.MustBegin()
	defer tx.Rollback()

	// 如不允许注册，查询是否没有任何用户
	if !common.Register {
		commonResult, err := dao.UserCount(tx)
		if err != nil {
			panic(common.NewErr("注册失败", err))
		}
		if commonResult.Count > 0 {
			panic(common.NewError("暂不支持注册"))
		}
	}

	// 去除用户名和密码的空白
	user.Name = util.RemoveBlank(user.Name)
	user.Password = util.RemoveBlank(user.Password)
	if user.Name == "" || user.Password == "" {
		panic(common.NewError("用户名或密码不可为空"))
	}

	// 用户名长度限制
	if util.StringLength(user.Name) > 30 {
		panic(common.NewError("用户名不可大于30个字符"))
	}

	// 查询用户名不可重复
	commonResult, err := dao.UserCountByName(tx, user.Name)
	if err != nil {
		panic(common.NewErr("注册失败", err))
	}
	if commonResult.Count > 0 {
		panic(common.NewError("用户名已被注册"))
	}

	// 保存用户信息
	user.Id = util.SnowflakeString()
	user.CreateTime = time.Now().UnixMilli()
	dao.UserAdd(tx, user)

	err = tx.Commit()
	if err != nil {
		panic(common.NewErr("注册失败", err))
	}
}
