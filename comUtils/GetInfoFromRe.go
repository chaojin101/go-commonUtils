package comUtils

import (
	"log"
	"regexp"
)

func GetInfoFromRe(s, pattern string) string {
	re := regexp.MustCompile(pattern)
	infoS := re.FindStringSubmatch(s)
	if infoS == nil {
		log.Fatal("Regexp get info fail")
	}

	return infoS[1]
}
