package httping

import (
	. "github.com/onsi/gomega"
	"net/http"
	"testing"
)

func TestOK(t *testing.T) {
	RegisterTestingT(t)
	resp := OK("test")
	checkResponseSuccess(resp, http.StatusOK)
}

func TestCreated(t *testing.T) {
	RegisterTestingT(t)
	resp := Created("test")
	checkResponseSuccess(resp, http.StatusCreated)
}

func TestAccepted(t *testing.T) {
	RegisterTestingT(t)
	resp := Accepted("test")
	checkResponseSuccess(resp, http.StatusAccepted)
}

func TestNoContent(t *testing.T) {
	RegisterTestingT(t)
	resp := NoContent()
	Expect(resp).ToNot(BeNil())
	Expect(resp.Status).To(BeEquivalentTo(StatusSuccess))
	Expect(resp.Code).To(BeEquivalentTo(""))
	Expect(resp.Data).To(BeNil())
	Expect(resp.Message).To(BeEquivalentTo(""))
	Expect(resp.statusCode).To(BeEquivalentTo(http.StatusNoContent))
	Expect(len(resp.headers)).To(BeEquivalentTo(0))
}

func TestBadRequest(t *testing.T) {
	RegisterTestingT(t)
	resp := BadRequest("test")
	checkResponseFail(resp, http.StatusBadRequest)
}

func TestUnauthorized(t *testing.T) {
	RegisterTestingT(t)
	resp := Unauthorized("test")
	checkResponseFail(resp, http.StatusUnauthorized)
}

func TestForbidden(t *testing.T) {
	RegisterTestingT(t)
	resp := Forbidden("test")
	checkResponseFail(resp, http.StatusForbidden)
}

func TestNotFound(t *testing.T) {
	RegisterTestingT(t)
	resp := NotFound("test")
	checkResponseFail(resp, http.StatusNotFound)
}

func TestMethodNotAllowed(t *testing.T) {
	RegisterTestingT(t)
	resp := MethodNotAllowed("test")
	checkResponseFail(resp, http.StatusMethodNotAllowed)
}

func TestNotAcceptable(t *testing.T) {
	RegisterTestingT(t)
	resp := NotAcceptable("test")
	checkResponseFail(resp, http.StatusNotAcceptable)
}

func TestInternalServerError(t *testing.T) {
	RegisterTestingT(t)
	resp := InternalServerError("test error message")
	Expect(resp).ToNot(BeNil())
	Expect(resp.Status).To(BeEquivalentTo(StatusError))
	Expect(resp.Code).To(BeEquivalentTo(""))
	Expect(resp.Data).To(BeNil())
	Expect(resp.Message).To(BeEquivalentTo("test error message"))
	Expect(resp.statusCode).To(BeEquivalentTo(http.StatusInternalServerError))
	Expect(len(resp.headers)).To(BeEquivalentTo(0))
}

func checkResponseSuccess(resp *ResponseMessage, status int) {
	Expect(resp).ToNot(BeNil())
	Expect(resp.Status).To(BeEquivalentTo(StatusSuccess))
	Expect(resp.Code).To(BeEquivalentTo(""))
	Expect(resp.Data).To(BeEquivalentTo("test"))
	Expect(resp.Message).To(BeEquivalentTo(""))
	Expect(resp.statusCode).To(BeEquivalentTo(status))
	Expect(len(resp.headers)).To(BeEquivalentTo(0))
}

func checkResponseFail(resp *ResponseMessage, status int) {
	Expect(resp).ToNot(BeNil())
	Expect(resp.Status).To(BeEquivalentTo(StatusFail))
	Expect(resp.Code).To(BeEquivalentTo(""))
	Expect(resp.Data).To(BeEquivalentTo("test"))
	Expect(resp.Message).To(BeEquivalentTo(""))
	Expect(resp.statusCode).To(BeEquivalentTo(status))
	Expect(len(resp.headers)).To(BeEquivalentTo(0))
}
