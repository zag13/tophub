/*
Copyright © 2023 zag13
*/
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zag13/tophub/cli/internal/config"
	"github.com/zag13/tophub/cli/internal/spider"
)

var (
	envFile = flag.String("e", "../.env", "the env file")
	cfgFile = flag.String("c", "./etc/tophub-cli.yaml", "the config file")

	rootCmd = &cobra.Command{
		Use:   "topcli",
		Short: "topcli is a cli tool for tophub.",
		Long:  `topcli is a cli tool for tophub.`,
	}
)

func main() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(spider.CmdSpider)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func initConfig() {
	flag.Parse()

	if err := godotenv.Load(*envFile); err != nil {
		fmt.Println(fmt.Sprintf("Loading env file error: %v", err))
		fmt.Println("Using default config file: ./etc/tophub-cli.yaml")
	}

	viper.SetConfigFile(*cfgFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("CLI")

	if err := viper.Unmarshal(&config.C); err != nil {
		log.Fatalf("Unable to unmarshal config: %v", err)
	}
}
