package models

import (
	"time"
)

// CornellCue : Sample struct for folder to test router
type CornellCue struct {
	ID            string    `db:"id" json:"ID"`
	CornellNoteID string    `db:"cornell_note_id" json:"CornellNoteID"`
	Cue           string    `db:"cue" json:"Cue"`
	CueOrder      int       `db:"cue_order" json:"CueOrder"`
	Answer        string    `db:"answer" json:"Answer"`
	Status        string    `db:"status" json:"Status"`
	DateCreated   time.Time `db:"date_created" json:"DateCreated"`
	DateEdited    time.Time `db:"date_edited" json:"DateEdited"`
}

// CornellCues : list of folders
var CornellCues []CornellCue
