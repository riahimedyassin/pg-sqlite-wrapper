package managers

import (
	"astor-bd-wrapper/cmd/config"
	"database/sql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type CommonManager struct {
	dbConfig config.DatabaseConfig
	rcConfig config.ReconnectConfig
	db       *gorm.DB
}

func NewCommonManager(dbConfig config.DatabaseConfig, rcConfig config.ReconnectConfig) (*CommonManager, error) {
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
	return &CommonManager{dbConfig: dbConfig, rcConfig: rcConfig, db: db}, nil
}

func (cManager *CommonManager) Connect() (*sql.DB, error) {
	db, err := sql.Open(cManager.dbConfig.Driver, cManager.dbConfig.DSN)
	if err != nil {
		return nil, err
	}
	return db, nil
}
func (cManager *CommonManager) Reconnect() error {
	return nil
}
