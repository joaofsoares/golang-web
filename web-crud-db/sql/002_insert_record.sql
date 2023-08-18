-- +goose Up

INSERT INTO record (id, title, description) VALUES (uuid(), "Title 1", "Description 1");

-- +goose Down

DELETE FROM record;
