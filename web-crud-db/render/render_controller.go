package render

import (
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"web-crud-db/db"
)

var templates = template.Must(template.ParseFiles(
	"tmpl/home.html",
	"tmpl/records.html",
	"tmpl/newRecord.html",
	"tmpl/editRecord.html",
))

var validPath = regexp.MustCompile("^/record/(edit|save|delete)/([0-9]+)$")

func RenderHome(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "home.html", nil)
}

func RenderRecords(w http.ResponseWriter, r *http.Request) {

	records, err := db.GetAllRecords()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templates.ExecuteTemplate(w, "records.html", records)
}

func RenderEditRecord(w http.ResponseWriter, r *http.Request) {

	path := validPath.FindStringSubmatch(r.URL.Path)
	id, err := strconv.Atoi(path[2])

	if err != nil {
		http.Redirect(w, r, "/records/", http.StatusFound)
	}

	record, err := db.GetRecordById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templates.ExecuteTemplate(w, "editRecord.html", record)
}

func RenderNewRecord(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "newRecord.html", nil)
}

func SaveRecord(w http.ResponseWriter, r *http.Request) {
	path := validPath.FindStringSubmatch(r.URL.Path)

	if path != nil {
		id, err := strconv.Atoi(path[2])
		title := r.FormValue("title")
		description := r.FormValue("description")

		if err != nil {
			http.Redirect(w, r, "/records/", http.StatusFound)
		}

		updated, err := db.UpdateRecord(id, title, description)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if updated != nil {
			log.Printf("Record Id = %d updated successfully", id)
		}
	} else {
		title := r.FormValue("title")
		description := r.FormValue("description")

		record, err := db.InsertRecord(title, description)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if record != nil {
			log.Printf("Record inserted successfully")
		}
	}

	http.Redirect(w, r, "/records/", http.StatusFound)
}

func DeleteRecord(w http.ResponseWriter, r *http.Request) {
	path := validPath.FindStringSubmatch(r.URL.Path)
	id, err := strconv.Atoi(path[2])

	if err != nil {
		http.Redirect(w, r, "/records/", http.StatusFound)
	}

	deleted, err := db.DeleteRecord(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if deleted {
		log.Printf("Record Id = %d deleted successfully", id)
	}

	http.Redirect(w, r, "/records/", http.StatusFound)
}
