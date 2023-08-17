package db

import (
	"database/sql"
	"fmt"
	"web-crud-db/model"
)

var conn *sql.DB

func GetAllRecords() (*model.Records, error) {

	if conn == nil {
		conn = createConnection()
	}

	const sql = "SELECT id,title,description FROM crud.record;"

	rows, err := conn.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var records []model.Record

	for rows.Next() {

		var r model.Record

		if err := rows.Scan(&r.ID, &r.Title, &r.Description); err != nil {
			return nil, fmt.Errorf("couldnt get records from database")
		}

		records = append(records, r)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("something else failed")
	}

	return &model.Records{Records: records}, nil
}

func GetRecordById(id string) (*model.Record, error) {

	if conn == nil {
		conn = createConnection()
	}

	rows, err := conn.Query("SELECT id,title,description FROM crud.record WHERE id = ?", id)

	if err != nil {
		return nil, fmt.Errorf("couldnt select to database")
	}

	defer rows.Close()

	var record model.Record

	for rows.Next() {
		if err := rows.Scan(&record.ID, &record.Title, &record.Description); err != nil {
			return nil, fmt.Errorf("couldnt get records from database")
		}
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("something else failed")
	}

	return &record, nil
}

func InsertRecord(title string, description string) (*model.Record, error) {
	if conn == nil {
		conn = createConnection()
	}

	inserted, err := conn.Query("INSERT INTO crud.record (id, title, description) VALUES (uuid(), ?,?)", title, description)

	if err != nil {
		return nil, fmt.Errorf("coulnt insert record")
	}

	defer inserted.Close()

	return &model.Record{Title: title, Description: description}, nil
}

func UpdateRecord(id string, title string, description string) (*model.Record, error) {
	if conn == nil {
		conn = createConnection()
	}

	updated, err := conn.Query("UPDATE crud.record SET title=?, description=? WHERE id=?", title, description, id)

	if err != nil {
		return nil, fmt.Errorf("couldnt update record = %v", id)
	}

	defer updated.Close()

	return &model.Record{ID: id, Title: title, Description: description}, nil
}

func DeleteRecord(uuid string) (bool, error) {
	if conn == nil {
		conn = createConnection()
	}

	deleted, err := conn.Query("DELETE FROM crud.record WHERE id=?", uuid)

	if err != nil {
		return false, fmt.Errorf("coulnt delete record = %s", uuid)
	}

	defer deleted.Close()

	return true, nil
}
