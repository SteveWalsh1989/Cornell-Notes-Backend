package models

import (
	"time"
)

// Folder : Sample struct for folder to test router
type User struct {
	ID          string    `db:"id" json:"ID"`
	FirstName   string    `db:"first_name" json:"FirstName"`
	LastName    string    `db:"last_name" json:"LastName"`
	Status      string    `db:"status" json:"Status"`
	Password    string    `db:"password" json:"Password"`
	DateCreated time.Time `db:"date_created" json:"DateCreated"`
	DateEdited  time.Time `db:"date_edited" json:"DateEdited"`
}

// Folders : list of folders
var Users []User
