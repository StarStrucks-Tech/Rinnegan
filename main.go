package main

import (
	"database/sql"
	"fmt"
	"time"

	"Rinnegan/common/database/postgres"

	_ "github.com/lib/pq" // Register the postgres driver
)

func main() {
	config := postgres.Config{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "root",
		Database: "Kubair",
	}

	// Create a new PostgresConnector instance
	connector := postgres.NewPostgresConnector(config)

	// Connect to the database
	db, err := connector.Connect(config)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close() // Close the connection when the function finishes

	fmt.Println("Successfully connected to database!")

	// Read all records from users table
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}
	defer rows.Close()

	// Define variables to store user data
	var id int
	var username string
	var email string
	var createdAt time.Time
	var updatedAt time.Time

	// Iterate through rows and scan data
	for rows.Next() {
		err := rows.Scan(&id, &username, &email, &createdAt, &updatedAt) // Modify based on your table columns
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		fmt.Printf("User ID: %d, Username: %s, Email: %s, Created at: %s, Modified at: %s\n", id, username, email, createdAt.Format("2006-01-02 15:04:05"), updatedAt.Format("2006-01-02 15:04:05"))
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return
	}

	fmt.Println("Successfully retrieved users!")

	// Example update query (assuming you have an ID to update)
	userID := 1 // Replace with the actual user ID you want to update
	newEmail := "updated_newmail@example.com"

	err = updateEmailWithTransaction(db, userID, newEmail)
	if err != nil {
		fmt.Println("Error updating user email:", err)
		return
	}

	fmt.Println("Successfully updated user email!")

	// Read all records from users table
	rows, err = db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &username, &email, &createdAt, &updatedAt) // Modify based on your table columns
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		fmt.Printf("User ID: %d, Username: %s, Email: %s, Created at: %s, Modified at: %s\n", id, username, email, createdAt.Format("2006-01-02 15:04:05"), updatedAt.Format("2006-01-02 15:04:05"))
	}
}

func updateEmailWithTransaction(db *sql.DB, userID int, newEmail string) error {
	// Create a PostgresTransactionManager instance
	transactionManager := postgres.NewPostgresTransactionManager(db)

	// Begin a transaction
	tx, err := transactionManager.BeginTx()
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback() // Rollback if any error occurs
			fmt.Println("Error updating user email: rolling back transaction")
		}
	}()

	// Prepare and execute update statement within the transaction
	stmt, err := tx.Prepare("UPDATE users SET email = $1 WHERE id = $2")
	if err != nil {
		return fmt.Errorf("error preparing update statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(newEmail, userID)
	if err != nil {
		return fmt.Errorf("error updating user email: %w", err)
	}

	// Commit the transaction if successful
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}
