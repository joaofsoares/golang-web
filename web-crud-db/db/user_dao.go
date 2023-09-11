package db

import (
	"fmt"
	"log"
	"web-crud-db/model"
)

func GetAllUsers() (*model.Users, error) {

	const sql = "SELECT id,created_at,updated_at,name FROM crud.user;"

	rows, err := conn.Query(sql)
	errorHandler(err)
	defer rows.Close()

	var users []model.User

	for rows.Next() {

		var u model.User

		if err := rows.Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt, &u.Name); err != nil {
			return nil, fmt.Errorf("couldnt get users from database: %v", err)
		}

		users = append(users, u)
	}

	return &model.Users{Users: users}, nil
}

func InsertUser(userName string) error {

	stmt, err := conn.Prepare("INSERT INTO crud.user (id, created_at, updated_at, name, api_key) VALUES (uuid(), now(), now(), ?, (SHA2(RANDOM_BYTES(10), 256)))")
	errorHandler(err)
	defer stmt.Close()

	inserted, err := stmt.Exec(userName)
	errorHandler(err)

	lastId, err := inserted.LastInsertId()
	errorHandler(err)

	log.Printf("User %v inserted with id %v", userName, lastId)

	return nil
}
