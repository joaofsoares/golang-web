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

	const sql = "SELECT id,title,description FROM learn.record;"

	rows, err := conn.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var records []model.Record

	for rows.Next() {

		var r model.Record

		if err := rows.Scan(&r.ID, &r.Title, &r.Description); err != nil {
			return nil, fmt.Errorf("Couldnt get records from database")
		}

		records = append(records, r)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Something else failed")
	}

	return &model.Records{Records: records}, nil
}

func GetRecordById(id int) (*model.Record, error) {

	if conn == nil {
		conn = createConnection()
	}

	rows, err := conn.Query("SELECT id,title,description FROM learn.record WHERE id = ?", id)

	if err != nil {
		return nil, fmt.Errorf("Couldnt select to database")
	}

	defer rows.Close()

	var record model.Record

	for rows.Next() {
		if err := rows.Scan(&record.ID, &record.Title, &record.Description); err != nil {
			return nil, fmt.Errorf("Couldnt get records from database")
		}
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Something else failed")
	}

	return &record, nil
}

func InsertRecord(title string, description string) (*model.RecordInsert, error) {
	if conn == nil {
		conn = createConnection()
	}

	inserted, err := conn.Query("INSERT INTO learn.record (title, description) VALUES (?,?)", title, description)

	if err != nil {
		return nil, fmt.Errorf("Coulnt insert record")
	}

	defer inserted.Close()

	return &model.RecordInsert{Title: title, Description: description}, nil
}

func UpdateRecord(id int, title string, description string) (*model.Record, error) {
	if conn == nil {
		conn = createConnection()
	}

	updated, err := conn.Query("UPDATE learn.record SET title=?, description=? WHERE id=?", title, description, id)

	if err != nil {
		return nil, fmt.Errorf("Couldnt update record = %d", id)
	}

	defer updated.Close()

	return &model.Record{ID: id, Title: title, Description: description}, nil
}

func DeleteRecord(id int) (bool, error) {
	if conn == nil {
		conn = createConnection()
	}

	deleted, err := conn.Query("DELETE FROM learn.record WHERE id=?", id)

	if err != nil {
		return false, fmt.Errorf("Coulnt delete record = %d", id)
	}

	defer deleted.Close()

	return true, nil
}
