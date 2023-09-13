package bootstrap

import (
	"flag"
	"log"
	"sync"

	"github.com/spf13/viper"
)

var (
	envFile = flag.String("e", "../.env", "the env file")

	C       Config
	onceCfg sync.Once
)

type Config struct {
	FilePath string `mapstructure:"CLI_FILE_PATH"`
	MySQLDSN string `mapstructure:"MYSQL_DSN"`
}

func InitConfig() {
	onceCfg.Do(func() {
		viper.SetConfigFile(*envFile)

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal("Can't find the env file: ", err)
		}

		err = viper.Unmarshal(&C)
		if err != nil {
			log.Fatal("Environment can't be loaded: ", err)
		}
	})

}
