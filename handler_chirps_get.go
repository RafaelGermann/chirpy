package main

import (
	"github/rafaelgermann/chirpy/internal/database"
	"net/http"
	"sort"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetChirps(w http.ResponseWriter, r *http.Request) {
	queryId := r.URL.Query().Get("author_id")
	authorID := uuid.Nil
	var err error
	if queryId != "" {
		authorID, err = uuid.Parse(queryId)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Invalid author Id", err)
			return
		}
	}
	sortDirection := "asc"
	querySort := r.URL.Query().Get("sort")

	if querySort == "desc" {
		sortDirection = "desc"
	}

	var dbChirps []database.Chirp
	if authorID != uuid.Nil {
		dbChirps, err = cfg.db.GetChirpsByUser(r.Context(), authorID)
	} else {
		dbChirps, err = cfg.db.GetChirps(r.Context())
	}
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get chirps", err)
		return
	}

	chirps := []Chirp{}
	for _, chirp := range dbChirps {
		chirps = append(chirps, Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		})
	}

	sort.Slice(chirps, func(i, j int) bool {
		if sortDirection == "asc" {
			return chirps[i].CreatedAt.Before(chirps[j].CreatedAt)
		}
		return chirps[i].CreatedAt.After(chirps[j].CreatedAt)
	})

	respondWithJSON(w, http.StatusOK, chirps)
}

func (cfg *apiConfig) handlerGetChirp(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("chirpID"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Id", err)
		return
	}

	chirp, err := cfg.db.GetChirp(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get chirp", err)
		return
	}

	respondWithJSON(w, http.StatusOK, Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	})
}
