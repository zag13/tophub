package bootstrap

import "github.com/spf13/viper"

var Env env

type env struct {
	IsShowErrors bool `mapstructure:"SERVER_IS_SHOW_ERRORS"`
}

func initEnv() {
	Env.IsShowErrors = viper.GetBool("SERVER_IS_SHOW_ERRORS")
}
