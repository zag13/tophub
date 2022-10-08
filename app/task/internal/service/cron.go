package service

import "github.com/robfig/cron/v3"

func InitCron() error {
	c := cron.New(cron.WithSeconds())
	c.Start()

	if _, err := c.AddJob("*/5 * * * * *", &Spider{}); err != nil {
		return err
	}

	return nil
}
