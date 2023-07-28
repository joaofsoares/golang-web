package api

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"web-crud-db/db"
	"web-crud-db/model"
)

var recordPath = regexp.MustCompile("^/api/record/([0-9]+)$")
var deletePath = regexp.MustCompile("^/api/record/delete/([0-9]+)$")

func RetrieveAllRecords(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method is not GET", http.StatusInternalServerError)
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
		http.Error(w, "Method must be GET", http.StatusInternalServerError)
		return
	}

	path := recordPath.FindStringSubmatch(r.URL.Path)
	id, err := strconv.Atoi(path[1])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	record, err := db.GetRecordById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(record)
}

func NewRecord(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Method must be POST", http.StatusInternalServerError)
		return
	}

	var rec model.Record

	if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newRecord, err := db.InsertRecord(rec.Title, rec.Description)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newRecord)
}

func UpdateRecord(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		http.Error(w, "Method must be PUT", http.StatusInternalServerError)
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
		http.Error(w, "Method must be DELETE", http.StatusInternalServerError)
		return
	}

	path := deletePath.FindStringSubmatch(r.URL.Path)

	id, err := strconv.Atoi(path[1])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	deleted, err := db.DeleteRecord(id)

	if err != nil && !deleted {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
