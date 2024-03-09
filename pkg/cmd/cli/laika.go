package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/bastean/laika"
)

const cli = "laika"

var jsonStore string
var isUrls bool
var silent bool
var err error

func usage() {
	fmt.Printf("Usage: %s [OPTIONS] sources... \n", cli)
	fmt.Printf("\nSniffs the content of the sources\n")
	fmt.Printf("\nE.g.: %s -jsonStore \"laika\" -urls http://localhost:8080\n\n", cli)
	flag.PrintDefaults()
}

func noContentError() {
	log.Println("No content to sniff")
	flag.PrintDefaults()
	os.Exit(2)
}

func Run() {
	flag.StringVar(&jsonStore, "jsonStore", "", "Store filepath to save the sniffed content (default \"In Memory\")")
	flag.BoolVar(&isUrls, "urls", false, "If the sources for sniffing content are urls (Required)")
	flag.BoolVar(&silent, "silent", false, "Do not show the sniffed content")

	flag.Usage = usage

	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		return
	}

	sources := flag.Args()

	store := laika.NewInMemoryStore()

	if jsonStore != "" {
		path := filepath.Dir(jsonStore)
		filename := filepath.Base(jsonStore)

		store = laika.NewLocalJsonStore(path, filename)
	}

	data := laika.NewEmptyData()

	if store != nil {
		data, err = laika.ReadDataFromStore(store)

		if err != nil {
			log.Fatal(err)
		}
	}

	sniff := laika.New(data)

	sniff.SetStore(store)

	switch {
	case isUrls:
		sniff.FromUrls(sources)
	default:
		noContentError()
	}

	switch {
	case !silent:
		log.Println(sniff.SniffedEmails())
	}

	if sniff.Store != nil {
		sniff.SaveSniffed()
	}
}
