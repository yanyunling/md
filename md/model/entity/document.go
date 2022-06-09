package entity

type Document struct {
	Id         string `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Content    string `json:"content" db:"content"`
	CreateTime int64  `json:"createTime" db:"create_time"`
	UpdateTime int64  `json:"updateTime" db:"update_time"`
	BookId     string `json:"bookId" db:"book_id"`
	UserId     string `json:"userId" db:"user_id"`
}
