-- +goose Up

INSERT INTO user (id, created_at, updated_at, name) VALUES (uuid(), now(), now(), 'MyAdmin');

-- +goose Down

DELETE FROM user;
