package http-service
import "file-downloader/internal/model"

func main()  {
	// initialize SQLFileDownloaderConfigDb
	// initialize SFTP Service and inject db as a dependency
	// Startserver to recieve download request. Pass service as a parameter
}

func startServer(service SFTPFileDownloaderService) {
	// recieve request to download file
}