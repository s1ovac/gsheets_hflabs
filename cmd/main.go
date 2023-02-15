package main

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"github.com/s1ovac/gdoc/internal/codes"
	"github.com/s1ovac/gdoc/internal/config"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())
	cfg := config.NewSheetConfig()
	router := httprouter.New()
	handler := codes.NewHandler(ctx, cfg)
	handler.Register(router)
	start(router)
}

func start(router *httprouter.Router) {
	var listener net.Listener
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(server.Serve(listener))
}
