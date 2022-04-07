package middleware

import (
	"md/model/common"
	"md/util"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/kataras/iris/v12"
	"github.com/muesli/cache2go"
)

// 用户接口授权
func UserAuth(ctx iris.Context) {
	// header中的token
	token := resolveHeader(ctx, "Bearer")

	// 检验缓存中是否存在此token
	if !cache2go.Cache(common.AccessTokenCache).Exists(token) {
		panic(common.NewErrorCode(common.HttpAuthFailure, "认证失败"))
	}

	// 自动续期
	cache2go.Cache(common.AccessTokenCache).Add(token, time.Hour*2, true)

	ctx.Next()
}

// token认证授权
func TokenAuth(ctx iris.Context) {
	// header中的token
	token := resolveHeader(ctx, "Basic")

	_, err := util.DecryptBASE64(token)
	if err != nil {
		common.NewErrCode(common.HttpAuthFailure, "头信息认证失败", err)
	}

	ctx.Next()
}

// 解析头信息中的认证信息
func resolveHeader(ctx iris.Context, prefix string) string {
	header := ctx.GetHeader("Authorization")
	prefixLen := utf8.RuneCountInString(prefix) + 1
	if header != "" && strings.Index(header, prefix) == 0 && utf8.RuneCountInString(header) > prefixLen {
		return string([]rune(header)[prefixLen:])
	}
	panic(common.NewErrorCode(common.HttpAuthFailure, "头信息认证失败"))
}
