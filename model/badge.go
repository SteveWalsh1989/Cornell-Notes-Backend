package models

import (
	"time"
)

// Badge : Sample struct for Badge to test router
type Badge struct {
	ID          string    `db:"id" json:"ID"`
	Name        string    `db:"name" json:"Name"`
	Status      string    `db:"status" json:"Status"`
	DateCreated time.Time `db:"date_created" json:"DateCreated"`
	DateEdited  time.Time `db:"date_edited" json:"DateEdited"`
}

// Folders : list of folders
var Badges []Badge
