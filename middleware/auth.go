package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"md/model/common"
	"md/util"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/kataras/iris/v12"
	"github.com/muesli/cache2go"
)

// 管理页面接口授权
func ManagerAuth(ctx iris.Context) {
	// header中的token
	token := resolveHeader(ctx, "Bearer")

	// 检验缓存中是否存在此token
	if !cache2go.Cache(common.ManagerTokenCache).Exists(token) {
		panic(common.NewErrorCode(common.HttpAuthFailure, "认证失败"))
	}

	// 自动续期
	cache2go.Cache(common.ManagerTokenCache).Add(token, time.Hour*2, true)

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

// AppSecret接口授权
func AppAuth(ctx iris.Context) {

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

// 解析body数据，执行ctx.Next()之前必须重新填充body数据，否则后续路由将无法解析参数
func resolveBody(ctx iris.Context, con interface{}) {
	if ctx.Request().Body == nil {
		panic(common.NewErrorCode(common.HttpAuthFailure, "参数解析失败"))
	}

	bodyBytes, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		panic(common.NewErrCode(common.HttpAuthFailure, "参数解析失败", err))
	}

	if err = json.Unmarshal(bodyBytes, &con); err != nil {
		panic(common.NewErrCode(common.HttpAuthFailure, "参数解析失败", err))
	}
}

// 重新填充body数据
func recoverBody(ctx iris.Context, bodyBytes []byte) {
	ctx.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
}
