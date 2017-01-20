package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func TestGetAccounts(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `[{"name":"UKCloud Bobby Demo","id":735}]`)
	}))

	var accounts []Accounts
	var err error

	papi := new(API)
	papi.SetAccountsURL(ts.URL)
	accounts, err = papi.GetAccounts()

	if err != nil || len(accounts) <= 0 {
		t.Errorf("Should list at least 1 account")
	}

	assertEqual(t, accounts[0].Name, "UKCloud Bobby Demo", "")
}
