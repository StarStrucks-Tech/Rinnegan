package postgres

type PostgresManager struct {
	Connector          *PostgresConnector
	Querier            *PostgresQuerier
	TransactionManager *PostgresTransactionManager
}

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
