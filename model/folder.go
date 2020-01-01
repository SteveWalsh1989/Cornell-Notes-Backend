package models

// Folder : Sample struct for folder to test router
type Folder struct {
	ID    string `json:"id"`
	Name  string `json:"title"`
	Notes string `json:"body"`
}

// Folders : list of folders
var Folders []Folder
