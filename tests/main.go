package main

import (
	"github.com/bastean/laika"
)

var urls = []string{"http://localhost:8080/"}

func main() {
	sniff := laika.Sniff(urls)

	emails := sniff.Emails()

	if len(emails) != 0 {
		err := sniff.Dump("laika")

		if err != nil {
			panic(err)
		}
	}
}
