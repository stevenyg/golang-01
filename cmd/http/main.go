package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/rs/zerolog/log"
	handlerAPI "github.com/stevenyg/golang-01/handler/api"
	internalConfig "github.com/stevenyg/golang-01/internal/config"
)

func main() {
	// read from vault
	var ctx = context.Background()
	var vaultConfig = internalConfig.Config{
		Env: internalConfig.EnvConfig{
			Env:         "local",
			ServiceName: "learn",
		},
		API: internalConfig.APIConfig{
			Port: "8089",
		},
	}

	globalConfig := internalConfig.New()
	globalConfig.ApplyConfig(vaultConfig)

	// assign handler
	handlerAPI := handlerAPI.New(handlerAPI.Config{
		Config: globalConfig,
	})

	// start handler
	log.Info().Msg("register API handler")
	handlerAPI.Register(ctx)

	// BOILERPLATE TO HOLD SERVER RUNNING
	sigs := make(chan os.Signal, 2)
	signal.Notify(sigs, os.Interrupt)

	<-sigs
	log.Info().Msg("Closing HTTP server ...")

	done := make(chan struct{})
	go func() {
		defer close(done)
	}()

	select {
	case <-sigs:
		log.Info().Msg("Quitting without waiting for graceful close ...")
	case <-done:
	}
}
