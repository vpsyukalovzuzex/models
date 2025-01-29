package models

type Message struct {
	UserId   string `json:"user_id"`
	Function string `json:"function"`

	Error Error `json:"error"`

	Data []byte `json:"data"`
}

type Error struct {
	Code int `json:"code"`

	Message string `json:"message"`
}
