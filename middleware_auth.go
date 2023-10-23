package main

import (
	"net/http"
	"github.com/AbassAdeyemi/rssagg/internal/database"
	"github.com/AbassAdeyemi/rssagg/internal/auth"
	"fmt"

)


type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiConfig *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if(err != nil) {
		respondWithError(w, 401, fmt.Sprintf("auth error %v", err))
        return
	}

	user, err := apiConfig.DB.GetUserByAPIKey(r.Context(), apiKey)

	if err!= nil {
        respondWithError(w, 404, fmt.Sprintf("user not found: %v", err))
        return
    }

	handler(w, r, user)

	}
}