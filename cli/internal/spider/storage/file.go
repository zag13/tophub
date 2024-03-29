package storage

import (
	"encoding/csv"
	"github.com/zag13/tophub/cli/bootstrap"
	"github.com/zag13/tophub/cli/internal/spider/site"
	"os"
)

func File(tops []site.Top) error {
	file, err := os.OpenFile(bootstrap.C.FilePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)

	for _, top := range tops {
		if err := w.Write(top.StringSlice()); err != nil {
			return err
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		return err
	}

	return nil
}
