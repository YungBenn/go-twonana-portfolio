package response

type Message struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewMessage(code int, message string) Message {
	response := Message{
		Code:    code,
		Message: message,
	}

	return response
}
