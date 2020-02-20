package httping

type JSendMessage struct {
	Status  JSendStatus `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
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
		return &JSendMessage{Status: StatusError}
	case statusCode >= 400 && statusCode < 500:
		return &JSendMessage{Status: StatusFail}
	default:
		return &JSendMessage{Status: StatusSuccess}
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
