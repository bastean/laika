package main

import (
	"os"

	"github.com/bastean/laika"
)

var port = os.Getenv("PORT")
var urls = []string{"localhost:" + port}

func main() {
	laika.Sniff(urls).Emails()
}
