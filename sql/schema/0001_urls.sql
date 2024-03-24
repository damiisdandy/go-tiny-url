-- +goose Up 
CREATE TABLE urls (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ DEFAULT now(),
    url_id TEXT UNIQUE NOT NULL,
    original_url TEXT NOT NULL
);


-- +goose Down
DROP TABLE urls;