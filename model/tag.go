package models

import (
	"time"
)

// Tag : struct for Tag
type Tag struct {
	ID          string    `db:"id" json:"ID"`
	Title       string    `db:"Title" json:"Title"`
	Color       string    `db:"color" json:"Color"`
	Status      string    `db:"status" json:"Status"`
	DateCreated time.Time `db:"date_created" json:"DateCreated"`
	DateEdited  time.Time `db:"date_edited" json:"DateEdited"`
}

//Tags : list of tags
var Tags []Tag
