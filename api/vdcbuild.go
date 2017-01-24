package api

import (
	"fmt"
	"strconv"
	"strings"
)

// VdcBuildData holds the json response
type VdcBuildData struct {
	Data   VdcBuild `json:"data"`
	Error  string   `json:"error"`
	Detail string   `json:"detail"`
}

// VdcBuild holds each vdc of the VdcArray
type VdcBuild struct {
	Type       string             `json:"type"`
	ID         string             `json:"id"`
	Attributes VdcBuildAttributes `json:"attributes"`
}

// VdcBuildAttributes holds the name of the VDC
type VdcBuildAttributes struct {
	CreatedAt string `json:"CreatedAt"`
	CreatedBy string `json:"CreatedBy"`
	State     string `json:"state"`
}

// VdcBuildURL is the accounts endpoint
var VdcBuildURL = "https://portal.skyscapecloud.com/api/vdc-builds/%s.json"

// GetVdcBuild gets the VDC Build from the API
func (a *API) GetVdcBuild(BuildID int) (VdcBuildData, error) {

	var vdcb VdcBuildData
	var err error

	n := strings.Count(VdcBuildURL, "%s")
	if n > 0 {
		VdcBuildURL = fmt.Sprintf(VdcBuildURL, strconv.Itoa(BuildID))
	}

	if err := getJSON(VdcBuildURL, &vdcb); err != nil {
		return vdcb, fmt.Errorf("Sorry, there was an error retrieving the VDC Build")
	}

	if len(vdcb.Error) > 0 {
		return vdcb, fmt.Errorf(vdcb.Error + " - " + vdcb.Detail)
	}

	return vdcb, err
}

// SetVdcBuildURL for testing purposes
func (a *API) SetVdcBuildURL(URL string) {
	VdcBuildURL = URL
}

// GetVdcBuildURL for testing purposes
func (a *API) GetVdcBuildURL() string {
	return VdcBuildURL
}
