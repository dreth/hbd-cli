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
	return port
}

func DefaultSSL() bool {
	// get ssl from env vars
	LoadEnvVars()
	ssl := viper.GetBool("HBD_SSL")
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
