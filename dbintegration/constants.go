package main

const (
	// SQL Queries
	SelectAllRecordsQuery = "SELECT * FROM %s"
	InsertRecordQuery     = "INSERT INTO %s (username, phonenumber, city, address) VALUES($1, $2, $3, $4)"
	UpdateRecordQuery     = "UPDATE %s SET city = $1 WHERE username = $2"
	DeleteRecordQuery     = "DELETE FROM %s WHERE username = $1"
)
