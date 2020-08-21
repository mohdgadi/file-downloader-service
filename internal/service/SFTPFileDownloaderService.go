package service

import (
	"file-downloader/model"
	"fmt"
)

// SFTPFileDownloaderService is the FTP implementation of Downloader Service.
type SFTPFileDownloaderService struct {
	// Injecting db as a dependecy to service
	db               model.IFileDownloaderDBConfig
	fileSaverService FileSaverService
}

// Download takes a config ID and downloads the file into local disk
func (service SFTPFileDownloaderService) Download(configID string) (filePath string, err error) {
	// Fetch config from db
	conf, err := service.db.Find(configID)
	fmt.Println(conf, err)

	// using config download file
	// pass ioStream to filesaverService and return filepath returned by FileSavserService.Save
	return
}
