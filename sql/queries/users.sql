-- name: GetUserByEmail :one
SELECT 
    _u.*, 
    _v._hashed_pw AS _password, 
    _v._active AS _active 
FROM _users AS _u
JOIN _vault AS _v ON _v._user_id = _u._id 
WHERE _u._email = $1;

-- name: GetUserByID :one
SELECT 
    _u.*, 
    _v._hashed_pw AS _password, 
    _v._active AS _active 
FROM _users AS _u
JOIN _vault AS _v ON _v._user_id = _u._id 
WHERE _u._id = $1;

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
    _hashed_pw = $2,
    _active = $3
WHERE _vault._user_id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM _users
WHERE _users._id = $1;
