package helper

import (
	"github.com/spf13/viper"
)

func DefaultHost() string {
	// get host from env vars
	LoadEnvVars()
	host := viper.GetString("HBD_HOST")

	// if the host is empty, use 0.0.0.0
	if host == "" {
		host = "0.0.0.0"
	}

	return host
}

func DefaultPort() string {
	// get port from env vars
	LoadEnvVars()
	port := viper.GetString("HBD_PORT")
	host := viper.GetString("HBD_HOST")

	// If both the host and port are empty, default the port to 8417
	if port == "" && host == "" {
		port = "8417"
	}

	return port
}

func DefaultSSL() bool {
	// get ssl from env vars
	LoadEnvVars()
	sslStr := viper.GetString("HBD_SSL")
	ssl := viper.GetBool("HBD_HOST")

	// if the host is empty, use 0.0.0.0
	if sslStr == "" {
		ssl = true
	}

	return ssl
}

func GenUrl(host string, port string, ssl bool) string {
	protocol := "http"
	if ssl {
		protocol = "https"
	}

	if port != "" {
		return protocol + "://" + host + ":" + port
	}

	return protocol + "://" + host
}
