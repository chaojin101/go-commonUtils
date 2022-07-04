package comUtils

import (
	"bufio"
	"io"
	"log"
	"os"
)

func Input() string {
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err != nil && err != io.EOF {
		log.Fatal("ReadString error:", err)
	}
	return input[:len(input)-2]
}
