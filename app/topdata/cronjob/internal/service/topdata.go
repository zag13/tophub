package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
	v1 "tophub/api/topdata/cronjob/v1"
)

type TopdataService struct {
	v1.UnimplementedTopdataServer

	cron *cron.Cron
	log  *log.Helper
}

func NewTopDataService(cron *cron.Cron, logger log.Logger) *TopdataService {
	return &TopdataService{
		cron: cron,
		log:  log.NewHelper(log.With(logger, "module", "topdata-cronjob/service")),
	}
}

func NewCron(db *gorm.DB) *cron.Cron {
	c := cron.New(cron.WithSeconds())
	c.Start()

	if _, err := c.AddJob("* */10 * * * *", &Spider{DB: db}); err != nil {
		log.Errorf("failed opening connection to sqlite: %v", err)
		panic("failed to add Spider Job")
	}

	return c
}
