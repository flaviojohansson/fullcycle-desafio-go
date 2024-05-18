package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s := "Full Cycle"

	if len(os.Args) > 1 {
		s = strings.Join(os.Args[1:], " ")
	}

	fmt.Println(s, "Rocks!!")
}
