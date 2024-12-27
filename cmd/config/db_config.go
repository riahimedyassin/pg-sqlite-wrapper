package config

type DatabaseConfig struct {
	DSN    string
	Driver string // Supported drivers Postgres or SQLite
}

type ReconnectConfig struct {
	AutoReconnect bool
	Attempts      int
}

func NewDatabaseConfig(dsn string) *DatabaseConfig {
	return &DatabaseConfig{DSN: dsn}
}

func NewReconnectConfig(autoReconnect bool, attempts int) *ReconnectConfig {
	return &ReconnectConfig{AutoReconnect: autoReconnect, Attempts: attempts}
}
