package tool

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
)

func PerformHTTP_Request(method, url string, requestBody io.Reader, responseBody interface{}) error {
	req, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return err
	}

	// TODO: Make this a confiruable option in the `initHTTP` function
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	err = json.NewDecoder(resp.Body).Decode(responseBody)
	if err != nil {
		return err
	}

	// Success!
	return nil
}
