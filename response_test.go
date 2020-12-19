package httping

import (
	. "github.com/onsi/gomega"
	"net/http"
	"testing"
	"time"
)

func TestNewResponse(t *testing.T) {
	RegisterTestingT(t)
	resp := NewResponse(http.StatusOK)
	Expect(resp).ToNot(BeNil())
	Expect(resp.data).To(BeNil())
	Expect(resp.statusCode).To(BeEquivalentTo(http.StatusOK))
	Expect(resp.StatusCode()).To(BeEquivalentTo(http.StatusOK))
	Expect(len(resp.headers)).To(BeEquivalentTo(0))
}

func TestAddData(t *testing.T) {
	RegisterTestingT(t)
	resp := NewResponse(http.StatusOK).SetData("test")
	Expect(resp).ToNot(BeNil())
	Expect(resp.data).To(BeEquivalentTo("test"))
	Expect(resp.statusCode).To(BeEquivalentTo(http.StatusOK))
	Expect(len(resp.headers)).To(BeEquivalentTo(0))
}

func TestAddHeader(t *testing.T) {
	RegisterTestingT(t)
	resp := NewResponse(http.StatusOK).AddHeader("test key", "test value")
	Expect(resp).ToNot(BeNil())
	Expect(resp.data).To(BeNil())
	Expect(resp.statusCode).To(BeEquivalentTo(http.StatusOK))
	Expect(len(resp.headers)).To(BeEquivalentTo(1))
	Expect(resp.headers["test key"][0]).To(BeEquivalentTo("test value"))
	headers := resp.Headers()
	Expect(len(headers)).To(BeEquivalentTo(1))
	Expect(headers["test key"][0]).To(BeEquivalentTo("test value"))
}

func TestAddCookie(t *testing.T) {
	RegisterTestingT(t)
	cookie := &http.Cookie{
		Name:    "cookie",
		Value:   "value",
		Path:    "/path",
		Domain:  "domain",
		Expires: time.Now().UTC().Add(time.Minute),
		MaxAge:  600,
	}
	resp := OK(nil).AddCookie(cookie)
	Expect(resp).ToNot(BeNil())
	Expect(resp.data).To(BeNil())
	Expect(resp.StatusCode()).To(BeEquivalentTo(http.StatusOK))
	Expect(len(resp.Headers())).To(BeEquivalentTo(0))
	Expect(len(resp.Cookies())).To(BeEquivalentTo(1))
	Expect(resp.Cookies()[0]).To(BeEquivalentTo(cookie))
}

func TestSetCookies(t *testing.T) {
	RegisterTestingT(t)
	var cookies []*http.Cookie
	cookie := &http.Cookie{
		Name:    "cookie",
		Value:   "value",
		Path:    "/path",
		Domain:  "domain",
		Expires: time.Now().UTC().Add(time.Minute),
		MaxAge:  600,
	}
	cookies = append(cookies, cookie)
	cookies = append(cookies, cookie)
	resp := OK(nil).SetCookies(cookies)
	Expect(resp).ToNot(BeNil())
	Expect(resp.data).To(BeNil())
	Expect(resp.StatusCode()).To(BeEquivalentTo(http.StatusOK))
	Expect(len(resp.Headers())).To(BeEquivalentTo(0))
	Expect(len(resp.Cookies())).To(BeEquivalentTo(2))
	Expect(resp.Cookies()[0]).To(BeEquivalentTo(cookie))
	Expect(resp.Cookies()[1]).To(BeEquivalentTo(cookie))
}

func TestOK(t *testing.T) {
	RegisterTestingT(t)
	resp := OK("test")
	checkResponse(resp, http.StatusOK)
}

func TestCreated(t *testing.T) {
	RegisterTestingT(t)
	resp := Created("test")
	checkResponse(resp, http.StatusCreated)
}

func TestAccepted(t *testing.T) {
	RegisterTestingT(t)
	resp := Accepted("test")
	checkResponse(resp, http.StatusAccepted)
}

func TestNonAuthoritativeInfo(t *testing.T) {
	RegisterTestingT(t)
	resp := NonAuthoritativeInfo("test")
	checkResponse(resp, http.StatusNonAuthoritativeInfo)
}
func TestResetContent(t *testing.T) {
	RegisterTestingT(t)
	resp := ResetContent("test")
	checkResponse(resp, http.StatusResetContent)
}
func TestPartialContent(t *testing.T) {
	RegisterTestingT(t)
	resp := PartialContent("test")
	checkResponse(resp, http.StatusPartialContent)
}
func TestMultiStatus(t *testing.T) {
	RegisterTestingT(t)
	resp := MultiStatus("test")
	checkResponse(resp, http.StatusMultiStatus)
}
func TestAlreadyReported(t *testing.T) {
	RegisterTestingT(t)
	resp := AlreadyReported("test")
	checkResponse(resp, http.StatusAlreadyReported)
}
func TestIMUsed(t *testing.T) {
	RegisterTestingT(t)
	resp := IMUsed("test")
	checkResponse(resp, http.StatusIMUsed)
}

func TestBadRequest(t *testing.T) {
	RegisterTestingT(t)
	resp := BadRequest("test")
	checkResponse(resp, http.StatusBadRequest)
}

func TestUnauthorized(t *testing.T) {
	RegisterTestingT(t)
	resp := Unauthorized("test")
	checkResponse(resp, http.StatusUnauthorized)
}

func TestForbidden(t *testing.T) {
	RegisterTestingT(t)
	resp := Forbidden("test")
	checkResponse(resp, http.StatusForbidden)
}

func TestNotFound(t *testing.T) {
	RegisterTestingT(t)
	resp := NotFound("test")
	checkResponse(resp, http.StatusNotFound)
}

func TestMethodNotAllowed(t *testing.T) {
	RegisterTestingT(t)
	resp := MethodNotAllowed("test")
	checkResponse(resp, http.StatusMethodNotAllowed)
}

func TestNotAcceptable(t *testing.T) {
	RegisterTestingT(t)
	resp := NotAcceptable("test")
	checkResponse(resp, http.StatusNotAcceptable)
}

func TestProxyAuthRequired(t *testing.T) {
	RegisterTestingT(t)
	resp := ProxyAuthRequired("test")
	checkResponse(resp, http.StatusProxyAuthRequired)
}
func TestRequestTimeout(t *testing.T) {
	RegisterTestingT(t)
	resp := RequestTimeout("test")
	checkResponse(resp, http.StatusRequestTimeout)
}
func TestConflict(t *testing.T) {
	RegisterTestingT(t)
	resp := Conflict("test")
	checkResponse(resp, http.StatusConflict)
}
func TestGone(t *testing.T) {
	RegisterTestingT(t)
	resp := Gone("test")
	checkResponse(resp, http.StatusGone)
}
func TestLengthRequired(t *testing.T) {
	RegisterTestingT(t)
	resp := LengthRequired("test")
	checkResponse(resp, http.StatusLengthRequired)
}
func TestPreconditionFailed(t *testing.T) {
	RegisterTestingT(t)
	resp := PreconditionFailed("test")
	checkResponse(resp, http.StatusPreconditionFailed)
}
func TestRequestEntityTooLarge(t *testing.T) {
	RegisterTestingT(t)
	resp := RequestEntityTooLarge("test")
	checkResponse(resp, http.StatusRequestEntityTooLarge)
}
func TestRequestURITooLong(t *testing.T) {
	RegisterTestingT(t)
	resp := RequestURITooLong("test")
	checkResponse(resp, http.StatusRequestURITooLong)
}
func TestUnsupportedMediaType(t *testing.T) {
	RegisterTestingT(t)
	resp := UnsupportedMediaType("test")
	checkResponse(resp, http.StatusUnsupportedMediaType)
}
func TestRequestedRangeNotSatisfiable(t *testing.T) {
	RegisterTestingT(t)
	resp := RequestedRangeNotSatisfiable("test")
	checkResponse(resp, http.StatusRequestedRangeNotSatisfiable)
}
func TestExpectationFailed(t *testing.T) {
	RegisterTestingT(t)
	resp := ExpectationFailed("test")
	checkResponse(resp, http.StatusExpectationFailed)
}
func TestTeapot(t *testing.T) {
	RegisterTestingT(t)
	resp := Teapot("test")
	checkResponse(resp, http.StatusTeapot)
}
func TestMisdirectedRequest(t *testing.T) {
	RegisterTestingT(t)
	resp := MisdirectedRequest("test")
	checkResponse(resp, http.StatusMisdirectedRequest)
}
func TestUnprocessableEntity(t *testing.T) {
	RegisterTestingT(t)
	resp := UnprocessableEntity("test")
	checkResponse(resp, http.StatusUnprocessableEntity)
}
func TestLocked(t *testing.T) {
	RegisterTestingT(t)
	resp := Locked("test")
	checkResponse(resp, http.StatusLocked)
}
func TestFailedDependency(t *testing.T) {
	RegisterTestingT(t)
	resp := FailedDependency("test")
	checkResponse(resp, http.StatusFailedDependency)
}
func TestTooEarly(t *testing.T) {
	RegisterTestingT(t)
	resp := TooEarly("test")
	checkResponse(resp, http.StatusTooEarly)
}
func TestUpgradeRequired(t *testing.T) {
	RegisterTestingT(t)
	resp := UpgradeRequired("test")
	checkResponse(resp, http.StatusUpgradeRequired)
}
func TestPreconditionRequired(t *testing.T) {
	RegisterTestingT(t)
	resp := PreconditionRequired("test")
	checkResponse(resp, http.StatusPreconditionRequired)
}
func TestTooManyRequests(t *testing.T) {
	RegisterTestingT(t)
	resp := TooManyRequests("test")
	checkResponse(resp, http.StatusTooManyRequests)
}
func TestRequestHeaderFieldsTooLarge(t *testing.T) {
	RegisterTestingT(t)
	resp := RequestHeaderFieldsTooLarge("test")
	checkResponse(resp, http.StatusRequestHeaderFieldsTooLarge)
}
func TestUnavailableForLegalReasons(t *testing.T) {
	RegisterTestingT(t)
	resp := UnavailableForLegalReasons("test")
	checkResponse(resp, http.StatusUnavailableForLegalReasons)
}

func TestNoContent(t *testing.T) {
	RegisterTestingT(t)
	resp := NoContent()
	Expect(resp.statusCode).To(BeEquivalentTo(http.StatusNoContent))
	Expect(resp.data).To(BeNil())
}

func TestInternalServerError(t *testing.T) {
	RegisterTestingT(t)
	resp := InternalServerError("test error message")
	Expect(resp).ToNot(BeNil())
	Expect(resp.data).To(BeEquivalentTo("test error message"))
	Expect(resp.statusCode).To(BeEquivalentTo(http.StatusInternalServerError))
	Expect(len(resp.headers)).To(BeEquivalentTo(0))
}

func checkResponse(resp *Response, status int) {
	Expect(resp).ToNot(BeNil())
	Expect(resp.data).To(BeEquivalentTo("test"))
	Expect(resp.Response()).To(BeEquivalentTo("test"))
	Expect(resp.statusCode).To(BeEquivalentTo(status))
	Expect(len(resp.headers)).To(BeEquivalentTo(0))
}
