package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/mrdkvcs/go-base-backend/internal/database"
	"net/http"
	"time"
)

func (apiCfg *apiConfig) SetActivityLog(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		ActivityID string `json:"activity_id"`
		StartTime  string `json:"start_time"`
		EndTime    string `json:"end_time"`
		Points     int32  `json:"points"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error decoding parameters: %v", err))
		return
	}
	parsedActivityID, err := uuid.Parse(params.ActivityID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing activity id: %v", err))
		return
	}
	startTime, err := time.Parse("2006-01-02T15:04:00.000Z", params.StartTime)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing start time: %v", err))
		return
	}
	endTime, err := time.Parse("2006-01-02T15:04:00.000Z", params.EndTime)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing end time: %v", err))
		return
	}
	duration := endTime.Sub(startTime)
	durationMinutes := int32(duration.Minutes())
	points := durationMinutes * params.Points
	activityLog, err := apiCfg.DB.SetActivityLog(r.Context(), database.SetActivityLogParams{
		ID:         uuid.New(),
		UserID:     user.ID,
		ActivityID: parsedActivityID,
		Duration:   durationMinutes,
		Points:     points,
		LoggedAt:   time.Now(),
		StartTime:  startTime,
		EndTime:    endTime,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error setting activity log: %v", err))
		return
	}
	respondWithJson(w, 200, databaseActivityLogToActivityLog(activityLog))
}

func (apiCfg *apiConfig) GetDailyActivityLogs(w http.ResponseWriter, r *http.Request, user database.User) {
	activityDailyLogs, err := apiCfg.DB.GetDailyActivityLogs(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error getting daily activity logs: %v", err))
		return
	}
	respondWithJson(w, 200, databaseDailyActivityLogsToDailyActivityLogs(activityDailyLogs))
}

func (apiCfg *apiConfig) GetDailyActivityPoints(w http.ResponseWriter, r *http.Request, user database.User) {
	points, err := apiCfg.DB.GetDailyActivityPoints(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error getting daily activity points: %v", err))
		return
	}
	respondWithJson(w, 200, points)
}
