package httping

type IRoute interface {
	AddMethod(method string, handler HandlerFunc) error
}
