package main

import (
	"flag"

	"github.com/bastean/laika/pkg/cmd/server"
)

var port int

func main() {
	flag.IntVar(&port, "p", 8080, "Port")

	flag.Parse()

	server.Run(port)
}
