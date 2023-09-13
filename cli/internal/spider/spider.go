package spider

import (
	"context"
	"errors"

	"github.com/spf13/cobra"
	"github.com/zag13/tophub/cli/internal/spider/site"
	"github.com/zag13/tophub/cli/internal/spider/storage"
	"golang.org/x/sync/errgroup"
)

var (
	CmdSpider = &cobra.Command{
		Use:   "spider",
		Short: "spider data from endpoint.",
		Long:  `spider data from endpoint.`,
		RunE:  run,
	}

	sites []string

	siteHandlers = map[string]func(...site.Options) (tops []site.Top, err error){
		"zhihu": site.ZhiHu,
	}

	stgHandlers = map[string]func(tops []site.Top) error{
		"db":   storage.DB,
		"file": storage.File,
	}
)

func init() {
	CmdSpider.Flags().StringSliceVarP(&sites, "sites", "s", []string{}, "site url")
}

func run(cmd *cobra.Command, args []string) error {
	g, _ := errgroup.WithContext(context.Background())

	if len(sites) == 0 {
		for key := range siteHandlers {
			sites = append(sites, key)
		}
	}
	results := map[string][]site.Top{}
	errs := map[string]error{}

	for _, s := range sites {
		s := s
		g.Go(func() error {
			handler, ok := siteHandlers[s]
			if !ok {
				errs[s] = errors.New("site not found")
				return nil
			}

			tops, err := handler()
			if err != nil {
				errs[s] = err
			} else {
				results[s] = tops
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return err
	}

	stgHandler := storage.DB
	if len(args) > 1 {
		if val, ok := stgHandlers[args[1]]; ok {
			stgHandler = val
		}
	}

	tops := []site.Top{}
	for _, result := range results {
		for _, top := range result {
			tops = append(tops, top)
		}
	}

	return stgHandler(tops)
}
