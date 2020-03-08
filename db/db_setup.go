package db

/*
 db_setup:
 --------------------------------------------
 file contains Maps containing SQL commands to create
 tables and insert sample data into the application
*/

//CreateTableCommands ... stores all create table commands needed for application - 20 tables
var CreateTables = map[string]string{
	"badges":        "CREATE TABLE badges (id VARCHAR(36) NOT NULL  PRIMARY KEY, title VARCHAR(36), item_id VARCHAR(36), item_type VARCHAR(10), status VARCHAR(10), date_created datetime, date_edited datetime);",
	"badge_users":   "CREATE TABLE badge_users (badge_id VARCHAR(36) NOT NULL, user_id VARCHAR(36) NOT NULL, PRIMARY KEY(badge_id, user_id));",
	"badge_scores":  "CREATE TABLE badge_scores (user_id VARCHAR(36) NOT NULL PRIMARY KEY, notes_created integer, cornell_notes_created integer,reviews_created integer,notes_shared integer, cornell_notes_shared integer,reviews_shared integer, reviews_completed integer);",
	"cornell_notes": "CREATE TABLE cornell_notes (id VARCHAR(36) NOT NULL PRIMARY KEY, title VARCHAR(36), status VARCHAR(10) DEFAULT 'Active', date_created datetime, date_edited datetime);",
	"cornell_cues":  "CREATE TABLE cornell_cues (id VARCHAR(36) NOT NULL PRIMARY KEY, cornell_note_id VARCHAR(36), cue VARCHAR(1000), answer VARCHAR(10000));",
	"cornell_users": "CREATE TABLE cornell_users  (cornell_note_id VARCHAR(36) NOT NULL, user_id VARCHAR(36) NOT NULL, PRIMARY KEY(cornell_note_id, user_id));",
	"folder_items":  "CREATE TABLE folder_items  (folder_id VARCHAR(36), item_id VARCHAR(36),item_type VARCHAR(36),  PRIMARY KEY(folder_id, item_id));",
	"folders":       "CREATE TABLE folders (id VARCHAR(36) NOT NULL PRIMARY KEY, title VARCHAR(36), status VARCHAR(10) DEFAULT 'Active', date_created datetime, date_edited datetime);",
	"folder_users":  "CREATE TABLE folder_users (folder_id VARCHAR(36) NOT NULL, user_id VARCHAR(36) NOT NULL, PRIMARY KEY(folder_id, user_id));",
	"group_users":   "CREATE TABLE group_users (group_id VARCHAR(36) NOT NULL, user_id VARCHAR(36) NOT NULL, PRIMARY KEY(group_id, user_id));",
	"notes":         "CREATE TABLE notes (id VARCHAR(36) NOT NULL PRIMARY KEY, title VARCHAR(36), body VARCHAR(10000), status VARCHAR(10) DEFAULT 'Active', date_created datetime, date_edited datetime );",
	"note_users":    "CREATE TABLE note_users  (note_id VARCHAR(36) NOT NULL, user_id VARCHAR(36) NOT NULL, PRIMARY KEY(note_id, user_id));",
	"pdfs":          "CREATE TABLE pdfs (id VARCHAR(36) NOT NULL  PRIMARY KEY, title VARCHAR(36), pdf VARCHAR(36));",
	"pdf_users":     "CREATE TABLE pdf_users  (pdf_id VARCHAR(36) NOT NULL, user_id VARCHAR(36) NOT NULL, PRIMARY KEY(pdf_id, user_id));",
	"reviews":       "CREATE TABLE reviews (id VARCHAR(36) NOT NULL PRIMARY KEY, title VARCHAR(36), item_id VARCHAR(36), item_type VARCHAR(10), status VARCHAR(10) DEFAULT 'Active', date_created datetime, date_edited datetime );",
	"review_cues":   "CREATE TABLE review_cues (review_id VARCHAR(36) NOT NULL, cue_id VARCHAR(36) NOT NULL, PRIMARY KEY(review_id, cue_id));",
	"shared_group":  "CREATE TABLE shared_group (id VARCHAR(36) NOT NULL PRIMARY KEY, group_name VARCHAR(36), admin_id VARCHAR(36));",
	"review_users":  "CREATE TABLE review_users  (review_id VARCHAR(36), user_id VARCHAR(36), PRIMARY KEY(review_id, user_id));",
	"tags":          "CREATE TABLE tags (id VARCHAR(36) NOT NULL PRIMARY KEY, title VARCHAR(36), user_id VARCHAR(36), color VARCHAR(30), status VARCHAR(10) DEFAULT 'Active', date_created datetime, date_edited datetime);",
	"tag_items":     "CREATE TABLE tag_items (tag_id VARCHAR(36) NOT NULL, item_id VARCHAR(36), item_type VARCHAR(36));",
	"users":         "CREATE TABLE users (id VARCHAR(36) NOT NULL PRIMARY KEY, user_name VARCHAR(36),email VARCHAR(50),  password VARCHAR(10), status VARCHAR(10) DEFAULT 'Active', date_created datetime, date_edited datetime) ;"}
