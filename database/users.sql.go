// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: users.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO _users(
    _name,
    _email,
    _phone_number
)
VALUES ($1,$2,$3)
RETURNING _id, _created_at, _updated_at, _last_login, _name, _email, _phone_number
`

type CreateUserParams struct {
	Name        string
	Email       string
	PhoneNumber string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Name, arg.Email, arg.PhoneNumber)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastLogin,
		&i.Name,
		&i.Email,
		&i.PhoneNumber,
	)
	return i, err
}

const createUserVault = `-- name: CreateUserVault :one
INSERT INTO _vault(
    _user_id,
    _hashed_pw
)
VALUES ($1,$2)
RETURNING _user_id, _hashed_pw, _active
`

type CreateUserVaultParams struct {
	UserID   pgtype.UUID
	HashedPw string
}

func (q *Queries) CreateUserVault(ctx context.Context, arg CreateUserVaultParams) (Vault, error) {
	row := q.db.QueryRow(ctx, createUserVault, arg.UserID, arg.HashedPw)
	var i Vault
	err := row.Scan(&i.UserID, &i.HashedPw, &i.Active)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM _users
WHERE _users._id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, ID pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteUser, ID)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT 
    _u._id, _u._created_at, _u._updated_at, _u._last_login, _u._name, _u._email, _u._phone_number, 
    _v._hashed_pw AS _password, 
    _v._active AS _active 
FROM _users AS _u
JOIN _vault AS _v ON _v._user_id = _u._id 
WHERE _u._email = $1
`

type GetUserByEmailRow struct {
	ID          pgtype.UUID
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
	LastLogin   pgtype.Timestamp
	Name        string
	Email       string
	PhoneNumber string
	Password    string
	Active      bool
}

func (q *Queries) GetUserByEmail(ctx context.Context, Email string) (GetUserByEmailRow, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, Email)
	var i GetUserByEmailRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastLogin,
		&i.Name,
		&i.Email,
		&i.PhoneNumber,
		&i.Password,
		&i.Active,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT 
    _u._id, _u._created_at, _u._updated_at, _u._last_login, _u._name, _u._email, _u._phone_number, 
    _v._hashed_pw AS _password, 
    _v._active AS _active 
FROM _users AS _u
JOIN _vault AS _v ON _v._user_id = _u._id 
WHERE _u._id = $1
`

type GetUserByIDRow struct {
	ID          pgtype.UUID
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
	LastLogin   pgtype.Timestamp
	Name        string
	Email       string
	PhoneNumber string
	Password    string
	Active      bool
}

func (q *Queries) GetUserByID(ctx context.Context, ID pgtype.UUID) (GetUserByIDRow, error) {
	row := q.db.QueryRow(ctx, getUserByID, ID)
	var i GetUserByIDRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastLogin,
		&i.Name,
		&i.Email,
		&i.PhoneNumber,
		&i.Password,
		&i.Active,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE _users
SET
    _name = $2,
    _email = $3,
    _phone_number = $4,
    _updated_at = NOW()
WHERE _users._id = $1
RETURNING _id, _created_at, _updated_at, _last_login, _name, _email, _phone_number
`

type UpdateUserParams struct {
	ID          pgtype.UUID
	Name        string
	Email       string
	PhoneNumber string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.PhoneNumber,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastLogin,
		&i.Name,
		&i.Email,
		&i.PhoneNumber,
	)
	return i, err
}

const updateUserVaultByID = `-- name: UpdateUserVaultByID :one
UPDATE _vault
SET
    _hashed_pw = $2,
    _active = $3
WHERE _vault._user_id = $1
RETURNING _user_id, _hashed_pw, _active
`

type UpdateUserVaultByIDParams struct {
	UserID   pgtype.UUID
	HashedPw string
	Active   bool
}

func (q *Queries) UpdateUserVaultByID(ctx context.Context, arg UpdateUserVaultByIDParams) (Vault, error) {
	row := q.db.QueryRow(ctx, updateUserVaultByID, arg.UserID, arg.HashedPw, arg.Active)
	var i Vault
	err := row.Scan(&i.UserID, &i.HashedPw, &i.Active)
	return i, err
}
