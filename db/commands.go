package db

/*
 Commands:
 --------------------------------------------
 file contains Maps containing SQL commands to create
 tables and insert sample data into the application
*/

//CreateTableCommands ... stores all create table commands needed for application - 20 tables
var CreateTableCommands = map[string]string{
	"badges":        "CREATE TABLE badges (id VARCHAR(36) NOT NULL  PRIMARY KEY, name VARCHAR(36), item_id VARCHAR(36), item_type VARCHAR(10), status VARCHAR(10), date_created datetime, date_edited datetime);",
	"badge_users":   "CREATE TABLE badge_users (badge_id VARCHAR(36) NOT NULL, user_id VARCHAR(36) NOT NULL, PRIMARY KEY(badge_id, user_id));",
	"badge_scores":  "CREATE TABLE badge_scores (user_id VARCHAR(36) NOT NULL PRIMARY KEY, notes_created integer, cornell_notes_created integer,reviews_created integer,notes_shared integer, cornell_notes_shared integer,reviews_shared integer, reviews_completed integer);",
	"cornell_notes": "CREATE TABLE cornell_notes (id VARCHAR(36) NOT NULL PRIMARY KEY, name VARCHAR(36), status VARCHAR(10) DEFAULT 'Active', date_created datetime, date_edited datetime);",
	"cornell_cues":  "CREATE TABLE cornell_cues (id VARCHAR(36) NOT NULL PRIMARY KEY, cornell_note_id VARCHAR(36), cue VARCHAR(100), answer VARCHAR(100));",
	"cornell_users": "CREATE TABLE cornell_users  (cornell_note_id VARCHAR(36) NOT NULL, user_id VARCHAR(36) NOT NULL, PRIMARY KEY(cornell_note_id, user_id));",
	"folders":       "CREATE TABLE folders (id VARCHAR(36) NOT NULL PRIMARY KEY, name VARCHAR(36), status VARCHAR(10) DEFAULT 'Active', date_created datetime, date_edited datetime);",
	"folder_users":  "CREATE TABLE folder_users (folder_id VARCHAR(36) NOT NULL, user_id VARCHAR(36) NOT NULL, PRIMARY KEY(folder_id, user_id));",
	"group_users":   "CREATE TABLE group_users (group_id VARCHAR(36) NOT NULL, user_id VARCHAR(36) NOT NULL, PRIMARY KEY(group_id, user_id));",
	"notes":         "CREATE TABLE notes (id VARCHAR(36) NOT NULL PRIMARY KEY, name VARCHAR(36), note VARCHAR(100), status VARCHAR(10) DEFAULT 'Active', date_created datetime, date_edited datetime );",
	"note_users":    "CREATE TABLE note_users  (note_id VARCHAR(36) NOT NULL, user_id VARCHAR(36) NOT NULL, PRIMARY KEY(note_id, user_id));",
	"pdfs":          "CREATE TABLE pdfs (id VARCHAR(36) NOT NULL  PRIMARY KEY, name VARCHAR(36), pdf VARCHAR(36));",
	"pdf_users":     "CREATE TABLE pdf_users  (pdf_id VARCHAR(36) NOT NULL, user_id VARCHAR(36) NOT NULL, PRIMARY KEY(pdf_id, user_id));",
	"reviews":       "CREATE TABLE reviews (id VARCHAR(36) NOT NULL PRIMARY KEY, name VARCHAR(36), item_id VARCHAR(36), item_type VARCHAR(10), status VARCHAR(10) DEFAULT 'Active', date_created datetime, date_edited datetime );",
	"review_cues":   "CREATE TABLE review_cues (review_id VARCHAR(36) NOT NULL, cue_id VARCHAR(36) NOT NULL, PRIMARY KEY(review_id, cue_id));",
	"shared_group":  "CREATE TABLE shared_group (id VARCHAR(36) NOT NULL PRIMARY KEY, name VARCHAR(36), admin_id VARCHAR(36));",
	"review_users":  "CREATE TABLE review_users  (review_id VARCHAR(36), user_id VARCHAR(36), PRIMARY KEY(review_id, user_id));",
	"tags":          "CREATE TABLE tags (id VARCHAR(36) NOT NULL PRIMARY KEY, name VARCHAR(36), user_id VARCHAR(36), status VARCHAR(10) DEFAULT 'Active', date_created datetime, date_edited datetime);",
	"tag_items":     "CREATE TABLE tag_items (tag_id VARCHAR(36) NOT NULL, item_id VARCHAR(36), item_type VARCHAR(36));",
	"users":         "CREATE TABLE users (id VARCHAR(36) NOT NULL PRIMARY KEY, first_name VARCHAR(36), last_name VARCHAR(36), password VARCHAR(10), status VARCHAR(10) DEFAULT 'Active', date_created datetime, date_edited datetime) ;"}

//InsertSampleDataCommands ... inserts sample data into DB
var InsertSampleDataCommands = map[string]string{
	"User 1":        "INSERT INTO users (id, first_name, last_name, password) VALUES('123332', 'Steve', 'Walsh', 'password');",
	"User 2":        "INSERT INTO users (id, first_name, last_name, password) VALUES('123343', 'Emma', 'Johnson', 'password');",
	"User 3":        "INSERT INTO users (id, first_name, last_name, password) VALUES('123449', 'Ted', 'Bundy', 'password');",
	"Folder 1":      "INSERT INTO folders (id, name, date_created, date_edited) VALUES('131311', 'Machine Learning', '2020-01-03', '2020-01-03');",
	"Folder 2":      "INSERT INTO folders (id, name, date_created, date_edited) VALUES('1234439', 'Cryptography', '2020-01-03', '2020-01-03');",
	"Folder User 1": "INSERT INTO folder_users (folder_id, user_id) VALUES('131311', '123332' );",
	"Folder User 2": "INSERT INTO folder_users (folder_id, user_id) VALUES('1234439', '123343' );"}
