package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/jub0bs/cors"
	_ "github.com/lib/pq"
	"github.com/mrdkvcs/go-base-backend/internal/database"
	"net/http"
	"os"
)

type apiConfig struct {
	DB *database.Queries
}

var db *sql.DB
var totalPoints int32 = 0
var goalPoints int32 = 0
var stopChan chan struct{} = make(chan struct{})

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
	db, err = sql.Open("postgres", dbUrl)
	if err != nil {
		fmt.Println("Could not connect to database")
	}
	apiConfig := apiConfig{DB: database.New(db)}
	corsMw, err := cors.NewMiddleware(cors.Config{
		Origins:        []string{"http://localhost:5173"},
		Methods:        []string{"GET", "POST", "DELETE"},
		RequestHeaders: []string{"Authorization"},
	})
	if err != nil {
		fmt.Println("Could not create cors middleware")
	}
	corsMw.SetDebug(true)
	router := http.NewServeMux()
	router.HandleFunc("POST /register", apiConfig.CreateUser)
	router.HandleFunc("POST /login", apiConfig.GetUserByEmail)
	router.HandleFunc("GET /user/{id}", apiConfig.GetUserApiKey)
	router.HandleFunc("GET /activities", apiConfig.middlewareAuth(apiConfig.GetActivites))
	router.HandleFunc("POST /activities", apiConfig.middlewareAuth(apiConfig.SetActivity))
	router.HandleFunc("DELETE /activities/{id}", apiConfig.DeleteActivity)
	router.HandleFunc("POST /activitieslog", apiConfig.middlewareAuth(apiConfig.SetActivityLog))
	router.HandleFunc("GET /dailyactivitylogs", apiConfig.middlewareAuth(apiConfig.GetDailyActivityLogs))
	router.HandleFunc("GET /dailypoints", apiConfig.middlewareAuth(apiConfig.GetDailyPoints))
	router.HandleFunc("POST /productivitystats", apiConfig.middlewareAuth(apiConfig.GetProductivityStats))
	router.HandleFunc("POST /productivitygoals", apiConfig.middlewareAuth(apiConfig.SetProductivityGoal))
	router.HandleFunc("POST /teams", apiConfig.middlewareAuth(apiConfig.CreateTeam))
	router.HandleFunc("GET /teams", apiConfig.middlewareAuth(apiConfig.GetUserTeams))
	router.HandleFunc("GET /teaminfo/{teamid}", apiConfig.GetTeamInfo)
	router.HandleFunc("GET /teamactivities/{teamid}", apiConfig.GetTeamActivities)
	handler := corsMw.Wrap(router)
	fmt.Println("Server running on port: " + port)
	http.ListenAndServe(":"+port, handler)
}
