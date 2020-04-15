package models

type UserStats struct {
	UserID              string `db:"user_id" json:"UserID"`
	Points              int    `db:"points" json:"Points"`
	NotesCreated        int    `db:"notes_created" json:"NotesCreated"`
	CornellNotesCreated int    `db:"cornell_notes_created" json:"CornellNotesCreated"`
	CuesCreated         int    `db:"cues_created" json:"CuesCreated"`
	NotesShared         int    `db:"notes_shared" json:"StNotesSharedatus"`
	CornellNotesShared  int    `db:"cornell_notes_shared" json:"CornellNotesShared"`
	ReviewsCompleted    int    `db:"reviews_completed" json:"ReviewsCompleted"`
	CuesReviewed        int    `db:"cues_reviewed" json:"CuesReviewed"`
}
