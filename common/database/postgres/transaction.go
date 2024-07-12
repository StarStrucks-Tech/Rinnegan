package postgres

import "database/sql"

// PostgresTransactionManager is a struct that holds a reference to a sql.DB and provides methods to manage database transactions.
type PostgresTransactionManager struct {
	db *sql.DB
}

/*
NewPostgresTransactionManager initializes a new PostgresTransactionManager with the given database connection.
Returns a pointer to PostgresTransactionManager.
*/
func NewPostgresTransactionManager(db *sql.DB) *PostgresTransactionManager {
	return &PostgresTransactionManager{db: db}
}

/*
BeginTx starts a new transaction.
Returns the transaction object or an error if starting the transaction fails.
*/
func (ptm *PostgresTransactionManager) BeginTx() (*sql.Tx, error) {
	return ptm.db.Begin()
}

/*
CommitTx commits the given transaction.
Accepts a transaction object as a parameter.
Returns an error if committing the transaction fails.
*/
func (ptm *PostgresTransactionManager) CommitTx(tx *sql.Tx) error {
	return tx.Commit()
}

/*
RollbackTx rolls back the given transaction.
Takes a transaction object as a parameter.
Returns an error if rolling back the transaction fails.
*/
func (ptm *PostgresTransactionManager) Rollback(tx *sql.Tx) error {
	return tx.Rollback()
}
