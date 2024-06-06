package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	user     = "postgres"
	host     = "localhost"
	database = "KuberData"
	password = "Mani@2204073"
	port     = 5432
)

type Record struct {
	Username    string
	PhoneNumber string
	City        string
	Address     string
}

type DBHelper struct {
	db *sql.DB
}

func NewDBHelper() (*DBHelper, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}
	return &DBHelper{db: db}, nil
}

func (h *DBHelper) Close() {
	h.db.Close()
}

func (h *DBHelper) getAllRecords(tableName string) error {
	rows, err := h.db.Query("SELECT * FROM " + tableName)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var username string
		var phonenumber string
		var city string
		var address string
		err := rows.Scan(&username, &phonenumber, &city, &address)
		if err != nil {
			return err
		}
		fmt.Printf("Username: %s, Phone: %s, City: %s, Address: %s\n", username, phonenumber, city, address)
	}

	return nil
}

func (h *DBHelper) insertRecord(record Record, tableName string) error {
	insertDymStmt := fmt.Sprintf("INSERT into %s (username, phonenumber, city, address) VALUES($1, $2, $3, $4)", tableName)
	_, err := h.db.Exec(insertDymStmt, record.Username, record.PhoneNumber, record.City, record.Address)
	return err
}

func (h *DBHelper) updateRecord(record Record, tableName string) error {
	updateStmt := fmt.Sprintf("UPDATE %s SET city = $1 WHERE username = $2", tableName)
	_, err := h.db.Exec(updateStmt, record.City, record.Username)
	return err
}

func (h *DBHelper) deleteRecord(tableName string, primaryKeyValue string) error {
	deleteStmt := fmt.Sprintf("DELETE FROM %s WHERE username = $1", tableName)
	_, err := h.db.Exec(deleteStmt, primaryKeyValue)
	return err
}

func main() {
	helper, err := NewDBHelper()
	if err != nil {
		panic(err)
	}
	defer helper.Close()

	// CRUD operations demonstration
	record := Record{Username: "Shashi", PhoneNumber: "96610335556", City: "Munich", Address: "Germany"}

	record.City = "Delhi"
	err = helper.updateRecord(record, "customers")
	if err != nil {
		panic(err)
	}

	err = helper.deleteRecord("customers", "Mani")
	if err != nil {
		panic(err)
	}

	err = helper.insertRecord(record, "customers")
	if err != nil {
		panic(err)
	}

	err = helper.getAllRecords("customers")
	if err != nil {
		panic(err)
	}

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
