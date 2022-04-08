package main

import (
	"embed"
	"flag"
	"io/fs"
	"md/controller"
	"md/middleware"
	"md/model/common"
	"md/util"
	"net/http"
	"time"

	"github.com/kataras/iris/v12"
)

//go:embed web
var web embed.FS

func init() {
	// 解析命令行参数
	flag.StringVar(&common.Port, "p", "9900", "监听端口")
	flag.StringVar(&common.LogPath, "log", "./logs", "日志目录")
	flag.StringVar(&common.DbPath, "db", "./", "数据库目录")
	flag.BoolVar(&common.Register, "reg", true, "允许注册(即使禁止注册，在没有任何用户的情况时仍可注册)")
	flag.Parse()
}

func main() {
	// 创建iris服务
	app := iris.New()

	// 初始化日志
	middleware.InitLog(common.LogPath, app.Logger())

	// 全局异常恢复
	app.Use(middleware.GlobalRecover)

	// gzip压缩
	app.Use(iris.Compression)

	// 预检请求
	app.Options("*", func(ctx iris.Context) {
		ctx.Next()
	})

	// CORS跨域
	app.UseGlobal(middleware.CORS)

	// 初始化雪花算法节点
	err := util.InitSnowflake(0)
	if err != nil {
		middleware.Log.Error("初始化雪花算法节点失败：", err)
		return
	}

	// 初始化数据库
	err = middleware.InitSqlite(common.DbPath)
	if err != nil {
		return
	}

	// 初始化API路由
	controller.InitRouter(app)

	// 网页静态资源路由
	webFs, err := fs.Sub(web, "web")
	if err == nil {
		app.Use(iris.StaticCache(time.Hour * 720))
		app.HandleDir("/", http.FS(webFs))
	}

	// 启动服务
	app.Logger().Error(app.Run(iris.Addr(":" + common.Port)))
}
