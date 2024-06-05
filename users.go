package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/mrdkvcs/go-base-backend/internal/database"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func (apiCfg *apiConfig) GetUserApiKey(w http.ResponseWriter, r *http.Request) {
	apiKey := r.PathValue("id")
	user, err := apiCfg.DB.GetUserByApikey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, 400, "Couldnt get user by api key")
	}
	respondWithJson(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldnt parse json: %s", err))
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		respondWithError(w, 400, "Couldnt hash password")
		return
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:           uuid.New(),
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
		Username:     params.Username,
		PasswordHash: string(hashedPassword),
		Email:        params.Email,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldnt create user: %s", err))
		return
	}
	respondWithJson(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldnt parse json: %s", err))
		return
	}
	user, err := apiCfg.DB.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, 400, "Couldnt get user by email : User does not exist")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(params.Password))
	if err != nil {
		respondWithError(w, 400, "Incorrect password")
		return
	}
	respondWithJson(w, 200, databaseUserToUser(user))
}
