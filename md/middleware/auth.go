package middleware

import (
	"md/model/common"
	"md/util"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/kataras/iris/v12"
	"github.com/muesli/cache2go"
)

// 获取当前登录用户信息
func CurrentUser(ctx iris.Context) {
	token := resolveHeader(ctx, "Bearer")

	cache2go.Cache(token)
}

// 数据接口授权
func DataAuth(ctx iris.Context) {
	token := resolveHeader(ctx, "Bearer")

	// 检验缓存中是否存在此token
	if !cache2go.Cache(common.AccessTokenCache).Exists(token) {
		panic(common.NewErrorCode(common.HttpAuthFailure, "认证失败"))
	}

	ctx.Next()
}

// token相关接口认证授权
func TokenAuth(ctx iris.Context) {
	token := resolveHeader(ctx, "Basic")

	// SHA256（BasicTokenKey + 时间戳的10分钟为基准的值，可上下浮动10分钟）
	current := time.Now().UnixMilli() / 600000
	t1 := util.EncryptSHA256([]byte(common.BasicTokenKey + strconv.FormatInt(current, 10)))
	t2 := util.EncryptSHA256([]byte(common.BasicTokenKey + strconv.FormatInt(current-1, 10)))
	t3 := util.EncryptSHA256([]byte(common.BasicTokenKey + strconv.FormatInt(current+1, 10)))

	if token != t1 && token != t2 && token != t3 {
		panic(common.NewErrorCode(common.HttpAuthFailure, "认证失败"))
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
	panic(common.NewErrorCode(common.HttpAuthFailure, "认证失败"))
}
