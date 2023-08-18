-- +goose Up

CREATE TABLE record(
	ID VARCHAR(36) NOT NULL,
	title VARCHAR(100) NOT NULL,
	description VARCHAR(255),
	PRIMARY KEY (ID)
);

-- +goose Down

DROP TABLE record;
