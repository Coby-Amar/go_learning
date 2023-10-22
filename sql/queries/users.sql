-- name: CreateUser :one
INSERT INTO _users(
    _name,
    _email,
    _phone_number
)
VALUES ($1,$2,$3)
RETURNING *;

-- name: UpdateUser :one
UPDATE _users
SET
    _name = $2,
    _email = $3,
    _phone_number = $4,
    _updated_at = NOW()
WHERE _users._id = $1
RETURNING *;

-- name: CreateUserVault :one
INSERT INTO _vault(
    _user_id,
    _hashed_pw
)
VALUES ($1,$2)
RETURNING *;

-- name: UpdateUserVaultByID :one
UPDATE _vault
SET
    _hashed_pw = $2
WHERE _vault._user_id = $1
RETURNING *;

-- name: GetUserByID :one
SELECT u.*, v._hashed_pw AS _password FROM _users AS u
JOIN _vault AS v ON v._user_id = u._id 
WHERE u._id = $1;

-- name: DeleteUser :exec
DELETE FROM _users
WHERE _users._id = $1;
