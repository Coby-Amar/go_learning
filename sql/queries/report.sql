-- name: CreateReport :one
INSERT INTO _reports(_date)
VALUES ($1)
RETURNING *;

-- name: GetAllReports :many
SELECT * FROM _reports;