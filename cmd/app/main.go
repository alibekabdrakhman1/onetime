package main

import (
	"context"
	"log"
	"onTime/config"
	"onTime/internal/service"
	"onTime/internal/storage"
	"onTime/internal/transport/http"
	"onTime/internal/transport/http/handler"
	"onTime/internal/transport/http/middleware"
	"os"
	"os/signal"
)

func main() {
	log.Fatal(run())
}
func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	gracefullyShutdown(cancel)
	conf, err := config.New()
	if err != nil {
		return err
	}
	repo, err := storage.NewStorage(ctx, conf)
	if err != nil {
		return err
	}
	svc, err := service.NewManager(repo)
	if err != nil {
		return err
	}

	jwt := middleware.NewJWTAuth(conf, *svc)
	h := handler.NewManager(svc, jwt) //R6J9M97WIKwcnIqe
	HTTPServer := http.NewServer(conf, h, jwt)
	return HTTPServer.StartHTTPServer(ctx)
}
func gracefullyShutdown(c context.CancelFunc) {
	osC := make(chan os.Signal, 1)
	signal.Notify(osC, os.Interrupt)
	go func() {
		log.Print(<-osC)
		c()
	}()
}
