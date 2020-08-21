package model

import (
	"database/sql"
	"file-downloader/entities"
)

// SQLFILEDownloaderConfig is SQL implementation of IFileDownloaderDBConfig
type SQLFILEDownloaderConfig struct {
	conn *sql.Conn
}

// NewSQLFILEDownloaderConfig is a constructor
func NewSQLFILEDownloaderConfig(dbUsername string,
	dbPassword string) (sqlDB SQLFILEDownloaderConfig, err error) {
	// Connect to sql an initialize database object and return
	return
}

//  Implement interface methods

// Find from SQL db
func (db SQLFILEDownloaderConfig) Find(configID string) (config entities.Config, err error) {
	// Find config in table by id and return
	return
}

// Insert into SQL db
func (db SQLFILEDownloaderConfig) Insert(entities.Config) (id int, err error) {
	// Insert config into table
	return
}

// Delete from SQL db
func Delete(configID string) (err error) {
	// delete config from db
	return
}

// Update into SQL db
func Update(entities.Config) (err error) {
	// Update config into db
	return
}
