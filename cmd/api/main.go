package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/romeufcrosa/best-route-finder/configurations"
	"github.com/romeufcrosa/best-route-finder/configurations/readers"
	"github.com/romeufcrosa/best-route-finder/providers"
	"github.com/romeufcrosa/best-route-finder/services/api"
)

var (
	env        = os.Getenv("ENV")
	listenAddr string
)

func main() {
	ctx := context.Background()

	channel := configureService(ctx)
	router := api.Router()
	listenAddr = configurations.ListenAddr()
	server := &http.Server{
		Addr:         listenAddr,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Println("service running in", listenAddr)

	go startServer(ctx, server)
	stopAtSignal(ctx, server, channel)

	log.Println("service stopped")
}

func configureService(ctx context.Context) chan os.Signal {
	log.Println("Provisioning configs for Env:", env)
	reader := readers.NewFileReader("./configurations/" + env + ".json")
	if err := configurations.Load(ctx, reader); err != nil {
		log.Fatalln("could not load configurations")
	}

	sqlPool, err := configurations.Pool()
	if err != nil {
		log.Fatal(err)
	}

	providers.Configure(
		providers.NewParams(sqlPool),
		configurations.IsConfigured,
	)
	providers.RegisterRepositoryProviders()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	return c
}

func startServer(ctx context.Context, server *http.Server) {
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func stopAtSignal(background context.Context, server *http.Server, stop chan os.Signal) {
	<-stop

	ctx, cancel := context.WithTimeout(background, 15*time.Second)
	defer cancel()

	log.Println("shutting down service")
	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln("could not gracefully shutdown service")
	}

	close(stop)
}
