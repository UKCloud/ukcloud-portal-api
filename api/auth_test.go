package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetAuth(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expire := time.Now().AddDate(0, 0, 1)
		cookie1 := http.Cookie{"test", "testcookie", "/", "www.skyscapecloud.com", expire, expire.Format(time.UnixDate), 86400, true, true, "test=testcookie", []string{"test=testcookie"}}
		cookie2 := http.Cookie{"test2", "testcookie2", "/", "www.skyscapecloud.com", expire, expire.Format(time.UnixDate), 86400, true, true, "test2=testcookie2", []string{"test2=testcookie2"}}
		http.SetCookie(w, &cookie1)
		http.SetCookie(w, &cookie2)
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"expire_after":900}`)
	}))

	papi := new(API)
	papi.SetAuthURL(ts.URL)
	r := papi.GetAuth("username", "password")

	assertEqual(t, r, 0, "")

	assertEqual(t, len(papi.GetCookieCollection().Collection), 2, "")
}

func TestGetAuthNoCookie(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"expire_after":900}`)
	}))

	papi := new(API)
	papi.SetAuthURL(ts.URL)
	r := papi.GetAuth("username", "password")

	assertEqual(t, r, 0, "")

	assertEqual(t, len(papi.GetCookieCollection().Collection), 0, "")
}

func TestGetAuthIncorrectStatus(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"expire_after":900}`)
	}))

	papi := new(API)
	papi.SetAuthURL(ts.URL)
	r := papi.GetAuth("username", "password")

	assertEqual(t, r, 4, "")

}
