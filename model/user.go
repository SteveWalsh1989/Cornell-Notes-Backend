package models

import (
	"time"
)

// User : Stores user information
type User struct {
	ID          string    `db:"id" json:"ID"`
	Name        string    `db:"name" json:"Name"`
	Email       string    `db:"email" json:"Email"`
	Status      string    `db:"status" json:"Status"`
	Password    string    `db:"password" json:"Password"`
	DateCreated time.Time `db:"date_created" json:"DateCreated"`
	DateEdited  time.Time `db:"date_edited" json:"DateEdited"`
}

// Users : list of Users
var Users []User
