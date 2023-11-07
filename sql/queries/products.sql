-- name: GetAllProducts :many
SELECT * 
FROM _products;

-- name: GetUserProducts :many
SELECT * 
FROM _products
WHERE _user_id = $1;

-- name: CreateProduct :one
INSERT INTO _products(
    _user_id,
    _name,
    _amount,
    _carbohydrate, 
    _protein,
    _fat
)
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING *;

-- name: UpdateProduct :exec
UPDATE _products
SET 
    _name = $2,
    _amount = $3,
    _carbohydrate = $4,
    _protein = $5,
    _fat = $6,
    _updated_at = NOW()
WHERE _id = $1;

-- name: DeleteProduct :exec
DELETE FROM _products
WHERE _id = $1;