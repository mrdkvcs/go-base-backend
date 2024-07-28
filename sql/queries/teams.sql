-- name: CreateTeam :one
INSERT INTO teams (id , name,  team_industry , team_size , is_private , created_by) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: CreateTeamMembership :exec
INSERT INTO team_memberships (id, team_id, user_id, role) VALUES ($1, $2, $3, $4); 

-- name: GetUserTeams :many

SELECT 
    tm.team_id,
    t.name AS team_name,
    tm.role
FROM 
    team_memberships tm
JOIN 
    teams t ON tm.team_id = t.id
WHERE 
    tm.user_id = $1;

-- name: GetTeamInFo :one

SELECT t.id ,  t.name , t.team_industry , t.team_size , t.is_private  , t.created_by , u.username 
FROM 
teams t
JOIN users u ON u.id = t.created_by
WHERE t.id = $1;

-- name: GetTeamActivities :many
SELECT activity_name, points FROM team_activities WHERE team_id = $1;

