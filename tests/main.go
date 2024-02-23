package main

import (
	"github.com/bastean/laika"
)

var urls = []string{"http://localhost:8080/"}

func main() {
	sniff := laika.Sniff(urls)

	sniff.ContentFromUrls()

	sniff.Emails()
}
