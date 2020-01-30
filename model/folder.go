package models

import (
	"time"
)

// Folder : Sample struct for folder to test router
type Folder struct {
	ID          string    `json:"ID"`
	Name        string    `json:"Name"`
	ItemID      string    `json:"ItemID"`
	ItemType    string    `json:"ItemType"`
	Status      string    `json:"Status"`
	DateCreated time.Time `json:"DateCreated"`
	DateEdited  time.Time `json:"DateEdited"`
	DateDeleted time.Time `json:"DateDeleted"`
}

// Folders : list of folders
var Folders []Folder
