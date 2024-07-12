package postgres

type Config struct {
	Host     string // Database server host.
	Port     int    // Port on which the database server is listening.
	User     string // Username used to authenticate with the database
	Password string // Password used to authenticate with the database.
	Database string // Name of the database to connect to.
}
