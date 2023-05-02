package http

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"onTime/config"
	"onTime/internal/transport/http/handler"
	middleware2 "onTime/internal/transport/http/middleware"
	"time"
)

type Server struct {
	cfg     *config.Config
	handler *handler.Manager
	App     *echo.Echo
	m       *middleware2.JWTAuth
}

func NewServer(cfg *config.Config, handler *handler.Manager, jwt *middleware2.JWTAuth) *Server {

	return &Server{
		cfg:     cfg,
		handler: handler,
		m:       jwt,
	}
}

func (s *Server) StartHTTPServer(ctx context.Context) error {
	s.App = s.BuildEngine()
	s.SetupRoutes()
	go func() {
		if err := s.App.Start(fmt.Sprintf(":%s", s.cfg.HTTP.PORT)); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%v\n", err)
		}
	}()
	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := s.App.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%v", err)
	}
	log.Print("server exited properly")
	return nil
}

func (s *Server) BuildEngine() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	return e
}
