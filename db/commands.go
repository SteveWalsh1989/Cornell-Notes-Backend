package db

//CreateTableCommands ... stores all create table commands needed for application
var CreateTableCommands = map[string]string{
	"users":         "CREATE TABLE users (id VARCHAR(36) NOT NULL, first_name VARCHAR(36), last_name VARCHAR(36), password VARCHAR(10), status VARCHAR(10) DEFAULT 'Active', date_created datetime, date_edited datetime, date_deleted datetime  PRIMARY KEY (id)) ;",
	"folders":       "CREATE TABLE folders (id VARCHAR(36) NOT NULL, name VARCHAR(36), status VARCHAR(10) DEFAULT 'Active', date_created datetime, date_edited datetime, date_deleted datetime  PRIMARY KEY (ID));",
	"folders_users": "CREATE TABLE folder_users (folder_id VARCHAR(36) NOT NULL, user_id VARCHAR(36) NOT NULL);",
	"groups":        "CREATE TABLE groups (id VARCHAR(36) NOT NULL, name VARCHAR(36), item_id VARCHAR(36), item_type VARCHAR(10), status VARCHAR(10) DEFAULT 'Active', date_created datetime, date_edited datetime, date_deleted datetime  PRIMARY KEY (id));",
	"groups_users":  "CREATE TABLE group_users (group_id VARCHAR(36) NOT NULL, user_id VARCHAR(36) NOT NULL);",
	"cornell_notes": "CREATE TABLE cornell_notes (id VARCHAR(36) NOT NULL, name VARCHAR(36), item_id VARCHAR(36), item_type VARCHAR(10), status VARCHAR(10) DEFAULT 'Active', date_created datetime, date_edited datetime, date_deleted datetime  PRIMARY KEY (id));",
	"cornell_cues":  "CREATE TABLE cornell_cues (id VARCHAR(36) NOT NULL, name VARCHAR(36), item_id VARCHAR(36), item_type VARCHAR(10), status VARCHAR(10) DEFAULT 'Active', date_created datetime, date_edited datetime, date_deleted datetime  PRIMARY KEY (id));",
	"cornell_users": "CREATE TABLE cornell_users  (cornell_note_id VARCHAR(36) NOT NULL, user_id VARCHAR(36) NOT NULL );",
	"notes":         "CREATE TABLE notes (id VARCHAR(36) NOT NULL, name VARCHAR(36), item_id VARCHAR(36), item_type VARCHAR(10), status VARCHAR(10) DEFAULT 'Active', date_created datetime, date_edited datetime, date_deleted datetime  PRIMARY KEY (id));",
	"notes_users":   "CREATE TABLE note_users  (note_id VARCHAR(36) NOT NULL, user_id VARCHAR(36) NOT NULL);",
	"reviews":       "CREATE TABLE reviews (id VARCHAR(36) NOT NULL, name VARCHAR(36), item_id VARCHAR(36), item_type VARCHAR(10), status VARCHAR(10) DEFAULT 'Active', date_created datetime, date_edited datetime, date_deleted datetime  PRIMARY KEY (id));",
	"review_users":  "CREATE TABLE review_users  (review_id VARCHAR(36), user_id VARCHAR(36));",
	"tags":          "CREATE TABLE tags (id VARCHAR(36) NOT NULL, name VARCHAR(36), user_id VARCHAR(36), status VARCHAR(10) DEFAULT 'Active', date_created datetime, date_edited datetime, date_deleted datetime);",
	"tags_items":    "CREATE TABLE tag_items (id VARCHAR(36) NOT NULL, name VARCHAR(36), item_id VARCHAR(36), item_type VARCHAR(10), color VARCHAR(10), status VARCHAR(10) DEFAULT 'Active', date_created datetime, date_edited datetime, date_deleted datetime  PRIMARY KEY (id));",
	"pdf":           "CREATE TABLE pdfs (id VARCHAR(36) NOT NULL, name VARCHAR(36), item_id VARCHAR(36), item_type VARCHAR(10), status VARCHAR(10) DEFAULT 'Active', date_created datetime, date_edited datetime, date_deleted datetime  PRIMARY KEY (id));",
	"pdf_users":     "CREATE TABLE pdf_users  (pdf VARCHAR(36) NOT NULL, user_id VARCHAR(36) NOT NULL);",
	"badges":        "CREATE TABLE badges (id VARCHAR(36) NOT NULL, name VARCHAR(36), item_id VARCHAR(36), item_type VARCHAR(10), status VARCHAR(10), date_created datetime, date_edited datetime, date_deleted datetime  PRIMARY KEY (id));",
	"badges_users":  "CREATE TABLE badge_users (badge_id VARCHAR(36) NOT NULL, user_id VARCHAR(36) NOT NULL);",
}

//InsertSampleDataCommands ... inserts sample data into DB
var InsertSampleDataCommands = map[string]string{
	"User 1": "INSERT INTO users (id, first_name, last_name, password, status) VALUES('123332', 'Steve', 'Walsh', 'password', 'Active');",
	"User 2": "INSERT INTO users (id, first_name, last_name, password, status) VALUES('123343', 'Emma', 'Johnson', 'password', 'Active');",
	"User 3": "INSERT INTO users (id, first_name, last_name, password, status) VALUES('123343', 'Ted', 'Bundy', 'password', 'Deleted');",
}
