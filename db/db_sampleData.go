package db

//InsertSampleData ... stores sample data for DB
var InsertSampleData = map[string]string{
	"User 1":        "INSERT INTO users (id, name,email, password) VALUES('324ddsf3', 'Steve', 'S@dfs.com' ,'testing');",
	"User 2":        "INSERT INTO users (id, name,email, password) VALUES('sfs34344', 'Emma Watson', 'ewat@ymail.co.uk' ,'password');",
	"User 3":        "INSERT INTO users (id, name,email, password) VALUES('sfsdf443', 'Ted Bundy', 'todiefor@mail.com', 'password');",
	"Folder 1":      "INSERT INTO folders (id, name, date_created, date_edited) VALUES('werwr433', 'Machine Learning', '2020-01-03', '2020-01-03');",
	"Folder 2":      "INSERT INTO folders (id, name, date_created, date_edited) VALUES('434rsfg4', 'Cryptography', '2020-01-03', '2020-01-03');",
	"Folder User 1": "INSERT INTO folder_users (folder_id, user_id) VALUES('werwr433', '324ddsf3' );",
	"Folder User 2": "INSERT INTO folder_users (folder_id, user_id) VALUES('434rsfg4', '324ddsf3' );"}
