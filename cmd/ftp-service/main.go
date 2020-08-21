package ftp-service
import "file-downloader/internal/model"

func main()  {
	// initialize SQLFileDownloaderConfigDb
	// initialize FTP Service and inject db as a dependency
	// Startserver to recieve download request. Pass service as a parameter
}

func startServer(service FTPFileDownloaderService) {
	// recieve request to download file
}