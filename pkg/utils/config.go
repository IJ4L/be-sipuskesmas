package utils

import "github.com/spf13/viper"

type Config struct {
	App           string `mapstructure:"app_port"`
	DBDriver      string `mapstructure:"db_driver"`
	ServerAddress string `mapstructure:"server_address"`

	DBHost     string `mapstructure:"db_host"`
	DBPort     string `mapstructure:"db_port"`
	DBUser     string `mapstructure:"db_user"`
	DBPassword string `mapstructure:"db_password"`
	DBName     string `mapstructure:"db_name"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(".env")
	viper.SetConfigName(".")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	return
}
