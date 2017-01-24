package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConvertStructToJSON(t *testing.T) {
	att := VdcAttributes{
		Name:   "BobbyTest",
		VMType: "POWER",
	}

	vdc := Vdc{
		Type:       "VDC",
		Attributes: att,
	}

	vdca := VdcArray{
		Data: []Vdc{vdc},
	}

	json, err := convertStructToJSON(vdca)

	assertEqual(t, err, nil, "")

	assertEqual(t, string(json), `{"data":[{"type":"VDC","id":"","attributes":{"name":"BobbyTest","vmType":"POWER"}}],"error":"","detail":""}`, "")
}

func TestPostJSON(t *testing.T) {
	//{"data":{"type":"VDC-build","id":"3557","attributes":{"createdAt":"2017-01-24T13:46:47+00:00","createdBy":"bdeveaux@ukcloud.com","state":"approved"}}}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Location", "/api/testing")
		w.WriteHeader(http.StatusAccepted)
	}))

	att := VdcAttributes{
		Name:   "BobbyTest",
		VMType: "POWER",
	}

	vdc := Vdc{
		Type:       "VDC",
		Attributes: att,
	}

	vdca := VdcArray{
		Data: []Vdc{vdc},
	}

	response, headers, err2 := postJSON(ts.URL, vdca)
	assertEqual(t, err2, nil, "")
	assertEqual(t, headers.Get("Location"), "/api/testing", "")
	assertEqual(t, response, "", "")

}
