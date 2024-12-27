package pglitewrapper

type DatabaseConfig struct {
	DSN    string
	Driver string // Supported drivers Postgres or SQLite
}

type ReconnectConfig struct {
	AutoReconnect bool
	Attempts      int
}

func NewDatabaseConfig(dsn, driver string) *DatabaseConfig {
	if driver == "" {
		panic("NewDatabaseConfig: driver is required")
	}
	if dsn == "" {
		panic("NewDatabaseConfig: DSN is required")
	}
	return &DatabaseConfig{DSN: dsn, Driver: driver}
}

func NewReconnectConfig(autoReconnect bool, attempts int) *ReconnectConfig {
	if attempts <= 0 {
		panic("NewReconnectConfig: attempts must be > 0")
	}
	return &ReconnectConfig{AutoReconnect: autoReconnect, Attempts: attempts}
}
