package middleware

import (
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
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
	type TEXT NOT NULL,
	published BOOLEAN NOT NULL,
	create_time INTEGER NOT NULL,
	update_time INTEGER NOT NULL,
	book_id TEXT NOT NULL,
	user_id TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS t_book
(
	id TEXT PRIMARY KEY NOT NULL,
	name TEXT NOT NULL,
	create_time INTEGER NOT NULL,
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

CREATE INDEX IF NOT EXISTS "book_user_id"
ON "t_book" (
  "user_id" ASC
);

CREATE INDEX IF NOT EXISTS "document_user_id_book_id"
ON "t_document" (
  "user_id" ASC,
  "book_id" ASC
);

CREATE INDEX IF NOT EXISTS "picture_size_hash"
ON "t_picture" (
  "size" ASC,
  "hash" ASC
);

CREATE INDEX IF NOT EXISTS "picture_user_id"
ON "t_picture" (
  "user_id" ASC
);

CREATE UNIQUE INDEX IF NOT EXISTS "user_name"
ON "t_user" (
  "name" ASC
);
`

// 初始化sqlite
func InitSqlite(dataPath string) error {
	// 开启数据库文件
	var err error
	Db, err = sqlx.Connect("sqlite", dataPath+"md.db")
	if err != nil {
		Log.Error("开启数据库文件失败：", err)
		return err
	}

	// 创建表结构
	Db.MustExec(createTableSql)

	return nil
}
