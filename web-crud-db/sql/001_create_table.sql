-- +goose Up

CREATE TABLE record(
	ID VARCHAR(36) NOT NULL,
	title VARCHAR(100) NOT NULL,
	description VARCHAR(255)
);

INSERT INTO record (id, title, description) VALUES (uuid(), "Title 1", "Description 1");

-- +goose Down

DROP TABLE record;
