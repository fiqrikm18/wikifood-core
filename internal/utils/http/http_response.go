package http

type SuccessResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type FailureResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewSuccessResponse(code int, message string, data interface{}) SuccessResponse {
	return SuccessResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func NewFailureResponse(code int, message string) FailureResponse {
	return FailureResponse{
		Code:    code,
		Message: message,
	}
}
