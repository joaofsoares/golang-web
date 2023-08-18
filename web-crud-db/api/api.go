package api

import (
	"encoding/json"
	"net/http"
	"regexp"
	"web-crud-db/db"
	"web-crud-db/model"

	"github.com/google/uuid"
)

var recordPath = regexp.MustCompile("^/api/record/([0-9]+)$")
var deletePath = regexp.MustCompile("^/api/record/delete/([0-9]+)$")

func RetrieveAllRecords(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method is not GET", http.StatusMethodNotAllowed)
		return
	}

	records, err := db.GetAllRecords()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(records)
}

func RetrieveRecord(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method must be GET", http.StatusMethodNotAllowed)
		return
	}

	path := recordPath.FindStringSubmatch(r.URL.Path)

	record, err := db.GetRecordById(uuid.MustParse(path[1]))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(record)
}

func NewRecord(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Method must be POST", http.StatusMethodNotAllowed)
		return
	}

	var rec model.Record

	if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	insertedRecord, err := db.InsertRecord(rec.Title, rec.Description)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newRecord := struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}{
		Title:       insertedRecord.Title,
		Description: insertedRecord.Description,
	}

	json.NewEncoder(w).Encode(newRecord)
}

func UpdateRecord(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		http.Error(w, "Method must be PUT", http.StatusMethodNotAllowed)
		return
	}

	var rec model.Record

	if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updateRecord, err := db.UpdateRecord(rec.ID, rec.Title, rec.Description)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updateRecord)

}

func DeleteRecord(w http.ResponseWriter, r *http.Request) {

	if r.Method != "DELETE" {
		http.Error(w, "Method must be DELETE", http.StatusMethodNotAllowed)
		return
	}

	path := deletePath.FindStringSubmatch(r.URL.Path)

	deleted, err := db.DeleteRecord(uuid.MustParse(path[1]))

	if err != nil && !deleted {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func HandleUsers(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" && r.Method != "POST" {
		http.Error(w, "Method must be GET", http.StatusMethodNotAllowed)
	}

	if r.Method == "GET" {
		users, err := db.GetAllUsers()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(users)
	} else {

		type param struct {
			Name string `json:"name"`
		}

		p := param{}

		err := json.NewDecoder(r.Body).Decode(&p)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = db.InsertUser(p.Name)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
