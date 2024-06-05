package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mrdkvcs/go-base-backend/internal/database"
	"net/http"
	"os"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("Couldnt get port number from .env file")
	}
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		fmt.Println("Could not find database url in .env file")
	}
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		fmt.Println("Could not connect to database")
	}
	apiConfig := apiConfig{DB: database.New(db)}
	router := http.NewServeMux()
	router.HandleFunc("POST /register", apiConfig.CreateUser)
	router.HandleFunc("POST /login", apiConfig.GetUserByEmail)
	router.HandleFunc("GET /user/{id}", apiConfig.GetUserApiKey)
	fmt.Println("Server running on port: " + port)
	http.ListenAndServe(":"+port, router)
}
