-- name: CreateUser :one
INSERT INTO _users(
    _name,
    _email,
    _phone_number
)
VALUES ($1,$2,$3)
RETURNING *;

-- name: CreateUserVault :one
INSERT INTO _vault(
    _user_id,
    _hashed_pw
)
VALUES ($1,$2)
RETURNING *;

-- name: GetUserByEmail :one
SELECT u.*, v._hashed_pw AS _password FROM _users AS u
JOIN _vault AS v ON v._user_id = u._id 
WHERE u._email = $1;

-- name: GetUserByID :one
SELECT u.*, v._hashed_pw AS _password FROM _users AS u
JOIN _vault AS v ON v._user_id = u._id 
WHERE u._id = $1;

-- name: DeleteUser :exec
DELETE FROM _users
WHERE _users._id = $1;
