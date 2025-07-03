package main

import (
	"github/rafaelgermann/chirpy/internal/auth"
	"net/http"
)

func (cfg *apiConfig) handlerRevokeToken(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error(), err)
		return
	}

	err = cfg.db.RevokeToken(r.Context(), refreshToken)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error revoking the token", err)
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}
