package config

import "github.com/spf13/viper"

const (
	//ServerType hi or caculate
	ServerType = "server_type"
	Port       = "port"
)

var (
	config *viper.Viper
)

func Init() {
	config = viper.New()
	config.AutomaticEnv()
}

func GetServerType() string {
	return config.GetString(ServerType)
}

func GetPort() string {
	return config.GetString(Port)
}
