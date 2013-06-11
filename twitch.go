// Package twitch provides an API wrapper around the Twitch APIs.
//
// https://github.com/justintv/Twitch-API/
package twitch

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	API_ENDPOINT = "api.twitch.tv/kraken"             // Twitch.tv API endpoint
	API_VERISON  = "application/vnd.twitchtv.v2+json" // The accept header used to specify the API version
)

type Twitch struct {
	httpClient *http.Client
	clientId   string
}

func New(clientId string) *Twitch {
	return &Twitch{httpClient: &http.Client{}, clientId: clientId}
}

func (client *Twitch) api(method string, path string) (body []byte, err error) {
	var req *http.Request
	url := fmt.Sprintf("https://%s/%s", API_ENDPOINT, path)

	req, err = http.NewRequest(method, url, nil)
	if err != nil {
		return
	}

	rsp, err := client.httpClient.Do(req)
	if err != nil {
		return
	}
	defer rsp.Body.Close()

	body, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		return
	}

	return
}
