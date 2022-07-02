package comUtils

import (
	"log"
	"net/http"
)

func Request(method, url string, header *http.Header, client *http.Client) *http.Response {
	req, _ := http.NewRequest(method, url, nil)
	req.Header = *header
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}
