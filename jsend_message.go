package httping

type JSendMessage struct {
	Status  JSendStatus `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	headers map[string][]string
}

type JSendStatus string

const (
	StatusSuccess JSendStatus = "success"
	StatusError   JSendStatus = "error"
	StatusFail    JSendStatus = "fail"
)

func NewJSend(statusCode int) *JSendMessage {
	switch {
	case statusCode >= 500:
		return &JSendMessage{Status: StatusError, headers: make(map[string][]string)}
	case statusCode >= 400 && statusCode < 500:
		return &JSendMessage{Status: StatusFail, headers: make(map[string][]string)}
	default:
		return &JSendMessage{Status: StatusSuccess, headers: make(map[string][]string)}
	}
}

func (j *JSendMessage) AddData(data interface{}) *JSendMessage {
	j.Data = data
	return j
}

func (j *JSendMessage) AddMessage(message string) *JSendMessage {
	j.Message = message
	return j
}

func (j *JSendMessage) AddHeader(key, value string) *JSendMessage {
	j.headers[key] = append(j.headers[key], value)
	return j
}
