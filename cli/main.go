/*
Copyright © 2023 zag13
*/
package main

import (
	"flag"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zag13/tophub/cli/internal/config"
	"github.com/zag13/tophub/cli/internal/spider"
)

var (
	envFile = flag.String("e", "../.env", "the env file")

	rootCmd = &cobra.Command{
		Use:   "topcli",
		Short: "topcli is a cli tool for tophub.",
		Long:  `topcli is a cli tool for tophub.`,
	}
)

func main() {
	flag.Parse()

	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(spider.CmdSpider)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func initConfig() {
	viper.SetConfigFile(*envFile)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the env file: ", err)
	}

	err = viper.Unmarshal(&config.C)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}
}
