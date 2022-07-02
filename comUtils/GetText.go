package comUtils

import (
	"io"
	"net/http"
)

func GetText(url string, header *http.Header, client *http.Client) string {
	resp := Request("GET", url, header, client)
	bs, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	return string(bs)
}
