package server

import (
	"context"
	"embed"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/bastean/laika/pkg/cmd/server/router"
)

//go:embed static
var Files embed.FS

func Run(port int) {
	log.Println("starting server")

	server := &http.Server{Addr: ":" + strconv.Itoa(port), Handler: router.New(&Files)}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err.Error())
		}
	}()

	log.Println("listening and serving HTTP on :" + strconv.Itoa(port))

	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-shutdown

	log.Println("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err.Error())
	}

	<-ctx.Done()

	log.Println("server exiting")
}
