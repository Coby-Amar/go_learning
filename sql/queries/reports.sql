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

-- name: GetAllUserReports :many
SELECT * FROM _reports
WHERE _reports._user_id = $1;