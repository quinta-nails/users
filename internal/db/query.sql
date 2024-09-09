-- name: GetByTelegramID :one
SELECT *
FROM users
WHERE telegram_id = $1;

--name: GetAll :many
SELECT *
FROM users;