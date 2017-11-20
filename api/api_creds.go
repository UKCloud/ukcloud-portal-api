package api

import (
	"fmt"
	"strconv"
)

// APICreds holds the JSON response
type APICreds struct {
	ServiceId string `json:"service_id"`
	Username  string `json:"username"`
}

type APICredsCollection struct {
	Creds map[string]APICreds
}

// AccountsURL is the accounts endpoint
var ApiCredsURL = "https://portal.skyscapecloud.com/api/accounts/%s/api_credentials"

// GetAccounts gets the accounts from the API
func (a *API) GetCredentials(AccountId int) (*APICredsCollection, error) {

	c := new(APICredsCollection)
	var creds = &c.Creds
	var err error

	ApiCredsURL = fmt.Sprintf(ApiCredsURL, strconv.Itoa(AccountId))

	if err := getJSON(ApiCredsURL, &creds); err != nil {
		return &APICredsCollection{}, err
	}

	return c, err
}

// SetAccountsURL for testing purposes
func (a *API) SetApiCredsURL(URL string) {
	ApiCredsURL = URL
}
