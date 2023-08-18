-- +goose Up

ALTER TABLE user ADD COLUMN api_key VARCHAR(64) NOT NULL;
ALTER TABLE user ADD UNIQUE (api_key);

-- +goose Down

ALTER TABLE user DROP COLUMN api_key;
