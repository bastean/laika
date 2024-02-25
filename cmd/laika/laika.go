package main

import (
	"log"

	"github.com/bastean/laika"
)

var urls = []string{"http://localhost:8080/"}

func main() {
	fromZero := laika.NewEmptyData()

	// store := laika.NewLocalJsonStore(".", "laika")

	// fromExistingData := laika.ReadDataFromStore(store)

	sniff := laika.New(fromZero)

	// sniff.SetStore(store)

	sniff.ContentFromUrls(urls)

	sniff.EmailsFromContent()

	log.Println(sniff.SniffedEmails())

	// sniff.SaveSniffed()
}
