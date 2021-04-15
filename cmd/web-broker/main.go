package main

import (
	"context"
	"flag"
	"github.com/phpCoder88/web-broker/internal/app/endpoint/queue"
	"github.com/phpCoder88/web-broker/internal/pkg/repository/file"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/phpCoder88/web-broker/internal/app/service"
)

func main() {
	log.Print("Starting the app")

	port := flag.String("port", "8000", "Port")
	storageName := flag.String("storage", "storage.json", "data storage")
	shutdownTimeout := flag.Int64("shutdown_timeout", 3, "shutdown timeout")

	repo := file.NewRepo(*storageName)
	svc := service.New(repo)

	serv := http.Server{
		Addr:    net.JoinHostPort("", *port),
		Handler: queue.RegisterPublicHTTP(svc),
	}
	go func() {
		if err := serv.ListenAndServe(); err != nil {
			log.Fatalf("listen and serve err: %v", err)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	log.Print("Started app")

	sig := <-interrupt

	log.Printf("Sig: %v, stopping app", sig)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*shutdownTimeout)*time.Second)
	defer cancel()
	if err := serv.Shutdown(ctx); err != nil {
		log.Printf("shutdown err: %v", err)
	}
}
