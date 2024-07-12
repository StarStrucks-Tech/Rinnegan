package postgres

import (
	"database/sql"
	"strconv"

	_ "github.com/lib/pq"
)

// PostgresConnector handles the connection to a PostgreSQL database.
type PostgresConnector struct {
	config Config
	db     *sql.DB
}

// NewPostgresConnector initializes a new PostgresConnector with the given configuration.
func NewPostgresConnector(config Config) *PostgresConnector {
	return &PostgresConnector{config: config}
}

// buildConnectionString constructs the PostgreSQL connection string from the given configuration.
func buildConnectionString(config Config) string {
	return "postgres://" + config.User + ":" + config.Password + "@" +
		config.Host + ":" + strconv.Itoa(config.Port) + "/" + config.Database + "?sslmode=disable"
}

/*
Connect establishes a connection to the PostgreSQL database using the provided configuration.
Returns a sql.DB instance representing the connection or an error if the connection fails.
*/

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

// Close closes the database connection.
func (pc *PostgresConnector) Close() error {
	return pc.db.Close()
}
