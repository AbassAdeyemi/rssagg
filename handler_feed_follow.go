package main

import (
	"encoding/json"
	"fmt"
	"github.com/AbassAdeyemi/rssagg/internal/database"
	"github.com/google/uuid"
	"net/http"
	"time"
	"github.com/go-chi/chi/v5"
)


func (apiConfig *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, dbUser database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err!= nil {
        respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
        return
    }

	feedFollow, err := apiConfig.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
        CreatedAt: time.Now().UTC(),
        UpdatedAt: time.Now().UTC(),
        UserID:    dbUser.ID,
        FeedID:    params.FeedID,
	})

	if err!= nil {
        respondWithError(w, 400, fmt.Sprintf("Error creating feed follow: %v", err))
        return
    }

	respondWithJson(w, 201, databaseFeedFollowToFeedFollow(feedFollow))
}

func (apiConfig *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, dbUser database.User) {

	feedFollows, err := apiConfig.DB.GetFeedFollows(r.Context(), dbUser.ID)

	if err!= nil {
        respondWithError(w, 400, fmt.Sprintf("Error fetching feed follows: %v", err))
        return
    }

	respondWithJson(w, 200, databaseFeedFollowsToFeedFollows(feedFollows))
}

func (apiConfig *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, dbUser database.User) {

	feedFollowIdParam := chi.URLParam(r, "feed_follow_id")
	feedFollowId, err := uuid.Parse(feedFollowIdParam) 

	if(err != nil) {
		respondWithError(w, 400, fmt.Sprintf("Error parsing feed id: %v", err))
        return
	}
	err = apiConfig.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams {
		ID:     feedFollowId,
        UserID: dbUser.ID,
	})

	if err!= nil {
        respondWithError(w, 400, fmt.Sprintf("Error creating feed follow: %v", err))
        return
    }

	// respondWithJson(w, 200, {})
}

