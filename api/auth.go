package api

import (
	"bytes"
	"fmt"
	//"log"
	"net/http"
	"net/http/cookiejar"
	"time"
)

// AuthURL is the Authenication endpoint
var AuthURL = "https://portal.skyscapecloud.com/api/authenticate.json"

// GetAuth gets Authorisation cookies
func (a *API) GetAuth(email string, password string) (bool, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return false, fmt.Errorf("Error setting cookies")
	}

	if len(email) < 1 || len(password) < 1 {
		return false, fmt.Errorf("Email (%s) and password (%s) must be supplied", email, password)
	}

	myClient := http.Client{Jar: jar, Timeout: 100 * time.Second}

	var jsonStr = []byte(`{"email": "` + email + ` ", "password": "` + password + `"}`)

	req, err := http.NewRequest("POST", AuthURL, bytes.NewBuffer(jsonStr))

	if err != nil {
		return false, fmt.Errorf("Error creating new request")
	}

	r, err := myClient.Do(req)

	if err != nil {
		return false, fmt.Errorf("Error posting to AuthURL: (%s)", AuthURL)
	}

	if r.StatusCode != 201 {
		return false, fmt.Errorf("Unauthorised")
	}

	cookieCollection.Collection = r.Cookies()
	return true, nil
}

// SetAuthURL for testing purposes
func (a *API) SetAuthURL(URL string) {
	AuthURL = URL
}

// GetCookieCollection for testing purposes
func (a *API) GetCookieCollection() CookiesCollection {
	return *cookieCollection
}
