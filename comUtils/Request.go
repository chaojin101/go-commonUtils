package comUtils

import (
	"log"
	"net/http"
)

func Request(method, url string) *http.Response {
	req, _ := http.NewRequest(method, url, nil)
	req.Header = Header
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}
