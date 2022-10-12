package service

import (
	"github.com/robfig/cron/v3"
)

func InitCron(serv *Service) error {
	c := cron.New(cron.WithSeconds())
	c.Start()

	if _, err := c.AddJob("*/10 * * * * *", &Spider{Service: serv}); err != nil {
		return err
	}

	return nil
}
