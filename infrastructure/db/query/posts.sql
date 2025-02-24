-- name: CreatePost :one
INSERT INTO posts (
    user_id, title, content
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetPostByID :one
SELECT
    *
FROM
    posts
WHERE
    id = $1;
