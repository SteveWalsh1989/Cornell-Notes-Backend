package models

import (
	"time"
)

// Folder : Sample struct for folder to test router
type CornellNote struct {
	ID          string    `db:"id" json:"ID"`
	Name        string    `db:"name" json:"Name"`
	Status      string    `db:"status" json:"Status"`
	DateCreated time.Time `db:"date_created" json:"DateCreated"`
	DateEdited  time.Time `db:"date_edited" json:"DateEdited"`
}

// Folders : list of folders
var CornellNotes []CornellNote
