package config

var C Config

type Config struct {
	FilePath string `mapstructure:"CLI_FILE_PATH"`
}
