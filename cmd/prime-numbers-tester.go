package main

import (
	"fmt"
	"github.com/icemont/prime-numbers-tester/internal/config"
	"github.com/icemont/prime-numbers-tester/router"
	"net/http"
)

func main() {
	config.Initialize()
	httpHandler := router.InitRouter()

	s := http.Server{
		Addr:        fmt.Sprintf(":%d", config.Server.HttpPort),
		Handler:     httpHandler,
		ReadTimeout: config.Server.ReadTimeout,
	}

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		httpHandler.Logger.Fatal(err)
	}
}
