package api

import (
	"bytes"
	//"log"
	"net/http"
	"net/http/cookiejar"
	"time"
)

// AuthURL is the Authenication endpoint
var AuthURL = "https://portal.skyscapecloud.com/api/authenticate.json"

// GetAuth gets Authorisation cookies
func (a *API) GetAuth(email string, password string) int {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	if len(email) < 1 || len(password) < 1 {
		return 1
	}

	myClient := http.Client{Jar: jar, Timeout: 100 * time.Second}

	var jsonStr = []byte(`{"email": "` + email + ` ", "password": "` + password + `"}`)

	req, err := http.NewRequest("POST", AuthURL, bytes.NewBuffer(jsonStr))

	if err != nil {
		return 2
	}

	r, err := myClient.Do(req)

	if err != nil {
		return 3
	}

	if r.StatusCode != 201 {
		return 4
	}

	cookieCollection.Collection = r.Cookies()
	return 0
}

// SetAuthURL for testing purposes
func (a *API) SetAuthURL(URL string) {
	AuthURL = URL
}

// GetCookieCollection for testing purposes
func (a *API) GetCookieCollection() CookiesCollection {
	return *cookieCollection
}
