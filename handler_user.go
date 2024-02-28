package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jorg1piano/rssapp/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	params := parameters{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json: %v", err))
		return
	}

	log.Default().Printf("name %s", params)

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: sql.NullString{
			String: params.Name,
			Valid:  true,
		},
	})
	if err != nil {
		respondWithError(w, 400, "Failed to create a user")
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUserByApiKey(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		ApiKey string `json:"apiKey"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), params.ApiKey)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Failed to get user by api key %v", params.ApiKey))
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}
