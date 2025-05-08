-- name: GetUser :one
SELECT * FROM "User"
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO "User"(
    name, password
) VALUES (
    $1, $2
) RETURNING *;
