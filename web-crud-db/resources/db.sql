CREATE schema learn;

CREATE TABLE record(
	ID INT NOT NULL AUTO_INCREMENT,
	title VARCHAR(100) NOT NULL,
	description VARCHAR(255),
	PRIMARY KEY(ID)
);

INSERT INTO record(title,description) VALUES ("Title1", "Description1");

COMMIT;
