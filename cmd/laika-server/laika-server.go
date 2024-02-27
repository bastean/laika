package main

import (
	"flag"
	"fmt"

	"github.com/bastean/laika/pkg/cmd/server"
)

const cli = "laika-server"

var port int

func usage() {
	fmt.Printf("Usage: %s [OPTIONS]\n", cli)
	fmt.Printf("\nSniff Test Server\n")
	fmt.Printf("\nE.g.: %s -p 8080\n\n", cli)
	flag.PrintDefaults()
}

func main() {
	flag.IntVar(&port, "p", 8080, "Port")

	flag.Usage = usage

	flag.Parse()

	server.Run(port)
}
