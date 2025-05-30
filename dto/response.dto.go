package dto

type Response[T any] struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func CreateResponseError(Message string) Response[string] {
	return Response[string]{
		Code:    "99",
		Message: Message,
		Data:    "",
	}
}
func CreateResponseErrorData(Message string, data map[string]string) Response[map[string]string] {
	return Response[map[string]string]{
		Code:    "99",
		Message: Message,
		Data:    data,
	}
}
func CreateResponseSuccess[T any](data T) Response[T] {
	return Response[T]{
		Code:    "00",
		Message: "Success",
		Data:    data,
	}
}
