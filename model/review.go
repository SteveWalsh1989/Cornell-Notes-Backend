package models

import (
	"time"
)

// Review : Sample struct for folder to test router
type Review struct {
	ID          string    `db:"id" json:"ID"`
	Name        string    `db:"name" json:"Name"`
	Status      string    `db:"status" json:"Status"`
	DateCreated time.Time `db:"date_created" json:"DateCreated"`
	DateEdited  time.Time `db:"date_edited" json:"DateEdited"`
	//Cues        CornellCues `db:"Cues" json:"Cues"`
}

// Reviews : list of Reviews
var Reviews []Review
