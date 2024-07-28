package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/mrdkvcs/go-base-backend/internal/database"
	"net/http"
)

func (apiCfg *apiConfig) CreateTeam(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name         string `json:"team_name"`
		TeamIndustry string `json:"team_industry"`
		TeamSize     int32  `json:"team_size"`
		IsPrivate    bool   `json:"is_private"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error in parsing json: %s", err))
		return
	}
	team, err := apiCfg.DB.CreateTeam(r.Context(), database.CreateTeamParams{
		ID:           uuid.New(),
		Name:         params.Name,
		TeamIndustry: params.TeamIndustry,
		TeamSize:     params.TeamSize,
		IsPrivate:    params.IsPrivate,
		CreatedBy:    user.ID,
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error in creating team: %s", err))
		return
	}
	err = apiCfg.DB.CreateTeamMembership(r.Context(), database.CreateTeamMembershipParams{
		ID:     uuid.New(),
		TeamID: team.ID,
		UserID: user.ID,
		Role:   "owner",
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error in creating team membership: %s", err))
		return
	}
	respondWithJson(w, 200, databaseTeamToTeam(team))
}

func (apiCfg *apiConfig) GetUserTeams(w http.ResponseWriter, r *http.Request, user database.User) {
	userteams, err := apiCfg.DB.GetUserTeams(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error in getting user teams: %s", err))
		return
	}
	respondWithJson(w, 200, databaseUserTeamsToUserTeams(userteams))
}

func (apiCfg *apiConfig) GetTeamInfo(w http.ResponseWriter, r *http.Request) {
	teamId := r.PathValue("teamid")
	parsedTeamUUID, err := uuid.Parse(teamId)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error in parsing uuid: %s", err))
		return
	}
	team, err := apiCfg.DB.GetTeamInFo(r.Context(), parsedTeamUUID)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error in getting user teams: %s", err))
		return
	}
	respondWithJson(w, 200, databaseTeamInfoToTeamInfo(team))
}

func (apiCfg *apiConfig) GetTeamActivities(w http.ResponseWriter, r *http.Request) {
	teamId := r.PathValue("teamid")
	parsedTeamUUID, err := uuid.Parse(teamId)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error in parsing uuid: %s", err))
		return
	}
	teamactivities, err := apiCfg.DB.GetTeamActivities(r.Context(), parsedTeamUUID)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error in getting  team activites: %s", err))
		return
	}
	respondWithJson(w, 200, databaseTeamActivityToTeamActivity(teamactivities))
}
