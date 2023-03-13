package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	msg := os.Args[1]
	length := len(msg)
	str := strings.Repeat("!", length)

	fmt.Println(strings.ToUpper(str + msg + str))
}
