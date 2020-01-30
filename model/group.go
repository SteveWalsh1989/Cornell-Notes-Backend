package models

import (
	"time"
)

// Group : Sample struct for Group
type Group struct {
	ID          string    `db:"id" json:"ID"`
	Name        string    `db:"name" json:"Name"`
	Status      string    `db:"status" json:"Status"`
	DateCreated time.Time `db:"date_created" json:"DateCreated"`
	DateEdited  time.Time `db:"date_edited" json:"DateEdited"`
}

// Groups : list of Group
var Groups []Group
