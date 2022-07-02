package comUtils

import (
	"io"
	"log"
	"net/http"
	"os"
)

func Download(url, path string, header *http.Header, client *http.Client) {
	resp := Request("GET", url, header, client)
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(f, resp.Body)

	resp.Body.Close()
	f.Close()
}
