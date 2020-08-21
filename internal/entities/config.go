package entities

import "net/url"

// Protocol denotes the type of communication user over network.
type Protocol string

const (
	// FTP denotes ftp network protocol
	FTP Protocol = "ftp"

	// HTTP denotes http network protocol
	HTTP = "http"

	// SFTP denotes sftp network protocol
	SFTP = "sftp"
)

// Config is the configuartion of newtork to download the file.
type Config struct {
	Protocol   Protocol
	RemoteHost url.URL
	Username   string
	Password   string
}
