package model

import "time"

type User struct {
	ID        string    `db:"id" json:"id"`
	CreatedAt time.Time `db:"CreatedAt" json:"created_at"`
	UpdatedAt time.Time `db:"UpdatedAt" json:"updated_at"`
	Name      string    `dg:"Name" json:"name"`
}

type Users struct {
	Users []User `json:"users"`
}
