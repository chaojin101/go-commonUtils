package comUtils

import (
	"io"
)

func GetText(url string) string {
	resp := Request("GET", url)
	bs, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	return string(bs)
}
