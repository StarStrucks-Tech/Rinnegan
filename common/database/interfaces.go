package db

import "database/sql"

type Connector interface {
	Connect() (*sql.DB, error)
	Close() error
}

type Querier interface {
	ExecuteQuery(query string, args ...interface{}) (*sql.Rows, error)
	ExecuteStatement(query string, args ...interface{}) (sql.Result, error)
}

type TransactionManager interface {
	BeginTx() (*sql.Tx, error)
	CommitTx(tx *sql.Tx) error
	RollbackTx(tx *sql.Tx) error
}

type DBManager interface {
	Connector
	Querier
	TransactionManager
}
