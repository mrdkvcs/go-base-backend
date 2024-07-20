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

type Activity struct {
	ActivityID uuid.UUID `json:"activity_id"`
	Name       string    `json:"name"`
	Points     int32     `json:"points"`
	Type       string    `json:"type"`
}

type ActivityLog struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	ActivityID uuid.UUID `json:"activity_id"`
	Duration   int32     `json:"duration"`
	Points     int32     `json:"points"`
	LoggedAt   time.Time `json:"logged_at"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
}
type DailyActivityLog struct {
	ActivityID   uuid.UUID   `json:"activity_id"`
	Points       int32       `json:"points"`
	Duration     int32       `json:"duration"`
	StartTime    string      `json:"start_time"`
	EndTime      string      `json:"end_time"`
	ActivityName interface{} `json:"activity_name"`
}

type TotalAndAveragePoints struct {
	TotalPoints   string `json:"total_points"`
	AveragePoints string `json:"average_points"`
}

type ProductivityDay struct {
	Date        time.Time   `json:"date"`
	TotalPoints interface{} `json:"total_points"`
}

type ProductivityStats struct {
	ProductivvityPoints TotalAndAveragePoints `json:"productivity_points"`
	BestProductivityDay ProductivityDay       `json:"best_productivity_day"`
	ProductivityDays    []ProductivityDay     `json:"productivity_days"`
}

type DailyPoints struct {
	TotalPoints interface{} `json:"total_points"`
	GoalPoints  interface{} `json:"goal_points"`
}

type Team struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"team_name"`
	TeamIndustry string    `json:"team_industry"`
	TeamSize     int32     `json:"team_size"`
	IsPrivate    bool      `json:"is_private"`
	CreatedBy    uuid.UUID `json:"created_by"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserTeams struct {
	TeamID   uuid.UUID `json:"team_id"`
	TeamName string    `json:"team_name"`
	Role     string    `json:"role"`
}

func databaseUserTeamsToUserTeams(dbuserteams []database.GetUserTeamsRow) []UserTeams {
	userTeams := []UserTeams{}
	for _, dbuserteam := range dbuserteams {
		userTeam := UserTeams{TeamID: dbuserteam.TeamID, TeamName: dbuserteam.TeamName, Role: dbuserteam.Role}
		userTeams = append(userTeams, userTeam)
	}
	return userTeams
}

func databaseTeamToTeam(dbteam database.Team) Team {
	return Team{
		ID:           dbteam.ID,
		Name:         dbteam.Name,
		TeamIndustry: dbteam.TeamIndustry,
		TeamSize:     dbteam.TeamSize,
		IsPrivate:    dbteam.IsPrivate,
		CreatedBy:    dbteam.CreatedBy,
		CreatedAt:    dbteam.CreatedAt,
		UpdatedAt:    dbteam.UpdatedAt,
	}
}

func DatabaseDailyPointsToDailyPoints(DbDailyPoints database.GetDailyPointsRow) DailyPoints {
	return DailyPoints{
		TotalPoints: DbDailyPoints.TotalPoints,
		GoalPoints:  DbDailyPoints.GoalPoints,
	}
}

func databaseProductivityStatsToProductivityStats(productivityStats DatabaseProductivityStats) ProductivityStats {
	productivityDays := []ProductivityDay{}
	for _, productivityDay := range productivityStats.ProductivityDays {
		productivityDays = append(productivityDays, ProductivityDay{Date: productivityDay.Date, TotalPoints: productivityDay.TotalPoints})
	}
	totalAveragePoints := TotalAndAveragePoints{
		TotalPoints:   productivityStats.ProductivityPoints.TotalPoints,
		AveragePoints: productivityStats.ProductivityPoints.AveragePointsPerDay,
	}
	bestProductivityDay := ProductivityDay{
		Date:        productivityStats.BestProductivityDay.Date,
		TotalPoints: productivityStats.BestProductivityDay.TotalPoints,
	}
	return ProductivityStats{
		ProductivvityPoints: totalAveragePoints,
		BestProductivityDay: bestProductivityDay,
		ProductivityDays:    productivityDays,
	}
}

func databaseDailyActivityLogsToDailyActivityLogs(dbdaily []database.GetDailyActivityLogsRow) []DailyActivityLog {
	dailyActivityLogs := []DailyActivityLog{}
	for _, dbdaily := range dbdaily {
		dailyActivityLogs = append(dailyActivityLogs, DailyActivityLog{
			ActivityID:   dbdaily.ActivityID,
			Points:       dbdaily.Points,
			Duration:     dbdaily.Duration,
			StartTime:    dbdaily.StartTime.Format("15:04"),
			EndTime:      dbdaily.EndTime.Format("15:04"),
			ActivityName: dbdaily.ActivityName,
		})
	}
	return dailyActivityLogs
}

func databaseActivityLogToActivityLog(dbLog database.ActivityLog) ActivityLog {
	return ActivityLog{
		ID:         dbLog.ID,
		UserID:     dbLog.UserID,
		ActivityID: dbLog.ActivityID,
		Duration:   dbLog.Duration,
		Points:     dbLog.Points,
		LoggedAt:   dbLog.LoggedAt,
		StartTime:  dbLog.StartTime,
		EndTime:    dbLog.EndTime,
	}
}

func databaseActivityToActivity(dbAccs []database.GetActivitiesRow) []Activity {
	activities := []Activity{}
	for _, dbAcc := range dbAccs {
		activity := Activity{ActivityID: dbAcc.ActivityID, Name: dbAcc.Name, Points: dbAcc.Points, Type: dbAcc.Type}
		activities = append(activities, activity)
	}
	return activities
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
