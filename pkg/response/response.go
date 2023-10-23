package response

type ApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(code int, message string, data interface{}) ApiResponse {
	response := ApiResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}

	return response
}
