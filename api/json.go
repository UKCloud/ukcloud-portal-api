package api

import (
	"encoding/json"
	"io/ioutil"
	//"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

// CookiesCollection is a collection of the cookies returned from Auth
type CookiesCollection struct {
	Collection []*http.Cookie
}

var cookieCollection = new(CookiesCollection)

// getJSON calls the API and returns the JSON
func getJSON(myURL string, target interface{}) error {

	jar, _ := cookiejar.New(nil)

	u, err := url.Parse(myURL)
	jar.SetCookies(u, cookieCollection.Collection)

	tr := &http.Transport{
		DisableCompression: true,
	}

	myClient := http.Client{Jar: jar, Timeout: 100 * time.Second, Transport: tr}

	r, err := myClient.Get(myURL)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err := json.Unmarshal(body, target); err != nil {
		return err
	}

	return json.Unmarshal(body, target)
}
