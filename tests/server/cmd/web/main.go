package main

import (
	"context"
	"embed"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bastean/laika/test/server/cmd/web/router"
)

//go:embed static
var Files embed.FS

var Port = os.Getenv("PORT")
var Server = &http.Server{Addr: ":" + Port, Handler: router.New(&Files)}

func main() {
	log.Println("starting server")

	go func() {
		if err := Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err.Error())
		}
	}()

	log.Println("listening and serving HTTP on :" + Port)

	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-shutdown

	log.Println("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := Server.Shutdown(ctx); err != nil {
		log.Fatal(err.Error())
	}

	<-ctx.Done()

	log.Println("server exiting")
}
