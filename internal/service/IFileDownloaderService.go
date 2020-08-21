package service

// IFileDownloaderService is an interface which abstracts out the functionality to download a file
// over different networking protocols.
type IFileDownloaderService interface {
	Download(configID string) (string, error)
}
