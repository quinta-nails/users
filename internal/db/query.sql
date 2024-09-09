-- name: GetByTelegramID :one
SELECT *
FROM users
WHERE telegram_id = $1;