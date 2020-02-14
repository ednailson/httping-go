package httping

type JSendMessage struct {
	Status  JSendStatus `json:"status"`
	Data    string      `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

type JSendStatus string

const (
	StatusSuccess JSendStatus = "success"
	StatusError   JSendStatus = "error"
	StatusFail    JSendStatus = "fail"
)

func NewJSendMessage(statusCode int) *JSendMessage {
	statusString := StatusSuccess
	switch {
	case statusCode >= 500:
		statusString = StatusError
	case statusCode >= 400 && statusCode < 500:
		statusString = StatusFail
	}
	return &JSendMessage{Status: statusString}
}
