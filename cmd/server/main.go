package main

import (
	"log"
	"sync"

	"github.com/egnitelabs/engine/internal/config"
	"github.com/egnitelabs/engine/internal/logger"
	"github.com/egnitelabs/engine/internal/server"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	conf, err := config.GetConfig()
	if err != nil {
		log.Fatalf("failed to initialize configs: %v", err.Error())
	}
	err = logger.Init(conf.LogLevel, conf.LogTimeFormat)
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	server := server.New(conf)
	go server.StartGRPC()
	go server.StartHTTP()

	wg.Wait()
}
