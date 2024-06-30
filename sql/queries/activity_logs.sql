-- name: GetActivities :many
(
    SELECT a.activity_id, a.name, a.points,aa.type 
    FROM activities a
    JOIN all_activities aa ON a.activity_id = aa.id
    WHERE aa.type = 'default'
)
UNION ALL
(
 SELECT  ca.activity_id,ca.name, ca.points, aa.type
    FROM custom_activities ca
    JOIN all_activities aa ON ca.activity_id = aa.id
    WHERE aa.type = 'custom' AND ca.user_id = $1  
)
ORDER BY points DESC;

-- name: SetAllActivity :exec
 
INSERT INTO all_activities (id, type) VALUES ($1 , $2);

-- name: SetCustomActivity :exec
INSERT INTO custom_activities (activity_id , user_id , name , points , created_at) VALUES ($1 , $2 , $3 , $4 , $5);

-- name: DeleteActivity :exec
DELETE FROM all_activities WHERE id = $1;

-- name: SetActivityLog :one
INSERT INTO activity_logs (id , user_id , activity_id , duration , points , logged_at , start_time  , end_time) VALUES ($1 , $2 , $3 , $4 , $5 , $6 , $7  ,$8) RETURNING *;

-- name: GetDailyActivityLogs :many
SELECT
    al.activity_id,
    al.points,
    al.duration,
    al.start_time,
    al.end_time,
    CASE
        WHEN aa.type = 'custom' THEN ca.name
        WHEN aa.type = 'default' THEN a.name
    END AS activity_name
FROM
    activity_logs al
JOIN
    all_activities aa ON al.activity_id = aa.id
LEFT JOIN
    custom_activities ca ON aa.id = ca.activity_id AND aa.type = 'custom'
LEFT JOIN
    activities a ON aa.id = a.activity_id AND aa.type = 'default'
WHERE
    al.user_id = $1
    AND al.logged_at::date = CURRENT_DATE;


-- name: GetDailyActivityPoints :one
SELECT COALESCE(sum(points) , 0)  FROM activity_logs WHERE user_id = $1 AND logged_at::date = CURRENT_DATE; 
