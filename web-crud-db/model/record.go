package model

import "github.com/google/uuid"

type Record struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

type Records struct {
	Records []Record `json:"records"`
}
