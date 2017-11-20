package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCredentials(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"Bobby Demo VDC (467-735-2-a736cd)":{"service_id":"467-735-2-a736cd","username":"6604.735.cbcdc8@467-735-2-a736cd"},"SAN HANA (467-735-11-3e6af1)":{"service_id":"467-735-11-3e6af1","username":"6604.735.cbcdc8@467-735-11-3e6af1"}}`)
	}))

	var err error

	papi := new(API)
	papi.SetApiCredsURL(ts.URL)
	c, err := papi.GetCredentials()

	if err != nil || len(c.Creds) < 2 {
		t.Errorf("Should list at least 2 api creds")
	}

	assertEqual(t, c.Creds["Bobby Demo VDC (467-735-2-a736cd)"].ServiceId, "467-735-2-a736cd", "")
}
