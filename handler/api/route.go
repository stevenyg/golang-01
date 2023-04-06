package api

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func (m *module) learnV1AboutMe(w http.ResponseWriter, r *http.Request) {
	var reply = map[string]interface{}{
		"status":  http.StatusOK,
		"message": "",
	}
	if err := writeReply(w, http.StatusOK, reply); err != nil {
		log.Error().Err(err).Msg("manualCancelRecoveryHandler - writeReply")
		return
	}
}
