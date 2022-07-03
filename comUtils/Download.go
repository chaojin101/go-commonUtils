package comUtils

import (
	"io"
	"log"
	"os"
)

func Download(url, path string) {
	resp := Request("GET", url)
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(f, resp.Body)

	resp.Body.Close()
	f.Close()
}
