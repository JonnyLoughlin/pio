package server

import (
	"fmt"
	"net/http"
	"time"
	// _ "github.com/joho/godotenv/autoload" // Used for loading .env variables
)

type Server struct {
	port int
}

func NewServer() *http.Server {
	NewServer := &Server{
		port: 9000,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
