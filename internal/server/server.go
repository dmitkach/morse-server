package server

import (
	"log"
	"morse-server/internal/handlers"
	"net/http"
	"time"
)

type Server struct {
	Logger log.Logger
	Server http.Server
}

func Create(logger log.Logger) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	server := &Server{
		Logger: logger,
		Server: http.Server{
			Addr:              "localhost:8080",
			ErrorLog:          &logger,
			ReadHeaderTimeout: time.Second * 5,
			WriteTimeout:      time.Second * 10,
			IdleTimeout:       time.Second * 15,
			Handler:           mux,
		},
	}

	return server
}
