package getvideo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// NewClient returns a getvideoClient capable of requests
func NewClient(key string) Client {
	var getvideoclient Client
	getvideoclient.apikey = key

	return getvideoclient
}

// NewRequest performs a get request
func (client Client) NewRequest(url string) (Video, error) {
	var video Video

	body, err := get(client.apikey, endpoint+url)
	if err != nil {
		return video, err
	}

	err = json.Unmarshal(body, &video)

	return video, err
}

func get(apikey string, url string) (body []byte, gerr error) {

	var client http.Client

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return body, err
	}

	req.Header.Add("X-RapidAPI-Host", "getvideo.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", apikey)

	response, err := client.Do(req)

	defer func() {
		if err := response.Body.Close(); err != nil {
			gerr = err
		}
	}()

	body, gerr = ioutil.ReadAll(response.Body)
	if gerr != nil {
		return
	}

	return body, nil
}
