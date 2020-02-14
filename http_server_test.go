package httping

import (
	. "github.com/onsi/gomega"
	"testing"
)

const port = 5001
const baseUrl = "http://localhost:5001"

func TestNewServer(t *testing.T) {
	RegisterTestingT(t)
	server := NewHttpServer(port)
	Expect(server.engine).ShouldNot(BeNil())
	Expect(server.server).ShouldNot(BeNil())
}

func TestNewRoute(t *testing.T) {
	RegisterTestingT(t)
	server := NewHttpServer(port)
	Expect(server.engine).ShouldNot(BeNil())
	Expect(server.server).ShouldNot(BeNil())
	route := server.NewRoute(nil, "/test")
	Expect(route.route).ShouldNot(BeNil())
}
