package config

var C Config

type Config struct {
	Port string `mapstructure:"PORT"`
}
