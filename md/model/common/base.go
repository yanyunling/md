package common

type Page struct {
	Records interface{} `json:"records"`
	Total   int         `json:"total"`
}

type ManagerLogin struct {
	Email string `json:"email"`
	Vcode string `json:"vcode"`
	Token string `json:"token"`
}
