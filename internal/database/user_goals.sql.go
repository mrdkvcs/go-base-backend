// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: user_goals.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const setProductivityGoal = `-- name: SetProductivityGoal :exec

INSERT INTO user_goals (user_id , goal_date , goal_points) VALUES ($1, $2, $3)
`

type SetProductivityGoalParams struct {
	UserID     uuid.UUID
	GoalDate   time.Time
	GoalPoints int32
}

func (q *Queries) SetProductivityGoal(ctx context.Context, arg SetProductivityGoalParams) error {
	_, err := q.db.ExecContext(ctx, setProductivityGoal, arg.UserID, arg.GoalDate, arg.GoalPoints)
	return err
}
