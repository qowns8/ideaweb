package models

var (
	getInfoWithEmail = "SELECT id, email, password, access_key, name, update_at FROM user WHERE email = ? AND password = ?"
	getInfoWithKey   = "SELECT id, email, password, access_key, name, update_at FROM user WHERE access_key = ?"
	getAccessKey     = "SELECT id, access_key FROM user WHERE email = ? AND password = ?"
	ok               = "SELECT id FROM user WHERE access_key = ?"
	insert           = "INSERT INTO user (email, password, access_key, name) VALUES (?, ?, ?, ?)"
	tokenUpdate      =  "UPDATE user SET access_key = ? WHERE id = ? "
)