package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (m *module) execute(ctx context.Context) {
	config := m.config.GetConfig()
	log.Info().Msg("api served on " + config.API.Port)
	if err := m.serv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Msg("executeAPI - ListenAndServe")
	}
}

func (m *module) attachHandler() http.Handler {
	r := chi.NewRouter()
	config := m.config.GetConfig()
	logger := httplog.NewLogger(config.Env.ServiceName, httplog.Options{
		JSON: true,
	})

	r.Use(middleware.RequestID)
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.Recoverer)

	r.Route("/learn", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/about-me", m.learnV1AboutMe)
		})
	})

	return r
}

func writeReply(w http.ResponseWriter, status int, payload map[string]interface{}) error {
	resp, err := json.Marshal(payload)
	if err != nil {
		return errors.Wrap(err, "writeReply - json.Marshal")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err := w.Write(resp); err != nil {
		return errors.Wrap(err, "writeReply - w.Write")
	}
	return nil
}
