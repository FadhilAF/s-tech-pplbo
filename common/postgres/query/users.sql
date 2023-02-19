-- name: GetUser :one
SELECT * FROM users
where id = ? LIMIT 1;

-- name: ListUser :many
SELECT id, name, email FROM users
ORDER BY name;

-- name: CreateUser :execresult
INSERT INTO users (
  name, email, password_hash, address, phone
) VALUES ( ?, ?, ?, ?, ? );

-- name: UpdateUser :execresult
UPDATE users SET
  name = ?,
  address = ?,
  phone = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;
