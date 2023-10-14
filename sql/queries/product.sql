-- name: CreateProduct :one
INSERT INTO _products(
    _name,
    _amount,
    _carbohydrate, 
    _protein,
    _fat
)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;

-- name: UpdateProduct :one
UPDATE _products
SET 
    _name = $2,
    _amount = $3,
    _carbohydrate = $4,
    _protein = $5,
    _fat = $6
WHERE _products._id = $1
RETURNING *;

-- name: GetAllProducts :many
SELECT * FROM _products;
