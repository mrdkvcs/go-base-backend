package main

import (
	"github.com/google/uuid"
	"github.com/mrdkvcs/go-base-backend/internal/database"
	"time"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	ApiKey       string    `json:"api_key"`
}

func databaseUserToUser(dbuser database.User) User {
	return User{
		ID:           dbuser.ID,
		CreatedAt:    dbuser.CreatedAt,
		UpdatedAt:    dbuser.UpdatedAt,
		Username:     dbuser.Username,
		Email:        dbuser.Email,
		PasswordHash: dbuser.PasswordHash,
		ApiKey:       dbuser.ApiKey,
	}
}
