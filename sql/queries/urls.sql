-- name: CreateURL :one
INSERT INTO urls (url_id, original_url) VALUES ($1, $2) RETURNING *;

-- name: DeleteURL :one
DELETE FROM urls WHERE url_id = $1 RETURNING *;

-- name: GetURL :one
SELECT * FROM urls WHERE url_id = $1;