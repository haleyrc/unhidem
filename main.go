package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	file := flag.String("f", "", "The file to check")
	flag.Parse()

	if file == nil || *file == "" {
		fmt.Printf("File name must be supplied.\n")
		os.Exit(2)
	}

	bytes, err := ioutil.ReadFile(*file)
	if err != nil {
		fmt.Printf("Could not read file: %v\n", err)
		os.Exit(1)
	}

	data := string(bytes)
	for num, line := range strings.Split(data, "\n") {
		for _, c := range []byte(line) {
			switch {
			case c == 10:
			case c == 13:
			case c >= 32 && c <= 126:
			default:
				fmt.Printf("%d (%v): %s\n", num, c, line)
			}
		}
	}
}
