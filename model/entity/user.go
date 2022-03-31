package entity

type User struct {
	Id         string `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Password   string `json:"password" db:"password"`
	Phone      string `json:"phone" db:"phone"`
	Email      string `json:"email" db:"email"`
	Portrait   string `json:"portrait" db:"portrait"`
	Remark     string `json:"remark" db:"remark"`
	Status     int    `json:"status" db:"status"`
	CreateTime int64  `json:"createTime" db:"create_time"`
	UpdateTime int64  `json:"updateTime" db:"update_time"`
}

type UserCon struct {
	User
	Page        int    `json:"page"`
	Size        int    `json:"size"`
	OldPassword string `json:"oldPassword"`
	Vcode       string `json:"vcode"`
}
