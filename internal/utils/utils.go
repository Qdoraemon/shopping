package utils

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}

func Success(Data interface{}, Message string) Response {
	return Response{
		Data:    Data,
		Message: Message,
		Code:    0,
	}

}

func Error(Code int, Message string) Response {
	return Response{
		// Data:    Data,
		Message: Message,
		Code:    Code,
	}

}
