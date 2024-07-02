package postgres

import (
	"database/sql"
	"strconv"

	_ "github.com/lib/pq"
)

type PostgresConnector struct {
	config Config
	db     *sql.DB
}

func NewPostgresConnector(config Config) *PostgresConnector {
	return &PostgresConnector{config: config}
}

func buildConnectionString(config Config) string {
	return "postgres://" + config.User + ":" + config.Password + "@" +
		config.Host + ":" + strconv.Itoa(config.Port) + "/" + config.Database + "?sslmode=disable"
}

func (pc *PostgresConnector) Connect(config Config) (*sql.DB, error) {
	var err error

	// Build the connection string using config values
	connectionString := buildConnectionString(config)

	pc.db, err = sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return pc.db, nil
}

func (pc *PostgresConnector) Close() error {
	return pc.db.Close()
}
