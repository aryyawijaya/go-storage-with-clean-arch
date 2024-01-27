package utilconfig

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBSourceLocal     string `mapstructure:"DB_SOURCE_LOCAL"`
	DBSourceContainer string `mapstructure:"DB_SOURCE_CONTAINER"`
	ServerAddress     string `mapstructure:"SERVER_ADDRESS"`
	APIKey            string `mapstructure:"API_KEY"`
	APIKeyValue       string `mapstructure:"API_KEY_VALUE"`
}

func LoadConfig(path string) (config *Config, err error) {
	configFile := fmt.Sprintf("%s/.env", path)
	viper.SetConfigFile(configFile)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
