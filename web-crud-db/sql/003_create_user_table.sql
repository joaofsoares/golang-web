-- +goose Up

CREATE TABLE user(
	ID VARCHAR(36) NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
    name varchar(255) NOT NULL,
	PRIMARY KEY (ID)
);

-- +goose Down

DROP TABLE user;
