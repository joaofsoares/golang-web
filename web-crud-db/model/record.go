package model

type Record struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Records struct {
	Records []Record `json:"records"`
}
