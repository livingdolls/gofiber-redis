package config

import "github.com/spf13/viper"

type Config struct {
	DBHost string `mapstructure:"SQL_HOST"`
	DBUsername string `mapstructure:"SQL_USER"`
	DBPassword string `mapstructure:"SQL_PASSWORD"`
	DBName string `mapstructure:"SQL_DB"`
	DBPort string `mapstructure:"SQL_PORT"`

	RedisUrl string `mapstructure:"REDIS_URL"`
}

func LoadConfig(path string) (config Config, err error){
	viper.AddConfigPath(path)
	viper.SetConfigFile("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	// Handle Null

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return

}