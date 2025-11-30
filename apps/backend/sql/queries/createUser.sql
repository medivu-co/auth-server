-- name: CreateUser :exec
INSERT INTO users (email, password_hash, name)
VALUES ($1, $2, $3);