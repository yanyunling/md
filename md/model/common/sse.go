package common

type SSEMessage struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}
