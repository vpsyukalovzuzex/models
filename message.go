package models

type Message struct {
	UserId   string `json:"user_id"`
	Function string `json:"function"`

	Data []byte `json:"data"`

	Successed bool `json:"successed"`
}

func SuccessedMessage(userID string, function string, data []byte) *Message {
	return &Message{
		UserId:    userID,
		Function:  function,
		Data:      data,
		Successed: true,
	}
}

func FailedMessage(userID string, function string, data []byte) *Message {
	return &Message{
		UserId:    userID,
		Function:  function,
		Data:      data,
		Successed: false,
	}
}
