package postgres

// PostgresManager manages the PostgreSQL database operations by combining a connector, querier, and transaction manager.
type PostgresManager struct {
	Connector          *PostgresConnector
	Querier            *PostgresQuerier
	TransactionManager *PostgresTransactionManager
}

/*
NewPostgresManager initializes a new PostgresManager with the given configuration.
Connects to the PostgreSQL database and sets up the connector, querier, and transaction manager.
Returns a pointer to PostgresManager or an error if the connection fails.
*/
func NewPostgresManager(config Config) (*PostgresManager, error) {
	connector := NewPostgresConnector(config)
	db, err := connector.Connect(config)
	if err != nil {
		return nil, err
	}
	return &PostgresManager{
		Connector:          connector,
		Querier:            NewPostgresQuerier(db),
		TransactionManager: NewPostgresTransactionManager(db),
	}, nil
}
