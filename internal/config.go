package internal

// Config /
type Config struct {
	Postgres PostgresConfig
}

// PostgresConfig /
type PostgresConfig struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     string
}

var prodConfig = Config{
	// In real production code, these values would be read from environment variables / secrets manager
	Postgres: PostgresConfig{
		Host:     "localhost",
		User:     "fizzbuzz",
		Password: "fizzbuzz_password",
		DbName:   "fizzbuzz_db",
		Port:     "5432",
	},
}
