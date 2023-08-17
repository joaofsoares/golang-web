package render

import (
	"html/template"
	"log"
	"net/http"
	"regexp"
	"web-crud-db/db"
)

var templates = template.Must(template.ParseFiles(
	"tmpl/home.html",
	"tmpl/records.html",
	"tmpl/newRecord.html",
	"tmpl/editRecord.html",
))

var validPath = regexp.MustCompile("^/record/(edit|save|delete)/([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})$")

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

	record, err := db.GetRecordById(path[2])

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
		title := r.FormValue("title")
		description := r.FormValue("description")

		updated, err := db.UpdateRecord(path[2], title, description)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if updated != nil {
			log.Printf("Record Id = %s updated successfully", path[2])
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

	deleted, err := db.DeleteRecord(path[2])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if deleted {
		log.Printf("Record Id = %s deleted successfully", path[2])
	}

	http.Redirect(w, r, "/records/", http.StatusFound)
}
