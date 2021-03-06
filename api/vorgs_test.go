package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetVorgs(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"data":[{"id":"735-2","type":"vOrg","attributes":{"name":"Bobby Demo VDC"}},{"id":"735-11","type":"vOrg","attributes":{"name":"SAN HANA"}}]}`)
	}))

	var vorgs VorgsArray
	var err error

	papi := new(API)
	papi.SetVorgsURL(ts.URL)
	vorgs, err = papi.GetVorgs(1)

	if err != nil || len(vorgs.Data) <= 1 {
		t.Errorf("Should list 2 Vorgs")
		return
	}

	assertEqual(t, vorgs.Data[0].ID, "735-2", "")
	assertEqual(t, vorgs.Data[0].Attributes.Name, "Bobby Demo VDC", "")

	assertEqual(t, vorgs.Data[1].ID, "735-11", "")
	assertEqual(t, vorgs.Data[1].Attributes.Name, "SAN HANA", "")
}

func TestGetVorgsInvalidPerms(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"error":"No permitted account found","detail":"We couldn't find an account with ID 736 that is accessible to the authenticated user"}`)
	}))

	var vorgs VorgsArray
	var err error

	papi := new(API)
	papi.SetVorgsURL(ts.URL)
	vorgs, err = papi.GetVorgs(1)

	if err == nil {
		t.Errorf("This should throw error!")
	} else {
		assertEqual(t, err.Error(), "No permitted account found - We couldn't find an account with ID 736 that is accessible to the authenticated user", "")
	}

	assertEqual(t, vorgs.Error, "No permitted account found", "")
	assertEqual(t, vorgs.Detail, "We couldn't find an account with ID 736 that is accessible to the authenticated user", "")
}
