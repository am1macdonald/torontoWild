-- name: AddUser :exec
INSERT INTO users
(email, first_name, last_name)
VALUES ($1, $2, $3);

-- name: GetUser :one
SELECT * FROM users
WHERE email = $1
LIMIT 1; 

-- name: LoginUser :exec
UPDATE users SET 
last_login = now();

-- name: SetUserActive :exec
UPDATE users SET 
is_active = true;
