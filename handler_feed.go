package main

import (
	"encoding/json"
	"net/http"
	"time"
	"log"

	"github.com/josefmay/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerFeedsCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string	`json:"name"`
		Url  string `json:"url"`
	}

	type response struct {
		Feed 		 Feed	`json:"feed"`
		FeedFollow   FeedFollow `json:"feed_follow"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed")
		return
	}

	feedFollow, err := cfg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create a feed follow")
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		Feed: databaseFeedToFeed(feed),
		FeedFollow: databaseFeedFollowToFeedFollow(feedFollow),
	})
	
}

func (cfg *apiConfig) handlerFeedsGet(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't get feeds")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedsToFeeds(feeds))
}

func (cfg *apiConfig) handlerFeedFollowsCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		ID string `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	real_feed_id, err := uuid.Parse(params.ID)
	if err != nil{
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode feed id")
		return
	}

	feedFollow, err := cfg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    real_feed_id,
	})
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't create a feed follow")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedFollowToFeedFollow(feedFollow))
}

func (cfg *apiConfig) handlerFeedFollowsGet(w http.ResponseWriter, r *http.Request, user database.User) {
	feed_follows, err := cfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't get feed follows for user")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedFollowsToFeedFollows(feed_follows))
}

func (cfg *apiConfig) handlerFeedFollowsDelete(w http.ResponseWriter, r *http.Request) {
	follow_id := r.PathValue("feedFollowID")
	id, err := uuid.Parse(follow_id)
	if err != nil{
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode feed id")
		return
	}
	err = cfg.DB.DeleteFeedFollows(r.Context(),id)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't delete a feed follow")
		return
	}

	respondWithJSON(w, http.StatusOK, "Success")
}