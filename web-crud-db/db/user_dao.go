package db

import (
	"fmt"
	"web-crud-db/model"
)

func GetAllUsers() (*model.Users, error) {

	createConnection()

	const sql = "SELECT id,created_at,updated_at,name FROM crud.user;"

	rows, err := conn.Query(sql)

	if err != nil {
		return nil, err
	}

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

	createConnection()

	inserted, err := conn.Query("INSERT INTO crud.user (id, created_at, updated_at, name, api_key) VALUES (uuid(), now(), now(), ?, (SHA2(RANDOM_BYTES(10), 256)))", userName)

	if err != nil {
		return fmt.Errorf("coulnt insert record: %v", err)
	}

	defer inserted.Close()

	return nil
}
