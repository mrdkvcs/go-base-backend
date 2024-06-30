// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: activity_reports.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const getBestProductivityDay = `-- name: GetBestProductivityDay :one

SELECT DATE(logged_at) AS date, COALESCE(SUM(points ) , 0) AS total_points
FROM activity_logs
WHERE user_id = $1 AND logged_at >= $2 AND logged_at < $3
GROUP BY DATE(logged_at)
ORDER BY total_points DESC
LIMIT 1
`

type GetBestProductivityDayParams struct {
	UserID     uuid.UUID
	LoggedAt   time.Time
	LoggedAt_2 time.Time
}

type GetBestProductivityDayRow struct {
	Date        time.Time
	TotalPoints interface{}
}

func (q *Queries) GetBestProductivityDay(ctx context.Context, arg GetBestProductivityDayParams) (GetBestProductivityDayRow, error) {
	row := q.db.QueryRowContext(ctx, getBestProductivityDay, arg.UserID, arg.LoggedAt, arg.LoggedAt_2)
	var i GetBestProductivityDayRow
	err := row.Scan(&i.Date, &i.TotalPoints)
	return i, err
}

const getProductivityDays = `-- name: GetProductivityDays :many

SELECT DATE(logged_at) AS date , COALESCE(SUM(points ) , 0) AS total_points
FROM activity_logs
WHERE user_id = $1 AND logged_at >= $2 AND logged_at < $3
GROUP BY DATE(logged_at)
ORDER BY DATE(logged_at)
`

type GetProductivityDaysParams struct {
	UserID     uuid.UUID
	LoggedAt   time.Time
	LoggedAt_2 time.Time
}

type GetProductivityDaysRow struct {
	Date        time.Time
	TotalPoints interface{}
}

func (q *Queries) GetProductivityDays(ctx context.Context, arg GetProductivityDaysParams) ([]GetProductivityDaysRow, error) {
	rows, err := q.db.QueryContext(ctx, getProductivityDays, arg.UserID, arg.LoggedAt, arg.LoggedAt_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetProductivityDaysRow
	for rows.Next() {
		var i GetProductivityDaysRow
		if err := rows.Scan(&i.Date, &i.TotalPoints); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTotalAndAverageProductivityPoints = `-- name: GetTotalAndAverageProductivityPoints :one
WITH points_per_day AS ( SELECT DATE(logged_at) AS date, 
        COALESCE(SUM(points), 0) AS total_points
    FROM activity_logs
    WHERE user_id = $1 AND logged_at >= $2 AND logged_at < $3
    GROUP BY DATE(logged_at)
)
SELECT 
    ROUND(COALESCE(SUM(total_points), 0) , 0) AS total_points,
    ROUND(COALESCE(AVG(total_points), 0), 2) AS average_points_per_day
FROM points_per_day
`

type GetTotalAndAverageProductivityPointsParams struct {
	UserID     uuid.UUID
	LoggedAt   time.Time
	LoggedAt_2 time.Time
}

type GetTotalAndAverageProductivityPointsRow struct {
	TotalPoints         string
	AveragePointsPerDay string
}

func (q *Queries) GetTotalAndAverageProductivityPoints(ctx context.Context, arg GetTotalAndAverageProductivityPointsParams) (GetTotalAndAverageProductivityPointsRow, error) {
	row := q.db.QueryRowContext(ctx, getTotalAndAverageProductivityPoints, arg.UserID, arg.LoggedAt, arg.LoggedAt_2)
	var i GetTotalAndAverageProductivityPointsRow
	err := row.Scan(&i.TotalPoints, &i.AveragePointsPerDay)
	return i, err
}