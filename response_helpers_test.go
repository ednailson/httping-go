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

func TestNonAuthoritativeInfo(t *testing.T) {
	RegisterTestingT(t)
	resp := NonAuthoritativeInfo("test")
	checkResponseSuccess(resp, http.StatusNonAuthoritativeInfo)
}
func TestResetContent(t *testing.T) {
	RegisterTestingT(t)
	resp := ResetContent("test")
	checkResponseSuccess(resp, http.StatusResetContent)
}
func TestPartialContent(t *testing.T) {
	RegisterTestingT(t)
	resp := PartialContent("test")
	checkResponseSuccess(resp, http.StatusPartialContent)
}
func TestMultiStatus(t *testing.T) {
	RegisterTestingT(t)
	resp := MultiStatus("test")
	checkResponseSuccess(resp, http.StatusMultiStatus)
}
func TestAlreadyReported(t *testing.T) {
	RegisterTestingT(t)
	resp := AlreadyReported("test")
	checkResponseSuccess(resp, http.StatusAlreadyReported)
}
func TestIMUsed(t *testing.T) {
	RegisterTestingT(t)
	resp := IMUsed("test")
	checkResponseSuccess(resp, http.StatusIMUsed)
}

func TestNoContent(t *testing.T) {
	RegisterTestingT(t)
	resp := NoContent()
	Expect(resp).ToNot(BeNil())
	Expect(resp.Status).To(BeEquivalentTo(""))
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

func TestProxyAuthRequired(t *testing.T) {
	RegisterTestingT(t)
	resp := ProxyAuthRequired("test")
	checkResponseFail(resp, http.StatusProxyAuthRequired)
}
func TestRequestTimeout(t *testing.T) {
	RegisterTestingT(t)
	resp := RequestTimeout("test")
	checkResponseFail(resp, http.StatusRequestTimeout)
}
func TestConflict(t *testing.T) {
	RegisterTestingT(t)
	resp := Conflict("test")
	checkResponseFail(resp, http.StatusConflict)
}
func TestGone(t *testing.T) {
	RegisterTestingT(t)
	resp := Gone("test")
	checkResponseFail(resp, http.StatusGone)
}
func TestLengthRequired(t *testing.T) {
	RegisterTestingT(t)
	resp := LengthRequired("test")
	checkResponseFail(resp, http.StatusLengthRequired)
}
func TestPreconditionFailed(t *testing.T) {
	RegisterTestingT(t)
	resp := PreconditionFailed("test")
	checkResponseFail(resp, http.StatusPreconditionFailed)
}
func TestRequestEntityTooLarge(t *testing.T) {
	RegisterTestingT(t)
	resp := RequestEntityTooLarge("test")
	checkResponseFail(resp, http.StatusRequestEntityTooLarge)
}
func TestRequestURITooLong(t *testing.T) {
	RegisterTestingT(t)
	resp := RequestURITooLong("test")
	checkResponseFail(resp, http.StatusRequestURITooLong)
}
func TestUnsupportedMediaType(t *testing.T) {
	RegisterTestingT(t)
	resp := UnsupportedMediaType("test")
	checkResponseFail(resp, http.StatusUnsupportedMediaType)
}
func TestRequestedRangeNotSatisfiable(t *testing.T) {
	RegisterTestingT(t)
	resp := RequestedRangeNotSatisfiable("test")
	checkResponseFail(resp, http.StatusRequestedRangeNotSatisfiable)
}
func TestExpectationFailed(t *testing.T) {
	RegisterTestingT(t)
	resp := ExpectationFailed("test")
	checkResponseFail(resp, http.StatusExpectationFailed)
}
func TestTeapot(t *testing.T) {
	RegisterTestingT(t)
	resp := Teapot("test")
	checkResponseFail(resp, http.StatusTeapot)
}
func TestMisdirectedRequest(t *testing.T) {
	RegisterTestingT(t)
	resp := MisdirectedRequest("test")
	checkResponseFail(resp, http.StatusMisdirectedRequest)
}
func TestUnprocessableEntity(t *testing.T) {
	RegisterTestingT(t)
	resp := UnprocessableEntity("test")
	checkResponseFail(resp, http.StatusUnprocessableEntity)
}
func TestLocked(t *testing.T) {
	RegisterTestingT(t)
	resp := Locked("test")
	checkResponseFail(resp, http.StatusLocked)
}
func TestFailedDependency(t *testing.T) {
	RegisterTestingT(t)
	resp := FailedDependency("test")
	checkResponseFail(resp, http.StatusFailedDependency)
}
func TestTooEarly(t *testing.T) {
	RegisterTestingT(t)
	resp := TooEarly("test")
	checkResponseFail(resp, http.StatusTooEarly)
}
func TestUpgradeRequired(t *testing.T) {
	RegisterTestingT(t)
	resp := UpgradeRequired("test")
	checkResponseFail(resp, http.StatusUpgradeRequired)
}
func TestPreconditionRequired(t *testing.T) {
	RegisterTestingT(t)
	resp := PreconditionRequired("test")
	checkResponseFail(resp, http.StatusPreconditionRequired)
}
func TestTooManyRequests(t *testing.T) {
	RegisterTestingT(t)
	resp := TooManyRequests("test")
	checkResponseFail(resp, http.StatusTooManyRequests)
}
func TestRequestHeaderFieldsTooLarge(t *testing.T) {
	RegisterTestingT(t)
	resp := RequestHeaderFieldsTooLarge("test")
	checkResponseFail(resp, http.StatusRequestHeaderFieldsTooLarge)
}
func TestUnavailableForLegalReasons(t *testing.T) {
	RegisterTestingT(t)
	resp := UnavailableForLegalReasons("test")
	checkResponseFail(resp, http.StatusUnavailableForLegalReasons)
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
