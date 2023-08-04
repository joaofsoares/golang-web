package model

type Record struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Records struct {
	Records []Record
}
