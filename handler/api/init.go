package api

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type (
	Config struct {
		Config InternalConfigInterface
	}

	module struct {
		config InternalConfigInterface
		serv   *http.Server
	}
)

func New(c Config) *module {
	return &module{
		config: c.Config,
	}
}

func (m *module) Register(ctx context.Context) {
	config := m.config.GetConfig()
	m.serv = &http.Server{
		Addr:              fmt.Sprintf(":%s", config.API.Port),
		Handler:           m.attachHandler(),
		ReadHeaderTimeout: 60 * time.Second,
	}
	go m.execute(ctx)
}

func (m *module) Stop(ctx context.Context) error {
	return m.serv.Shutdown(ctx)
}
