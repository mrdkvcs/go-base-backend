-- name: GetUserByApikey :one
SELECT * FROM users WHERE api_key = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: CreateUser :one
INSERT INTO users (id , created_at , updated_at , username , email , password_hash , api_key)
VALUES ($1, $2, $3, $4 , $5 , $6 , encode(sha256(random()::text::bytea) , 'hex'))
RETURNING *;


