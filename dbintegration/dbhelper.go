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
	query := fmt.Sprintf(SelectAllRecordsQuery, tableName)
	rows, err := h.db.Query(query)
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
	query := fmt.Sprintf(InsertRecordQuery, tableName)
	_, err := h.db.Exec(query, record.Username, record.PhoneNumber, record.City, record.Address)
	return err
}

func (h *DBHelper) updateRecord(record Record, tableName string) error {
	query := fmt.Sprintf(UpdateRecordQuery, tableName)
	_, err := h.db.Exec(query, record.City, record.Username)
	return err
}

func (h *DBHelper) deleteRecord(primaryKeyValue string, tableName string) error {
	query := fmt.Sprintf(DeleteRecordQuery, tableName)
	_, err := h.db.Exec(query, primaryKeyValue)
	return err
}
