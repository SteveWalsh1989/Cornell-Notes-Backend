package db

//InsertSampleData ... stores sample data for DB
var InsertSampleData = map[string]string{
	"User 1":        "INSERT INTO users (id, user_name, email, password) VALUES('324ddsf3', 'Steve', 'S@dfs.com' ,'testing');",
	"User 2":        "INSERT INTO users (id, user_name, email, password) VALUES('sfs34344', 'Emma Watson', 'ewat@ymail.co.uk' ,'password');",
	"User 3":        "INSERT INTO users (id, user_name, email, password) VALUES('sfsdf443', 'Ted Bundy', 'todiefor@mail.com', 'password');",
	"Folder 1":      "INSERT INTO folders (id, title, date_created, date_edited) VALUES('werwr433', 'Machine Learning', '2020-01-03', '2020-01-03');",
	"Folder 2":      "INSERT INTO folders (id, title, date_created, date_edited) VALUES('434rsfg4', 'Cryptography', '2020-01-03', '2020-01-03');",
	"Folder User 1": "INSERT INTO folder_users (folder_id, user_id) VALUES('werwr433', '324ddsf3' );",
	"Folder User 2": "INSERT INTO folder_users (folder_id, user_id) VALUES('434rsfg4', '324ddsf3' );",
	"Note 1":        "INSERT INTO notes (id, title) VALUES('12313c1ccc3223cd', 'Spyder' );",
	"Note 2":        "INSERT INTO notes (id, title) VALUES('14234254533334cd', 'Ciphers' );",
	"CornellNote 1": "INSERT INTO cornell_notes (id, title) VALUES('ewrwerewercccvv333', 'Veniger Ciphers' );",
	"CornellNote 2": "INSERT INTO cornell_notes (id, title) VALUES('23rresfdf434r34r34', 'Casear Ciphers' );",
	"cn User 1":     "INSERT INTO cornell_users (cornell_note_id, user_id ) VALUES('ewrwerewercccvv333', '324ddsf3' );",
	"cn User 2":     "INSERT INTO cornell_users (cornell_note_id, user_id ) VALUES('23rresfdf434r34r34', '324ddsf3' );",
	"folder_item1":  "INSERT INTO folder_items (folder_id, item_id, item_type) VALUES('werwr433', '12313c1ccc3223cd', 'Note' );",
	"folder_item2":  "INSERT INTO folder_items (folder_id, item_id, item_type) VALUES('434rsfg4', '14234254533334cd',  'Note');",
	"folder_item3":  "INSERT INTO folder_items (folder_id, item_id, item_type) VALUES('434rsfg4', 'ewrwerewercccvv333',  'CornellNote');",
	"folder_item4":  "INSERT INTO folder_items (folder_id, item_id, item_type) VALUES('434rsfg4', '23rresfdf434r34r34',  'CornellNote');",
}
