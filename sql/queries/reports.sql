-- name: CreateReport :one
INSERT INTO _reports(
    _date,
    _amout_of_entries,
    _carbohydrates, 
    _proteins,
    _fats
)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;

-- name: GetAllReports :many
SELECT * FROM _reports;