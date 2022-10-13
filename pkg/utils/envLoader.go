package utils

import "github.com/spf13/viper"

type EnvConfig struct {
	Db_connection string `mapstructure:"DB_CONNECTION"`
	Jwt_secret    string `mapstructure:"JWT_SECRET"`
}

func LoadConfig(envName string) (config EnvConfig, err error) {
	viper.AddConfigPath("./environments")
	viper.SetConfigName(envName)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
