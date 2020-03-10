package httping

import "net/http"

func OK(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusOK).AddData(data)
}

func Created(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusCreated).AddData(data)
}

func Accepted(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusAccepted).AddData(data)
}

func NoContent() *ResponseMessage {
	return NewResponse(http.StatusNoContent)
}

func BadRequest(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusBadRequest).AddData(data)
}

func Unauthorized(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusUnauthorized).AddData(data)
}

func Forbidden(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusForbidden).AddData(data)
}

func NotFound(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusNotFound).AddData(data)
}

func MethodNotAllowed(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusMethodNotAllowed).AddData(data)
}

func NotAcceptable(data interface{}) *ResponseMessage {
	return NewResponse(http.StatusNotAcceptable).AddData(data)
}

func InternalServerError(message string) *ResponseMessage {
	return NewResponse(http.StatusInternalServerError).AddMessage(message)
}
