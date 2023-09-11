package db

import (
	"fmt"
	"log"
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

	row := conn.QueryRow("SELECT id,title,description FROM crud.record WHERE id = ?", id)

	var record model.Record

	if err := row.Scan(&record.ID, &record.Title, &record.Description); err != nil {
		return nil, fmt.Errorf("couldnt get records from database: %v", err)
	}

	return &record, nil
}

func InsertRecord(title string, description string) (*model.Record, error) {

	stmt, err := conn.Prepare("INSERT INTO crud.record (id, title, description) VALUES (uuid(), ?,?)")
	errorHandler(err)
	defer stmt.Close()

	if _, err := stmt.Exec(title, description); err != nil {
		return nil, fmt.Errorf("coulnt insert record: %v", err)
	}

	return &model.Record{Title: title, Description: description}, nil
}

func UpdateRecord(id uuid.UUID, title string, description string) (*model.Record, error) {

	stmt, err := conn.Prepare("UPDATE crud.record SET title=?, description=? WHERE id=?")
	errorHandler(err)
	defer stmt.Close()

	updated, err := stmt.Exec(title, description, id)
	if err != nil {
		return nil, fmt.Errorf("couldnt update record = %v", id)
	}

	nRows, err := updated.RowsAffected()
	errorHandler(err)
	log.Printf("Record updated successfully: %v records updated", nRows)

	return &model.Record{ID: id, Title: title, Description: description}, nil
}

func DeleteRecord(uuid uuid.UUID) (bool, error) {

	stmt, err := conn.Prepare("DELETE FROM crud.record WHERE id=?")
	errorHandler(err)
	defer stmt.Close()

	deleted, err := stmt.Exec(uuid.String())
	if err != nil {
		return false, fmt.Errorf("coulnt delete record = %s", uuid)
	}

	nRows, err := deleted.RowsAffected()
	errorHandler(err)
	log.Printf("Record deleted successfully: %v records deleted", nRows)

	return true, nil
}
