package bootstrap

import (
	"flag"
	"log"

	"github.com/spf13/viper"
)

var envFile = flag.String("e", "../.env", "the env file")

type Config struct {
	AppEnv   string `mapstructure:"SERVER_APP_ENV"`
	Port     string `mapstructure:"SERVER_PORT"`
	MySQLDSN string `mapstructure:"SERVER_MYSQL_DSN"`
}

func NewConfig() *Config {
	c := Config{}

	viper.SetConfigFile(*envFile)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the env file: ", err)
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if c.AppEnv == "dev" {
		log.Println("The App is running in development env")
	}

	return &c
}
