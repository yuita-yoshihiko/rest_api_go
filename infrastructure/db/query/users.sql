-- name: CreateUser :one
INSERT INTO users (
    name
) VALUES (
    $1
) RETURNING *;

-- name: GetUserByID :one
SELECT
    *
FROM
    users
WHERE
    id = $1;
