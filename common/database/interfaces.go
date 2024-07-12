/*
Package db provides interfaces for database operations, including connection management,query execution, and transaction handling.
*/
package db

import "database/sql"

// Connector defines methods for managing database connections.
type Connector interface {
	/*
		Connect establishes a new database connection.
		It returns a pointer to the sql.DB object and an error if the connection fails.
	*/
	Connect() (*sql.DB, error)

	/*
		Close terminates the database connection.
		It returns an error if the operation fails.
	*/
	Close() error
}

// Querier defines methods for executing database queries and statements.
type Querier interface {

	/*
		ExecuteQuery executes a given SQL query with the specified arguments.
		It returns a pointer to sql.Rows and an error if the query fails.
	*/
	ExecuteQuery(query string, args ...interface{}) (*sql.Rows, error)

	/*
		ExecuteStatement executes a given SQL statement with the specified arguments.
		It returns an sql.Result object and an error if the statement execution fails.
	*/
	ExecuteStatement(query string, args ...interface{}) (sql.Result, error)
}

// TransactionManager defines methods for handling database transactions.
type TransactionManager interface {

	/*
		BeginTx starts a new database transaction.
		It returns a pointer to sql.Tx and an error if the transaction fails to start.
	*/
	BeginTx() (*sql.Tx, error)

	/*
		CommitTx commits the specified transaction.
		It returns an error if the commit operation fails.
	*/
	CommitTx(tx *sql.Tx) error

	/*
		RollbackTx rolls back the specified transaction.
		It returns an error if the rollback operation fails.
	*/
	RollbackTx(tx *sql.Tx) error
}

// DBManager combines Connector, Querier, and TransactionManager interfaces, providing a comprehensive interface for managing database operations.
type DBManager interface {
	Connector
	Querier
	TransactionManager
}
