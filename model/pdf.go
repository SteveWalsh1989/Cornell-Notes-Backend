package models

import (
	"time"
)

// PDF : Sample struct for PDF to test router
type PDF struct {
	ID          string    `db:"id" json:"ID"`
	Name        string    `db:"name" json:"Name"`
	pdfContent  []byte    `db:"pdfContent" json:"pdfContent"`
	Status      string    `db:"status" json:"Status"`
	DateCreated time.Time `db:"date_created" json:"DateCreated"`
	DateEdited  time.Time `db:"date_edited" json:"DateEdited"`
}

// PDFs : list of PDF
var PDFs []PDF
