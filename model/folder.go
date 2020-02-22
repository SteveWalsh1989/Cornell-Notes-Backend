package models

import (
	"time"
)

// Folder : struct for folder
type Folder struct {
	ID          string       `db:"id" json:"ID"`
	Title       string       `db:"title" json:"Title"`
	DateCreated time.Time    `db:"date_created" json:"Date_Created"`
	DateEdited  time.Time    `db:"date_edited" json:"Date_Edited"`
	Items       []FolderItem `json:"Folder_Items"`
}

// FolderItem : stores item in folders - notes, cornell notes, pdf
type FolderItem struct {
	Title     string `db:"Title" json:"Title"`
	ID        string `db:"ID" json:"ID"`
	ItemType  string `db:"item_type" json:"Item_Type"`
	ItemID    string `db:"item_ID" json:"Item_ID"`
	ItemTitle string `db:"itemTitle" json:"Item_Title"`
}

// Folders : list of folders
var Folders []Folder
