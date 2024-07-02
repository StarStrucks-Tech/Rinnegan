package postgres

import "database/sql"

type PostgresTransactionManager struct {
	db *sql.DB
}

func NewPostgresTransactionManager(db *sql.DB) *PostgresTransactionManager {
	return &PostgresTransactionManager{db: db}
}

func (ptm *PostgresTransactionManager) BeginTx() (*sql.Tx, error) {
	return ptm.db.Begin()
}

func (ptm *PostgresTransactionManager) CommitTx(tx *sql.Tx) error {
	return tx.Commit()
}

func (ptm *PostgresTransactionManager) Rollback(tx *sql.Tx) error {
	return tx.Rollback()
}
