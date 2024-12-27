package interfaces

import (
	"astor-bd-wrapper/cmd/config"
	"database/sql"
)

type CommonDatabaseManager[database sql.DB] interface {
	Connect(driver string, dbConfig config.DatabaseConfig) (*database, error)
	// Reconnect is used to reconnect to database when the connection fails.
	Reconnect(reconnectConfig config.ReconnectConfig) error
}
