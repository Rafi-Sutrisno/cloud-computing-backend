package utils

type Response2 struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    uint   `json:"code"`
	Data    any    `json:"data"`
	Role    string `json:"role"`
}

func BuildResponse2(message string, statusCode uint, data any, role string) Response2 {
	res := Response2{
		Success: true,
		Message: message,
		Code:    statusCode,
		Data:    data,
		Role:    role,
	}
	return res
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    uint   `json:"code"`
	Data    any    `json:"data"`
}

func BuildResponse(message string, statusCode uint, data any) Response {
	res := Response{
		Success: true,
		Message: message,
		Code:    statusCode,
		Data:    data,
	}
	return res
}

func BuildErrorResponse(message string, statusCode uint) Response {
	res := Response{
		Success: false,
		Message: message,
		Code:    statusCode,
		Data:    nil,
	}
	return res
}
