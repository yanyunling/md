package common

type Page struct {
	Records interface{} `json:"records"`
	Total   int         `json:"total"`
}

type CountResult struct {
	Count int `json:"count" db:"count"`
}

type TokenResult struct {
	Name         string `json:"name"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type TokenCache struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
