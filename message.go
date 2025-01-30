package models

type Message struct {
	UserId   string `json:"user_id"`
	Function string `json:"function"`

	Data []byte `json:"data"`

	successed bool `json:"successed"`
}

func (m *Message) IsSuccessed() bool {
	return !m.successed
}

func SuccessedMessage(userID string, function string, data []byte) *Message {
	return &Message{
		UserId:    userID,
		Function:  function,
		Data:      data,
		successed: true,
	}
}

func FailedMessage(userID string, function string, data []byte) *Message {
	return &Message{
		UserId:    userID,
		Function:  function,
		Data:      data,
		successed: false,
	}
}
