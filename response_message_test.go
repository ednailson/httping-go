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
	Expect(resp.Status).To(BeEquivalentTo(StatusSuccess))
	Expect(resp.Code).To(BeEquivalentTo(""))
	Expect(resp.Data).To(BeNil())
	Expect(resp.Message).To(BeEquivalentTo(""))
	Expect(resp.statusCode).To(BeEquivalentTo(http.StatusOK))
	Expect(resp.StatusCode()).To(BeEquivalentTo(http.StatusOK))
	Expect(len(resp.headers)).To(BeEquivalentTo(0))
	resp.AddMessage("test message")
	Expect(resp.Message).To(BeEquivalentTo(""))
	resp.AddCode("test code")
	Expect(resp.Code).To(BeEquivalentTo(""))
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
	Expect(resp.StatusCode()).To(BeEquivalentTo(http.StatusInternalServerError))
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
	resp.AddMessage("test message")
	Expect(resp.Message).To(BeEquivalentTo(""))
	resp.AddCode("test code")
	Expect(resp.Code).To(BeEquivalentTo(""))
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
	Expect(resp.Status).To(BeEquivalentTo(StatusSuccess))
	Expect(resp.Code).To(BeEquivalentTo(""))
	Expect(resp.Data).To(BeNil())
	Expect(resp.Message).To(BeEquivalentTo(""))
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
	Expect(resp.Status).To(BeEquivalentTo(StatusSuccess))
	Expect(resp.Code).To(BeEquivalentTo(""))
	Expect(resp.Data).To(BeNil())
	Expect(resp.Message).To(BeEquivalentTo(""))
	Expect(resp.StatusCode()).To(BeEquivalentTo(http.StatusOK))
	Expect(len(resp.Headers())).To(BeEquivalentTo(0))
	Expect(len(resp.Cookies())).To(BeEquivalentTo(2))
	Expect(resp.Cookies()[0]).To(BeEquivalentTo(cookie))
	Expect(resp.Cookies()[1]).To(BeEquivalentTo(cookie))
}
