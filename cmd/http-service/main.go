package http-service
import "file-downloader/internal/model"

func main()  {
	// initialize SQLFileDownloaderConfigDb
	// initialize HTTP Service and inject db as a dependency
	// Startserver to recieve download request. Pass service as a parameter
}

func startServer(service HTTPFileDownloaderService) {
	// recieve request to download file
}