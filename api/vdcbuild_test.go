package api

import (
	"fmt"
	//"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetVdcBuild(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"data":{"type":"VDC-build","id":"3557","attributes":{"createdAt":"2017-01-24T13:46:47+00:00","createdBy":"bdeveaux@ukcloud.com","state":"approved"}}}`)
	}))

	var vdcb VdcBuildData
	var err error

	papi := new(API)
	papi.SetVdcBuildURL(ts.URL)
	vdcb, err = papi.GetVdcBuild(1)

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	assertEqual(t, vdcb.Data.Type, "VDC-build", "")
	assertEqual(t, vdcb.Data.ID, "3557", "")
	assertEqual(t, vdcb.Data.Attributes.CreatedAt, "2017-01-24T13:46:47+00:00", "")
	assertEqual(t, vdcb.Data.Attributes.CreatedBy, "bdeveaux@ukcloud.com", "")
	assertEqual(t, vdcb.Data.Attributes.State, "approved", "")
}
