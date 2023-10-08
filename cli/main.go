/*
Copyright © 2023 zag13
*/
package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/zag13/tophub/cli/bootstrap"
	"github.com/zag13/tophub/cli/internal/spider"
)

var rootCmd = &cobra.Command{
	Use:   "tophub-cli",
	Short: "tophub-cli is a cli tool for tophub.",
	Long:  `tophub-cli is a cli tool for tophub.`,
}

func main() {
	cobra.OnInitialize(bootstrap.InitConfig)

	rootCmd.PersistentFlags().StringVarP(&bootstrap.C.EnvPath, "env", "e", "../.env", "the env path")

	rootCmd.AddCommand(spider.CmdSpider)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
