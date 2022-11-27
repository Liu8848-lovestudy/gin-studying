package dto

type Result struct {
	Status  int
	Message string
	Data    interface{}
}

func NewResult(status int, message string, data interface{}) Result {
	return Result{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
