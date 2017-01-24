package api

import (
	"fmt"
	"strconv"
	"strings"
)

// VdcArray holds the json response for multiple VDCS
type VdcArray struct {
	Data   []Vdc  `json:"data"`
	Error  string `json:"error"`
	Detail string `json:"detail"`
}

// VdcSingle holds the json request to create a single VDC
type VdcSingle struct {
	Data Vdc `json:"data"`
}

// Vdc holds each vdc of the VdcArray
type Vdc struct {
	Type       string        `json:"type"`
	ID         string        `json:"id"`
	Attributes VdcAttributes `json:"attributes"`
}

// VdcAttributes holds the name of the VDC
type VdcAttributes struct {
	Name   string `json:"name"`
	VMType string `json:"vmType"`
}

// VdcURL is the accounts endpoint
var VdcURL = "https://portal.skyscapecloud.com/api/accounts/%s/vorgs/%s/vdcs"

// GetVdc gets the accounts from the API
func (a *API) GetVdc(AccountID int, VorgID int) (VdcArray, error) {

	var vdcs VdcArray
	var err error

	n := strings.Count(VdcURL, "%s")
	if n > 0 {
		VdcURL = fmt.Sprintf(VdcURL, strconv.Itoa(AccountID), strconv.Itoa(VorgID))
	}

	if err := getJSON(VdcURL, &vdcs); err != nil {
		return vdcs, fmt.Errorf("Sorry, there was an error retrieving the VDCs")
	}

	if len(vdcs.Error) > 0 {
		return vdcs, fmt.Errorf(vdcs.Error + " - " + vdcs.Detail)
	}

	return vdcs, err
}

// CreateVdc is for creating a VDC
func (a *API) CreateVdc(AccountID int, VorgID int, name string) (string, error) {

	n := strings.Count(VdcURL, "%s")
	if n > 0 {
		VdcURL = fmt.Sprintf(VdcURL, strconv.Itoa(AccountID), strconv.Itoa(VorgID))
	}

	att := VdcAttributes{
		Name:   name,
		VMType: "POWER",
	}

	vdc := Vdc{
		Type:       "VDC",
		Attributes: att,
	}

	vdca := VdcSingle{
		Data: vdc,
	}

	response, headers, err := postJSON(VdcURL, vdca)
	if err != nil {
		return response, fmt.Errorf("Sorry, there was an error creating the VDC")
	}

	loc := headers.Get("Location")

	return loc, err
	///api/accounts/:account_id/vorgs/:vorg_id/vdcs
	//{"data": {"type": "VDC", "attributes": {"vmType": "POWER", "name": "DEMO"}}}
}

// SetVdcURL for testing purposes
func (a *API) SetVdcURL(URL string) {
	VdcURL = URL
}

// GetVdcURL for testing purposes
func (a *API) GetVdcURL() string {
	return VdcURL
}
