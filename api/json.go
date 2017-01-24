package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
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

func postJSON(myURL string, data interface{}) (string, http.Header, error) {
	var jsonStr, err = convertStructToJSON(data)

	if err != nil {
		//return string(jsonStr),, err
	}

	jar, _ := cookiejar.New(nil)

	u, err := url.Parse(myURL)
	jar.SetCookies(u, cookieCollection.Collection)

	tr := &http.Transport{
		DisableCompression: true,
	}

	fmt.Println("COOKIES " + strconv.Itoa(len(cookieCollection.Collection)))
	fmt.Println("JSON " + string(jsonStr))

	req, err := http.NewRequest("POST", myURL, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	myClient := http.Client{Jar: jar, Timeout: 10 * time.Second, Transport: tr}

	r, err := myClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()

	fmt.Println("response Status:", r.Status)
	fmt.Println("response Headers:", r.Header)
	body, _ := ioutil.ReadAll(r.Body)

	fmt.Println("response Body:", string(body))

	for k, v := range r.Header {
		log.Println("key:", k, "value:", v)
	}

	return string(body), r.Header, err

}

func convertStructToJSON(data interface{}) ([]byte, error) {
	var jsonStr, err = json.Marshal(data)

	if err != nil {
		return jsonStr, err
	}

	return jsonStr, err
}
