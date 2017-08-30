package server

import (
	"fmt"

	"github.com/spf13/viper"
)

// AppID is our infrastructure app identifier.
const AppID = "api"

// Config ...
type Config struct {
	AppID string
	Debug bool
	Port  int
}

// NewConfig creates new config.
func NewConfig(
	debug bool,
	port int,
) Config {
	return Config{
		AppID: AppID,
		Debug: debug,
		Port:  port,
	}
}

// LoadConfig loads config from file.
func LoadConfig(debug bool, environment string, configDir string) Config {
	v := viper.GetViper()

	v.AddConfigPath(configDir)
	v.SetConfigName(fmt.Sprintf("config.%s", environment))

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Unable to load config from %s", configDir))
	}

	port := v.GetInt("api.server.port")
	return NewConfig(debug, port)
}
