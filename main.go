package main

import (
	"Rinnegan/common/database/postgres"
	"fmt"
	"time"
)

func main() {
	// Configuration
	config := postgres.Config{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "root", // Replace with actual password (consider environment variables)
		Database: "Kubair",
	}

	// Database Manager
	manager, err := postgres.NewPostgresManager(config)
	if err != nil {
		fmt.Println("Error creating database manager:", err)
		return
	}
	defer manager.Connector.Close() // Close connection when main exits

	// Read users
	users, err := getAllUsers(*manager.Querier)
	if err != nil {
		fmt.Println("Error retrieving users:", err)
		return
	}
	fmt.Println("Successfully retrieved users!")
	for _, user := range users {
		fmt.Printf("User ID: %d, Username: %s, Email: %s, Created at: %s, Modified at: %s\n",
			user.ID, user.Username, user.Email, user.CreatedAt.Format("2006-01-02 15:04:05"), user.UpdatedAt.Format("2006-01-02 15:04:05"))
	}

	// Update user email (example)
	userID := 1 // Replace with actual ID
	newEmail := "updated_email@example.com"
	err = updateEmail(*manager.TransactionManager, userID, newEmail)
	if err != nil {
		fmt.Println("Error updating user email:", err)
		return
	}
	fmt.Println("Successfully updated user email!")
}

// Define a struct to represent a user
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetAllUsers retrieves all users from the database
func getAllUsers(querier postgres.PostgresQuerier) ([]User, error) {
	rows, err := querier.ExecuteQuery("SELECT * FROM users")
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return users, nil
}

// UpdateEmail updates the email address of a user within a transaction
func updateEmail(txManager postgres.PostgresTransactionManager, userID int, newEmail string) error {
	tx, err := txManager.BeginTx()
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback() // Rollback if any error occurs
			fmt.Println("Error updating user email: rolling back transaction")
		}
	}()

	stmt, err := tx.Prepare("UPDATE users SET email = $1 WHERE id = $2")
	if err != nil {
		return fmt.Errorf("error preparing update statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(newEmail, userID)
	if err != nil {
		return fmt.Errorf("error updating user email: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}
