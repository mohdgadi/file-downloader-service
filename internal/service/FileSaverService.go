package service

import "io"

// FileSaverService is used to save and file to local disk
type FileSaverService struct {
}

// Save reads from input stream and writes to file and resturns that filePath
func (service FileSaverService) Save(stream io.Reader, err error) (filePath string) {
	// From stream write the stream to path and return file path
	return
}
