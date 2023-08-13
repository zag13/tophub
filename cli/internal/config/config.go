package config

var C Config

type Config struct {
	FilePath string `mapstructure:"FILE_PATH"`
}
