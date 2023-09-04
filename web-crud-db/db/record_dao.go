package db

import (
	"fmt"
	"web-crud-db/model"

	"github.com/google/uuid"
)

func GetAllRecords() (*model.Records, error) {

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
			return nil, fmt.Errorf("couldnt get records from database: %v", err)
		}

		records = append(records, r)
	}

	return &model.Records{Records: records}, nil
}

func GetRecordById(id uuid.UUID) (*model.Record, error) {

	rows, err := conn.Query("SELECT id,title,description FROM crud.record WHERE id = ?", id)

	if err != nil {
		return nil, fmt.Errorf("couldnt select to database: %v", err)
	}

	defer rows.Close()

	var record model.Record

	for rows.Next() {
		if err := rows.Scan(&record.ID, &record.Title, &record.Description); err != nil {
			return nil, fmt.Errorf("couldnt get records from database: %v", err)
		}
	}

	return &record, nil
}

func InsertRecord(title string, description string) (*model.Record, error) {

	inserted, err := conn.Query("INSERT INTO crud.record (id, title, description) VALUES (uuid(), ?,?)", title, description)

	if err != nil {
		return nil, fmt.Errorf("coulnt insert record: %v", err)
	}

	defer inserted.Close()

	return &model.Record{Title: title, Description: description}, nil
}

func UpdateRecord(id uuid.UUID, title string, description string) (*model.Record, error) {

	updated, err := conn.Query("UPDATE crud.record SET title=?, description=? WHERE id=?", title, description, id)

	if err != nil {
		return nil, fmt.Errorf("couldnt update record = %v", id)
	}

	defer updated.Close()

	return &model.Record{ID: id, Title: title, Description: description}, nil
}

func DeleteRecord(uuid uuid.UUID) (bool, error) {

	deleted, err := conn.Query("DELETE FROM crud.record WHERE id=?", uuid.String())

	if err != nil {
		return false, fmt.Errorf("coulnt delete record = %s", uuid)
	}

	defer deleted.Close()

	return true, nil
}
