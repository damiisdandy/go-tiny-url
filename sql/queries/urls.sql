-- name: CreateURL :one
INSERT INTO urls (url_id, original_url) VALUES ($1, $2) RETURNING *;