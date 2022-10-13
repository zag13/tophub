package service

import (
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

func InitCron(db *gorm.DB) error {
	c := cron.New(cron.WithSeconds())
	c.Start()

	if _, err := c.AddJob("*/10 * * * * *", &Spider{DB: db}); err != nil {
		return err
	}

	return nil
}
