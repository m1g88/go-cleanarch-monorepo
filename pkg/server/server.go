package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	Port   string
	Router http.Handler
}

func New(port string, r http.Handler) *Server {
	return &Server{Port: port, Router: r}
}

func (s *Server) RunWithGracefulShutdown() {
	srv := &http.Server{
		Addr:    ":" + s.Port,
		Handler: s.Router,
	}
	go func() {
		log.Printf("Listen on %s", s.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error while listen and serve: %v", err)
		}
	}()

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, os.Interrupt, syscall.SIGTERM)
	<-wait
	log.Print("Shutting down http server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Cannot shutdown server: %v", err)
	}

	log.Println("Done gracefully")
}
