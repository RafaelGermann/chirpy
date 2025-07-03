package main

import (
	"github/rafaelgermann/chirpy/internal/auth"
	"net/http"
	"time"
)

func (cfg *apiConfig) handlerRefreshToken(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Token string `json:"token"`
	}

	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error(), err)
		return
	}

	dbToken, err := cfg.db.GetRefreshToken(r.Context(), refreshToken)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "invalid token", err)
		return
	}

	if dbToken.ExpiresAt.Before(time.Now().UTC()) || dbToken.RevokedAt.Valid {
		respondWithError(w, http.StatusUnauthorized, "expired refresh token", err)
		return
	}

	token, err := auth.MakeJWT(dbToken.UserID, cfg.jwtSecret, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't create token", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		Token: token,
	})
}
