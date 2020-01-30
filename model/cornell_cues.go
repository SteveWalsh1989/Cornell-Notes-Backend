package models

import (
	"time"
)

// CornellCue : Sample struct for folder to test router
type CornellCue struct {
	ID          string    `db:"id" json:"ID"`
	Name        string    `db:"name" json:"Name"`
	Status      string    `db:"status" json:"Status"`
	DateCreated time.Time `db:"date_created" json:"DateCreated"`
	DateEdited  time.Time `db:"date_edited" json:"DateEdited"`
}

// CornellCues : list of folders
var CornellCues []CornellCue
