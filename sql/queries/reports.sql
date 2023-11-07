-- name: GetAllUserReports :many
SELECT * FROM _reports
WHERE _user_id = $1;

-- name: GetReportByID :one
SELECT * FROM _reports
WHERE _id = $1;

-- name: CreateReport :one
INSERT INTO _reports(
    _date,
    _amout_of_entries,
    _carbohydrates_total,
    _proteins_total,
    _fats_total,
    _user_id
)
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING *;

-- name: UpdateReport :one
UPDATE _reports
SET
    _amout_of_entries = $2,
    _carbohydrates_total = $3,
    _proteins_total = $4,
    _fats_total = $5,
    _updated_at = NOW()
WHERE _reports._id = $1
RETURNING *;

-- name: DeleteReport :exec
DELETE FROM _reports
WHERE _reports._id = $1;
