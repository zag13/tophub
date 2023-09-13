/*
Copyright © 2023 zag13
*/
package main

import (
	"flag"
	"log"

	"github.com/spf13/cobra"
	"github.com/zag13/tophub/cli/bootstrap"
	"github.com/zag13/tophub/cli/internal/spider"
)

var rootCmd = &cobra.Command{
	Use:   "topcli",
	Short: "topcli is a cli tool for tophub.",
	Long:  `topcli is a cli tool for tophub.`,
}

func main() {
	flag.Parse()

	cobra.OnInitialize(bootstrap.InitConfig)

	rootCmd.AddCommand(spider.CmdSpider)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
