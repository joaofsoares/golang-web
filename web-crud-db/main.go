package main

import (
	"fmt"
	"log"
	"net/http"
	"web-crud-db/api"
	"web-crud-db/render"
)

func main() {
	fmt.Println("Starting server, port :8080")

	http.HandleFunc("/", render.RenderHome)
	http.HandleFunc("/records/", render.RenderRecords)
	http.HandleFunc("/record/new/", render.RenderNewRecord)
	http.HandleFunc("/record/edit/", render.RenderEditRecord)
	http.HandleFunc("/record/save/", render.SaveRecord)
	http.HandleFunc("/record/delete/", render.DeleteRecord)

	http.HandleFunc("/api/records", api.RetrieveAllRecords)
	http.HandleFunc("/api/record/", api.RetrieveRecord)
	http.HandleFunc("/api/record/new", api.NewRecord)
	http.HandleFunc("/api/record/update", api.UpdateRecord)
	http.HandleFunc("/api/record/delete/", api.DeleteRecord)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
