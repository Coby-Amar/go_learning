// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: products.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO _products(
    _user_id,
    _name,
    _amount,
    _carbohydrate, 
    _protein,
    _fat
)
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING _id, _created_at, _updated_at, _name, _amount, _carbohydrate, _protein, _fat, _user_id
`

type CreateProductParams struct {
	UserID       pgtype.UUID
	Name         string `json:"name" validate:"required,min=4,max=200"`
	Amount       int16  `json:"amount" validate:"required,min=1"`
	Carbohydrate int16  `json:"carbohydrate"`
	Protein      int16  `json:"protein"`
	Fat          int16  `json:"fat"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, createProduct,
		arg.UserID,
		arg.Name,
		arg.Amount,
		arg.Carbohydrate,
		arg.Protein,
		arg.Fat,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Amount,
		&i.Carbohydrate,
		&i.Protein,
		&i.Fat,
		&i.UserID,
	)
	return i, err
}

const getAllProducts = `-- name: GetAllProducts :many
SELECT _id, _created_at, _updated_at, _name, _amount, _carbohydrate, _protein, _fat, _user_id FROM _products
`

func (q *Queries) GetAllProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.db.Query(ctx, getAllProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Amount,
			&i.Carbohydrate,
			&i.Protein,
			&i.Fat,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduct = `-- name: UpdateProduct :one
UPDATE _products
SET 
    _name = $2,
    _amount = $3,
    _carbohydrate = $4,
    _protein = $5,
    _fat = $6,
    _updated_at = NOW()
WHERE _products._id = $1
RETURNING _id, _created_at, _updated_at, _name, _amount, _carbohydrate, _protein, _fat, _user_id
`

type UpdateProductParams struct {
	ID           pgtype.UUID `json:"id"`
	Name         string      `json:"name" validate:"required,min=4,max=200"`
	Amount       int16       `json:"amount" validate:"required,min=1"`
	Carbohydrate int16       `json:"carbohydrate"`
	Protein      int16       `json:"protein"`
	Fat          int16       `json:"fat"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, updateProduct,
		arg.ID,
		arg.Name,
		arg.Amount,
		arg.Carbohydrate,
		arg.Protein,
		arg.Fat,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Amount,
		&i.Carbohydrate,
		&i.Protein,
		&i.Fat,
		&i.UserID,
	)
	return i, err
}