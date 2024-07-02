package postgres

import "database/sql"

type PostgresQuerier struct {
	db *sql.DB
}

func NewPostgresQuerier(db *sql.DB) *PostgresQuerier {
	return &PostgresQuerier{db: db}
}

func (pq *PostgresQuerier) ExecuteQuery(query string, args ...interface{}) (*sql.Rows, error) {
	return pq.db.Query(query, args...)
}

func (pq *PostgresQuerier) ExecuteStatement(query string, args ...interface{}) (sql.Result, error) {
	return pq.db.Exec(query, args...)
}
