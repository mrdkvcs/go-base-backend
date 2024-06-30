package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/mrdkvcs/go-base-backend/internal/database"
	"net/http"
	"time"
)

func (apiCfg *apiConfig) GetActivites(w http.ResponseWriter, r *http.Request, user database.User) {
	activities, err := apiCfg.DB.GetActivities(r.Context(), user.ID)
	if err != nil {
		respondWithJson(w, 400, fmt.Sprintf("Error getting activities: %v", err))
	}
	respondWithJson(w, 200, databaseActivityToActivity(activities))
}

func (apiCfg *apiConfig) SetActivity(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name   string `json:"name"`
		Points int32  `json:"points"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error decoding parameters: %v", err))
		return
	}
	activityUUID := uuid.New()
	err = apiCfg.DB.SetAllActivity(r.Context(), database.SetAllActivityParams{
		ID:   activityUUID,
		Type: "custom",
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error setting activity: %v", err))
		return
	}
	err = apiCfg.DB.SetCustomActivity(r.Context(), database.SetCustomActivityParams{
		ActivityID: activityUUID,
		UserID:     user.ID,
		Name:       params.Name,
		Points:     params.Points,
		CreatedAt:  time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error setting custom activity: %v", err))
		return
	}
	respondWithJson(w, 200, fmt.Sprintf("Activity successfully set: %s", activityUUID))
}

func (apiCfg *apiConfig) DeleteActivity(w http.ResponseWriter, r *http.Request) {
	activityUUID := r.PathValue("id")
	parsedactvityUUID, err := uuid.Parse(activityUUID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing activity UUID: %v", err))
		return
	}
	err = apiCfg.DB.DeleteActivity(r.Context(), parsedactvityUUID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error deleting activity: %v", err))
		return
	}
	respondWithJson(w, 200, "Activity successfully deleted")
}
