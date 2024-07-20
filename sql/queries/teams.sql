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


