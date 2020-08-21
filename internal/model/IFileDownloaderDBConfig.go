package model

import "../entities"

// IFileDownloaderDBConfig is an interface abstracting IfilerDownloaderConfigDB
type IFileDownloaderDBConfig interface {
	Find(configID string) (entities.Config, error)
	Insert(entities.Config) (int, error)
	Delete(configID string) error
	Update(entities.Config) error
}
