package managers

import (
	"astor-bd-wrapper/cmd/config"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DatabaseWrapper struct {
	dbConfig *config.DatabaseConfig
	rcConfig *config.ReconnectConfig
	DB       *gorm.DB
}

func NewDatabaseWrapper(dbConfig *config.DatabaseConfig, rcConfig *config.ReconnectConfig) (*DatabaseWrapper, error) {
	var db *gorm.DB
	var err error
	switch dbConfig.Driver {
	case "postgres":
		db, err = gorm.Open(postgres.Open(dbConfig.DSN), &gorm.Config{})
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(dbConfig.DSN), &gorm.Config{})
	default:
		panic("Unsupported database driver: " + dbConfig.Driver)
	}
	if err != nil {
		return nil, err
	}
	return &DatabaseWrapper{dbConfig: dbConfig, rcConfig: rcConfig, DB: db}, nil
}

func (dbWrapper *DatabaseWrapper) Reconnect() error {
	return nil
}
