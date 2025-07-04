package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github/rafaelgermann/chirpy/internal/auth"
	"net/http"

	"github.com/google/uuid"
)

const (
	eventUserUpgraded = "user.upgraded"
)

func (cfg *apiConfig) handlerWeebhook(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserId uuid.UUID `json:"user_id"`
		} `json:"data"`
	}

	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error(), err)
		return
	}

	if apiKey != cfg.polkaKey {
		respondWithError(w, http.StatusUnauthorized, "API Key is invalid", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	if params.Event != eventUserUpgraded {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	_, err = cfg.db.UpgradeUserChirpRed(r.Context(), params.Data.UserId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "Couldn't find user", err)
			return
		}
		respondWithError(w, http.StatusNotFound, "Couldn't update user", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
