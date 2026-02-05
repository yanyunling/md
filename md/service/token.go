package service

import (
	"fmt"
	"md/dao"
	"md/middleware"
	"md/model/common"
	"md/model/entity"
	"md/util"
	"time"

	"github.com/wenlng/go-captcha/v2/slide"
)

const AccessTokenExpire = time.Hour
const RefreshTokenExpire = time.Hour * 24 * 30
const CaptchaExpire = time.Minute * 5

// 注册
func SignUp(signIn common.SignIn) {
	tx := middleware.Db.MustBegin()
	defer tx.Rollback()

	// 校验验证码
	ok := CaptchaValidate(signIn)
	if !ok {
		panic(common.NewError("验证失败"))
	}
	// 移除验证码缓存
	middleware.Cache.Delete(common.CaptchaCache + signIn.CaptchaId)

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

	// 去除用户名的空白
	signIn.Name = util.RemoveBlank(signIn.Name)
	if signIn.Name == "" || signIn.Password == "" {
		panic(common.NewError("用户名或密码不可为空"))
	}

	// 用户名长度限制
	if util.StringLength(signIn.Name) > 30 {
		panic(common.NewError("用户名不可大于30个字符"))
	}

	// 查询用户名不可重复
	commonResult, err := dao.UserCountByName(tx, signIn.Name)
	if err != nil {
		panic(common.NewErr("注册失败", err))
	}
	if commonResult.Count > 0 {
		panic(common.NewError("用户名已被注册"))
	}

	// 保存用户信息
	user := entity.User{}
	user.Id = util.SnowflakeString()
	user.Name = signIn.Name
	user.Password = util.EncryptSHA256([]byte(user.Id + user.Password))
	user.CreateTime = time.Now().UnixMilli()
	dao.UserAdd(tx, user)

	err = tx.Commit()
	if err != nil {
		panic(common.NewErr("注册失败", err))
	}
}

// 登录
func SignIn(signIn common.SignIn) common.TokenResult {
	// 去除用户名的空白
	signIn.Name = util.RemoveBlank(signIn.Name)
	if signIn.Name == "" || signIn.Password == "" {
		panic(common.NewError("用户名或密码不可为空"))
	}

	// 校验登录次数
	checkSignInTimes(signIn.Name)

	// 校验验证码
	ok := CaptchaValidate(signIn)
	if !ok {
		panic(common.NewError("验证失败"))
	}
	// 移除验证码缓存
	middleware.Cache.Delete(common.CaptchaCache + signIn.CaptchaId)

	// 根据用户名查询用户
	userResult, err := dao.UserGetByName(middleware.Db, signIn.Name)
	if err != nil {
		panic(common.NewErr("用户名或密码错误", err))
	}

	// 匹配密码：sha256(id + password)
	if util.EncryptSHA256([]byte(userResult.Id+signIn.Password)) != userResult.Password {
		panic(common.NewError("用户名或密码错误"))
	}

	// 生成token
	tokenResult := common.TokenResult{}
	tokenResult.Name = userResult.Name
	tokenResult.AccessToken = util.RandomString(64)
	tokenResult.RefreshToken = util.RandomString(64)

	tokenCache := common.TokenCache{}
	tokenCache.Id = userResult.Id
	tokenCache.TokenResult = tokenResult

	// 缓存token
	middleware.Cache.Set(common.AccessTokenCache+tokenResult.AccessToken, &tokenCache, AccessTokenExpire)
	middleware.Cache.Set(common.RefreshTokenCache+tokenResult.RefreshToken, &tokenCache, RefreshTokenExpire)
	middleware.Cache.Delete(common.SignInTimesCache + signIn.Name)

	return tokenResult
}

// 退出登录
func SignOut(tokenResult common.TokenResult) {
	res, found := middleware.Cache.Get(common.RefreshTokenCache + tokenResult.RefreshToken)
	if found {
		tokenCache := res.(*common.TokenCache)
		if tokenCache.RefreshToken != "" {
			middleware.Cache.Delete(common.RefreshTokenCache + tokenCache.RefreshToken)
		}
		if tokenCache.AccessToken != "" {
			middleware.Cache.Delete(common.AccessTokenCache + tokenCache.AccessToken)
		}
	}
}

// 刷新token
func TokenRefresh(refreshToken string) common.TokenResult {
	res, found := middleware.Cache.Get(common.RefreshTokenCache + refreshToken)
	if !found {
		panic(common.NewError("认证信息已过期，请重新登录"))
	}
	tokenCache := res.(*common.TokenCache)
	if tokenCache.RefreshToken == "" {
		panic(common.NewError("认证信息已过期，请重新登录"))
	}

	// 重新生成token
	tokenResult := common.TokenResult{}
	tokenResult.Name = tokenCache.Name
	tokenResult.AccessToken = util.RandomString(64)
	tokenResult.RefreshToken = util.RandomString(64)

	newTokenCache := common.TokenCache{}
	newTokenCache.Id = tokenCache.Id
	newTokenCache.TokenResult = tokenResult

	// 缓存token
	middleware.Cache.Set(common.AccessTokenCache+newTokenCache.AccessToken, &newTokenCache, AccessTokenExpire)
	middleware.Cache.Set(common.RefreshTokenCache+newTokenCache.RefreshToken, &newTokenCache, RefreshTokenExpire)

	return tokenResult
}

// 获取验证码
func Captcha(signIn common.SignIn) middleware.CaptchaResult {
	// 传入的id是上一次的验证码，此操作为刷新验证码，移除上一次的缓存
	if signIn.CaptchaId != "" {
		middleware.Cache.Delete(common.CaptchaCache + signIn.CaptchaId)
	}

	// 生成验证码
	captchaResult, err := middleware.CaptchaGet()
	captchaResult.CaptchaId = util.Uuid()
	if err != nil {
		panic(common.NewErr("验证码生成失败", err))
	}

	// 缓存验证信息
	captchaCache := middleware.CaptchaCache{X: captchaResult.X, Y: captchaResult.Y}
	middleware.Cache.Set(common.CaptchaCache+captchaResult.CaptchaId, &captchaCache, CaptchaExpire)

	return captchaResult
}

// 校验验证码
func CaptchaValidate(signIn common.SignIn) bool {
	// 根据验证码id从缓存获取验证码信息
	res, found := middleware.Cache.Get(common.CaptchaCache + signIn.CaptchaId)
	if !found {
		return false
	}
	captchaCache := res.(*middleware.CaptchaCache)

	// 校验
	ok := slide.Validate(signIn.CaptchaX, signIn.CaptchaY, captchaCache.X, captchaCache.Y, 6)

	// 移除验证码缓存
	if !ok {
		middleware.Cache.Delete(common.CaptchaCache + signIn.CaptchaId)
	}

	return ok
}

// 校验登录次数，如已超出则抛出异常
func checkSignInTimes(name string) {
	times := 1
	signInTimes, exp, found := middleware.Cache.GetWithExpiration(common.SignInTimesCache + name)
	now := time.Now()
	if !found || exp.Before(now) {
		middleware.Cache.Set(common.SignInTimesCache+name, times, time.Minute*5)
	} else {
		times = signInTimes.(int) + 1
		middleware.Cache.Set(common.SignInTimesCache+name, times, exp.Sub(now))
	}
	if times > 5 {
		panic(common.NewError(fmt.Sprintf("登录次数已达上限，请于%v分钟后再试", int(exp.Sub(now).Minutes())+1)))
	}
}
