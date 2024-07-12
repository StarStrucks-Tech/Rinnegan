package postgres

import "database/sql"

// PostgresQuerier is a struct that holds a reference to a sql.DB and provides methods to execute queries and statements.
type PostgresQuerier struct {
	db *sql.DB
}

/*
NewPostgresQuerier initializes a new PostgresQuerier with the given database connection.
Returns a pointer to PostgresQuerier.
*/
func NewPostgresQuerier(db *sql.DB) *PostgresQuerier {
	return &PostgresQuerier{db: db}
}

/*
ExecuteQuery executes a query that returns rows, typically a SELECT.
Takes a query string and optional arguments.
Returns the resulting rows or an error if the execution fails.
*/
func (pq *PostgresQuerier) ExecuteQuery(query string, args ...interface{}) (*sql.Rows, error) {
	return pq.db.Query(query, args...)
}

/*
ExecuteStatement executes a query without returning any rows, typically an INSERT, UPDATE, or DELETE.
Takes a query string and optional arguments.
Returns the result of the execution or an error if the execution fails.
*/
func (pq *PostgresQuerier) ExecuteStatement(query string, args ...interface{}) (sql.Result, error) {
	return pq.db.Exec(query, args...)
}
