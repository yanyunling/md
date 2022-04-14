package service

import (
	"md/dao"
	"md/middleware"
	"md/model/common"
	"md/model/entity"
	"md/util"
	"time"

	"github.com/muesli/cache2go"
)

// 注册
func SignUp(user entity.User) {
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
	user.Password = util.EncryptSHA256([]byte(user.Id + user.Password))
	user.CreateTime = time.Now().UnixMilli()
	dao.UserAdd(tx, user)

	err = tx.Commit()
	if err != nil {
		panic(common.NewErr("注册失败", err))
	}
}

// 登录
func SignIn(user entity.User) common.TokenResult {
	// 去除用户名和密码的空白
	user.Name = util.RemoveBlank(user.Name)
	user.Password = util.RemoveBlank(user.Password)
	if user.Name == "" || user.Password == "" {
		panic(common.NewError("用户名或密码不可为空"))
	}

	// 根据用户名查询用户
	userResult, err := dao.UserGetByName(middleware.Db, user.Name)
	if err != nil {
		panic(common.NewErr("用户名或密码错误", err))
	}

	// 匹配密码：sha256(id + password)
	if util.EncryptSHA256([]byte(userResult.Id+user.Password)) != userResult.Password {
		panic(common.NewError("用户名或密码错误"))
	}

	// 生成token
	tokenResult := common.TokenResult{}
	tokenResult.Name = userResult.Name
	tokenResult.AccessToken = util.RandomString(64)
	tokenResult.RefreshToken = util.RandomString(64)

	// 缓存token
	tokenCache := common.TokenCache{}
	tokenCache.Id = userResult.Id
	tokenCache.Name = userResult.Name
	tokenCache.AccessToken = tokenResult.AccessToken
	tokenCache.RefreshToken = tokenResult.RefreshToken

	cache2go.Cache(common.AccessTokenCache).Add(tokenResult.AccessToken, time.Hour, &tokenCache)
	cache2go.Cache(common.RefreshTokenCache).Add(tokenResult.RefreshToken, time.Hour*24*180, &tokenCache)

	return tokenResult
}
