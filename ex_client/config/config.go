package config

import "github.com/spf13/viper"

const (
	ServerAddress = "server_address"
	Port          = "port"

	ClientType = "client_type"
)

var (
	config *viper.Viper
)

func Init() {
	config = viper.New()
	config.AutomaticEnv()
}

func GetServerAddress() string {
	return config.GetString(ServerAddress)
}

func GetPort() string {
	return config.GetString(Port)
}

func GetClientType() string {
	return config.GetString(ClientType)
}
