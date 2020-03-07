package models

import (
	"time"
)

// CornellNote : Sample struct for CornellNote
type CornellNote struct {
	ID          string       `db:"id" json:"ID"`
	Title       string       `db:"title" json:"Title"`
	Status      string       `db:"status" json:"Status"`
	DateCreated time.Time    `db:"date_created" json:"DateCreated"`
	DateEdited  time.Time    `db:"date_edited" json:"DateEdited"`
	Cues        []CornellCue `json:"Cues"`
	Tags        []Tag        `json:"Tags"`
}

// CornellNotes : list of CornellNotes
var CornellNotes []CornellNote
