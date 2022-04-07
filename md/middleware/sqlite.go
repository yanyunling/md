package middleware

import (
	"os"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// 数据库连接
var Db *sqlx.DB

// 建表语句
var createTableSql = `
CREATE TABLE IF NOT EXISTS t_user
(
	id TEXT PRIMARY KEY NOT NULL,
	name TEXT NOT NULL,
	password TEXT NOT NULL,
	create_time INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS t_document
(
	id TEXT PRIMARY KEY NOT NULL,
	name TEXT NOT NULL,
	content TEXT NOT NULL,
	create_time INTEGER NOT NULL,
	update_time INTEGER NOT NULL,
	user_id TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS t_picture
(
	id TEXT PRIMARY KEY NOT NULL,
	name TEXT NOT NULL,
	path TEXT NOT NULL,
	hash TEXT NOT NULL,
	size INTEGER NOT NULL,
	create_time INTEGER NOT NULL,
	user_id TEXT NOT NULL
);
`

// 初始化sqlite
func InitSqlite(dbPath string) error {
	// 创建目录
	if dbPath == "" {
		dbPath = "./"
	} else if !strings.HasSuffix(dbPath, "/") {
		dbPath += "/"
	}
	err := os.MkdirAll(dbPath, 0666)
	if err != nil {
		Log.Error("创建数据库文件目录失败：", err)
		return err
	}

	// 开启数据库文件
	Db, err = sqlx.Connect("sqlite3", dbPath+"md.db")
	if err != nil {
		Log.Error("开启数据库文件失败：", err)
		return err
	}

	// 创建表结构
	Db.MustExec(createTableSql)

	return nil
}
