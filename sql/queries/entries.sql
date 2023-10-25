-- name: GetReportEntries :many
SELECT _report_entries.*, _products._name FROM _report_entries 
JOIN _reports ON _reports._id =_report_entries._report_id
JOIN _products ON _products._id =_report_entries._product_id
WHERE _reports._id = $1;

-- name: CreateReportEntries :copyfrom
INSERT INTO _report_entries(
    _product_id,
    _report_id,
    _amount,
    _carbohydrates,
    _proteins,
    _fats
)
VALUES ($1,$2,$3,$4,$5,$6);

-- name: UpdateReportEntry :one
UPDATE _report_entries
SET 
    _amount = $2,
    _carbohydrates = $3,
    _proteins = $4,
    _fats = $5,
    _updated_at = NOW()
WHERE _report_entries._id = $1
RETURNING *;

-- name: DeleteReportEntry :exec
DELETE FROM _report_entries
WHERE _report_entries._id = $1;