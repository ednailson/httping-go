package httping

type ResponseMessage struct {
	Status     ResponseStatus `json:"status"`
	Data       interface{}    `json:"data,omitempty"`
	Message    string         `json:"message,omitempty"`
	headers    map[string][]string
	statusCode int
}

type ResponseStatus string

const (
	StatusSuccess ResponseStatus = "success"
	StatusError   ResponseStatus = "error"
	StatusFail    ResponseStatus = "fail"
)

func NewResponse(statusCode int) *ResponseMessage {
	switch {
	case statusCode >= 500:
		return &ResponseMessage{Status: StatusError, headers: make(map[string][]string), statusCode: statusCode}
	case statusCode >= 400 && statusCode < 500:
		return &ResponseMessage{Status: StatusFail, headers: make(map[string][]string), statusCode: statusCode}
	default:
		return &ResponseMessage{Status: StatusSuccess, headers: make(map[string][]string), statusCode: statusCode}
	}
}

func (r *ResponseMessage) AddData(data interface{}) *ResponseMessage {
	r.Data = data
	return r
}

func (r *ResponseMessage) AddMessage(message string) *ResponseMessage {
	r.Message = message
	return r
}

func (r *ResponseMessage) AddHeader(key, value string) *ResponseMessage {
	r.headers[key] = append(r.headers[key], value)
	return r
}
