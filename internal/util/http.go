package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func HttpRequest(accessToken string, method string, url string, body []byte, useCache bool) ([]byte, error) {
	if useCache {
		respBody := getFromCache(url, body)
		if respBody != nil {
			log.Printf("fetching url: %s [cached]\n", url)
			return respBody, nil
		}
	}

	log.Printf("fetching url: %s\n", url)

	bearer := fmt.Sprintf("Bearer %s", accessToken)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	req.Header.Add("Authorization", bearer)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if useCache {
		err = saveToCache(url, body, respBody)
		if err != nil {
			return nil, err
		}
	}

	return respBody, nil
}

