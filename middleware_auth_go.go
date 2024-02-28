package main

import (
	"fmt"
	"net/http"

	"github.com/jorg1piano/rssapp/internal/database"
	"github.com/jorg1piano/rssapp/internal/database/auth"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("auth error: %v", err))
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 404, fmt.Sprintf("could not get user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
