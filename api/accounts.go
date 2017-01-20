package api

import ()

// Accounts holds the JSON response
type Accounts struct {
	Name string
	ID   int
}

// AccountsURL is the accounts endpoint
var AccountsURL = "https://portal.skyscapecloud.com/api/accounts.json"

// GetAccounts gets the accounts from the API
func (a *API) GetAccounts() ([]Accounts, error) {

	var accounts []Accounts
	var err error
	if err := getJSON(AccountsURL, &accounts); err != nil {
		return accounts, err
	}

	return accounts, err
}

// SetAccountsURL for testing purposes
func (a *API) SetAccountsURL(URL string) {
	AccountsURL = URL
}
