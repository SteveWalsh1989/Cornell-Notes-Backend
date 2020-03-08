package models

import (
	"time"
)

// Note : Sample struct for Note to test router
type Note struct {
	ID          string    `db:"id" json:"ID"`
	FolderID    string    `db:"folder_id" json:"FolderID"`
	Title       string    `db:"title" json:"Title"`
	Body        string    `db:"Body" json:"Body"`
	DateCreated time.Time `db:"date_created" json:"DateCreated"`
	DateEdited  time.Time `db:"date_edited" json:"DateEdited"`
	Tags        []Tag     `json:"Tags"`
}

// Notes : list of Notes
var Notes []Note
