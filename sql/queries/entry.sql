-- name: CreateReportEntry :one
INSERT INTO _report_entries(
    _product_id,
    _report_id,
    _amount,
    _carbohydrate,
    _protein,
    _fat
)
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING *;

-- name: UpdateReportEntry :one
UPDATE _report_entries
SET 
    _amount = $2,
    _carbohydrate = $3,
    _protein = $4,
    _fat = $5
WHERE _report_entries._id = $1
RETURNING *;

-- name: GetReportEntries :many
SELECT _report_entries.*, _products._name FROM _report_entries 
JOIN _reports ON _reports._id =_report_entries._report_id
JOIN _products ON _products._id =_report_entries._product_id
WHERE _reports._id = $1;
