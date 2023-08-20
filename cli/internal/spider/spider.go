package spider

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zag13/tophub/cli/internal/config"
	"github.com/zag13/tophub/cli/internal/spider/endpoint"
	"github.com/zag13/tophub/cli/internal/spider/storage"
)

var (
	CmdSpider = &cobra.Command{
		Use:   "spider",
		Short: "spider data from endpoint.",
		Long:  `spider data from endpoint.`,
		RunE:  run,
	}

	epUrl string

	epHandlers = map[string]func(...endpoint.Options) (tops []endpoint.Top, err error){
		"zhihu": endpoint.ZhiHu,
	}

	stgHandlers = map[string]func(tops []endpoint.Top) error{
		"file": storage.File,
	}
)

func init() {
	CmdSpider.Flags().StringVarP(&epUrl, "url", "u", "", "the url of endpoint")
}

func run(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("please enter the endpoint name")
	}

	fmt.Println(config.C)
	epHandler, ok := epHandlers[args[0]]
	if !ok {
		return errors.New("endpoint not found")
	}
	tops, err := epHandler(endpoint.WithUrl(epUrl))
	if err != nil {
		return err
	}

	stgHandler := storage.File
	if len(args) > 1 {
		if val, ok := stgHandlers[args[1]]; ok {
			stgHandler = val
		}
	}
	return stgHandler(tops)
}
