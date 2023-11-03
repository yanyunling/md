package middleware

import (
	"io"
	"md/util"
	"os"
	"path"
	"time"

	"github.com/kataras/golog"
)

var (
	lastTime string
	lastFile *os.File
	Log      *golog.Logger
)

// 初始化日志
func InitLog(prefixPath string, logger *golog.Logger) {
	Log = logger
	Log.SetLevel("info")
	Log.SetTimeFormat("2006-01-02 15:04:05")

	// 首次执行
	if lastTime == "" {
		// 创建目录
		prefixPath = util.PathCompletion(prefixPath)
		err := os.MkdirAll(prefixPath, 0777)
		if err != nil {
			panic(err)
		}
		// 生成日志文件
		lastTime = time.Now().Format("20060102")
		currentFile, err := os.OpenFile(prefixPath+lastTime+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		lastFile = currentFile
		Log.SetOutput(io.MultiWriter(lastFile, os.Stdout))
	}

	// 定时扫描日志文件是否需要生成
	logTicker := time.NewTicker(60 * time.Second)
	go func(ticker *time.Ticker) {
		for {
			<-ticker.C
			currentTime := time.Now().Format("20060102")
			// 时间不一致，则生成新的日志文件
			if lastTime != currentTime {
				currentFile, err := os.OpenFile(prefixPath+currentTime+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
				if err != nil {
					Log.Error("日志文件创建失败：", err)
					continue
				}
				// 关闭上一个文件
				err = lastFile.Close()
				if err != nil {
					Log.Error("日志文件关闭失败：", err)
				}
				lastTime = currentTime
				lastFile = currentFile
				// 设置新日志文件
				Log.SetOutput(io.MultiWriter(lastFile, os.Stdout))
				// 删除超过30天的日志文件
				removeOvertimeFile(prefixPath, 30)
			}
		}
	}(logTicker)
}

// 删除早于指定天数的文件
func removeOvertimeFile(dirPath string, days int64) {
	files, err := os.ReadDir(dirPath)
	if err == nil {
		for _, f := range files {
			// 删除早于指定时间的日志文件
			if !f.IsDir() && path.Ext(f.Name()) == ".log" {
				info, err := f.Info()
				if err == nil && info.ModTime().UnixNano() < time.Now().UnixNano()-days*int64(time.Hour)*24 {
					_ = os.Remove(dirPath + f.Name())
				}
			}
		}
	}
}
