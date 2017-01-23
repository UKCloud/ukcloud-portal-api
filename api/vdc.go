package api

import (
	"fmt"
	"strconv"
	"strings"
)

// VdcArray holds the json response
type VdcArray struct {
	Data []Vdc `json:"data"`
}

// Vdc holds each vdc of the VdcArray
type Vdc struct {
	Type       string
	ID         string
	Attributes VdcAttributes
}

// VdcAttributes holds the name of the VDC
type VdcAttributes struct {
	Name string
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
		return vdcs, fmt.Errorf("Sorry, there was an error retrieving the accounts")
	}

	return vdcs, err
}

// SetVdcURL for testing purposes
func (a *API) SetVdcURL(URL string) {
	VdcURL = URL
}

// GetVdcURL for testing purposes
func (a *API) GetVdcURL() string {
	return VdcURL
}
