package models

type Message struct {
	UserId   string `json:"user_id"`
	Function string `json:"function"`

	Data []byte `json:"data"`
}
