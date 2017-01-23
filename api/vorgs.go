package api

import (
	"fmt"
	"strconv"
	"strings"
)

// VorgsArray holds the json response
type VorgsArray struct {
	Data []Vorgs `json:"data"`
}

// Vorgs holds each vdc of the VdcArray
type Vorgs struct {
	ID         string
	Type       string
	Attributes VorgsAttributes
}

// VorgsAttributes holds the name of the VDC
type VorgsAttributes struct {
	Name string
}

// VorgsURL is the accounts endpoint
var VorgsURL = "https://portal.skyscapecloud.com/api/accounts/%s/vorgs.json"

// GetVorgs gets the Vorgs from the API
func (a *API) GetVorgs(AccountID int) (VorgsArray, error) {

	var vorgs VorgsArray
	var err error

	n := strings.Count(VorgsURL, "%s")
	if n > 0 {
		VorgsURL = fmt.Sprintf(VorgsURL, strconv.Itoa(AccountID))
	}

	if err := getJSON(VorgsURL, &vorgs); err != nil {
		return vorgs, fmt.Errorf("Sorry, there was an error retrieving the Vorgs")
	}

	return vorgs, err
}

// SetVorgsURL for testing purposes
func (a *API) SetVorgsURL(URL string) {
	VorgsURL = URL
}

// GetVorgsURL for testing purposes
func (a *API) GetVorgsURL() string {
	return VorgsURL
}
