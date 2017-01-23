package api

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetVdc(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"data":[{"type":"VDC","id":"urn:vcloud:vdc:b40807eb-cead-42b3-ba60-b3c8566c9873","attributes":{"name":"A-ESS-Bobby Demo VDC"}}]}`)
	}))

	var vdcs VdcArray
	var err error

	papi := new(API)
	papi.SetVdcURL(ts.URL)
	vdcs, err = papi.GetVdc(1, 1)

	if err != nil || len(vdcs.Data) <= 0 {
		t.Errorf("Should list at least 1 vdc")
		return
	}

	assertEqual(t, vdcs.Data[0].ID, "urn:vcloud:vdc:b40807eb-cead-42b3-ba60-b3c8566c9873", "")
	assertEqual(t, vdcs.Data[0].Attributes.Name, "A-ESS-Bobby Demo VDC", "")
}

func TestGetVdcInvalidPerms(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"error":"Insufficient Permission","detail":"Compute-Service permission is required in order to list VDCs."}`)
	}))

	var vdcs VdcArray
	var err error

	papi := new(API)
	papi.SetVdcURL(ts.URL)
	vdcs, err = papi.GetVdc(1, 1)

	log.Println(vdcs)
	if err == nil {
		t.Errorf("This should throw error!")
	} else {
		assertEqual(t, err.Error(), "Insufficient Permission - Compute-Service permission is required in order to list VDCs.", "")
	}

	assertEqual(t, vdcs.Error, "Insufficient Permission", "")
	assertEqual(t, vdcs.Detail, "Compute-Service permission is required in order to list VDCs.", "")
}
