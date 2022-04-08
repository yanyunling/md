package entity

type Book struct {
	Id         string `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	CreateTime int64  `json:"createTime" db:"create_time"`
	UserId     string `json:"UserId" db:"user_id"`
}
