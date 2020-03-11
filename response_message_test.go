package httping

import (
	. "github.com/onsi/gomega"
	"net/http"
	"testing"
)

func TestNewResponse(t *testing.T) {
	RegisterTestingT(t)
	resp := NewResponse(http.StatusOK)
	Expect(resp).ToNot(BeNil())
	Expect(resp.Status).To(BeEquivalentTo(StatusSuccess))
	Expect(resp.Code).To(BeEquivalentTo(""))
	Expect(resp.Data).To(BeNil())
	Expect(resp.Message).To(BeEquivalentTo(""))
	Expect(resp.statusCode).To(BeEquivalentTo(http.StatusOK))
	Expect(len(resp.headers)).To(BeEquivalentTo(0))
}

func TestAddMessage(t *testing.T) {
	RegisterTestingT(t)
	resp := NewResponse(http.StatusInternalServerError).AddMessage("message error")
	Expect(resp).ToNot(BeNil())
	Expect(resp.Status).To(BeEquivalentTo(StatusError))
	Expect(resp.Code).To(BeEquivalentTo(""))
	Expect(resp.Data).To(BeNil())
	Expect(resp.Message).To(BeEquivalentTo("message error"))
	Expect(resp.statusCode).To(BeEquivalentTo(http.StatusInternalServerError))
	Expect(len(resp.headers)).To(BeEquivalentTo(0))
}

func TestAddCode(t *testing.T) {
	RegisterTestingT(t)
	resp := NewResponse(http.StatusInternalServerError).AddCode("CODE-ERROR")
	Expect(resp).ToNot(BeNil())
	Expect(resp.Status).To(BeEquivalentTo(StatusError))
	Expect(resp.Code).To(BeEquivalentTo("CODE-ERROR"))
	Expect(resp.Data).To(BeNil())
	Expect(resp.Message).To(BeEquivalentTo(""))
	Expect(resp.statusCode).To(BeEquivalentTo(http.StatusInternalServerError))
	Expect(len(resp.headers)).To(BeEquivalentTo(0))
}

func TestFailMessage(t *testing.T) {
	RegisterTestingT(t)
	resp := NewResponse(http.StatusBadRequest)
	Expect(resp).ToNot(BeNil())
	Expect(resp.Status).To(BeEquivalentTo(StatusFail))
	Expect(resp.Code).To(BeEquivalentTo(""))
	Expect(resp.Data).To(BeNil())
	Expect(resp.Message).To(BeEquivalentTo(""))
	Expect(resp.statusCode).To(BeEquivalentTo(http.StatusBadRequest))
	Expect(len(resp.headers)).To(BeEquivalentTo(0))
}

func TestAddData(t *testing.T) {
	RegisterTestingT(t)
	resp := NewResponse(http.StatusOK).AddData("test")
	Expect(resp).ToNot(BeNil())
	Expect(resp.Status).To(BeEquivalentTo(StatusSuccess))
	Expect(resp.Data).To(BeEquivalentTo("test"))
	Expect(resp.Code).To(BeEquivalentTo(""))
	Expect(resp.Message).To(BeEquivalentTo(""))
	Expect(resp.statusCode).To(BeEquivalentTo(http.StatusOK))
	Expect(len(resp.headers)).To(BeEquivalentTo(0))
}

func TestAddHeader(t *testing.T) {
	RegisterTestingT(t)
	resp := NewResponse(http.StatusOK).AddHeader("test key", "test value")
	Expect(resp).ToNot(BeNil())
	Expect(resp.Status).To(BeEquivalentTo(StatusSuccess))
	Expect(resp.Code).To(BeEquivalentTo(""))
	Expect(resp.Data).To(BeNil())
	Expect(resp.Message).To(BeEquivalentTo(""))
	Expect(resp.statusCode).To(BeEquivalentTo(http.StatusOK))
	Expect(len(resp.headers)).To(BeEquivalentTo(1))
	Expect(resp.headers["test key"][0]).To(BeEquivalentTo("test value"))
}
