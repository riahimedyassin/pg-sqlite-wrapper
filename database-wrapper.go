package pglitewrapper

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type DatabaseWrapper struct {
	dbConfig *DatabaseConfig
	rcConfig *ReconnectConfig
	DB       *gorm.DB
}

func Connect(dbConfig *DatabaseConfig, rcConfig *ReconnectConfig) (*DatabaseWrapper, error) {
	var db *gorm.DB
	var err error
	switch dbConfig.Driver {
	case "postgres":
		db, err = gorm.Open(postgres.Open(dbConfig.DSN), &gorm.Config{})
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(dbConfig.DSN), &gorm.Config{})
	default:
		return nil, errors.New("Unsupported database driver: " + dbConfig.Driver)
	}
	if err != nil {
		return nil, err
	}
	return &DatabaseWrapper{dbConfig: dbConfig, rcConfig: rcConfig, DB: db}, nil
}

func NewDatabaseWrapper(dbConfig *DatabaseConfig, rcConfig *ReconnectConfig) (*DatabaseWrapper, error) {
	dbWrapper, err := Connect(dbConfig, rcConfig)
	if !rcConfig.AutoReconnect {
		return dbWrapper, nil
	}
	for rcConfig.Attempts > 0 {
		dbWrapper, err = Connect(dbConfig, rcConfig)
		if err == nil {
			return dbWrapper, nil
		}
		fmt.Println("Cannot connect to database retrying")
		time.Sleep(time.Second * 2)
		rcConfig.Attempts--
	}
	return nil, errors.New("Cannot connect to database ")

}
