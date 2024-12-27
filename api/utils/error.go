package utils

type ErrorStruct struct {
	Code    int
	Message string
}

func NewErrorStruct(code int, message string) ErrorStruct {
	return ErrorStruct{
		Code:    code,
		Message: message,
	}
}

func (err ErrorStruct) Error() string {
	return err.Message
}
