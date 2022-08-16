package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Job is a Job model.
type Job struct {
	Id      int64
	JobName string
}

// JobRepo is a Job repo.
type JobRepo interface {
	CreateJob(context.Context, *Job) (*Job, error)
	UpdateJob(context.Context, *Job) (*Job, error)
	DeleteJob(context.Context, int64) (*Job, error)
	GetJob(context.Context, string) ([]*Job, error)
	ListJob(context.Context) ([]*Job, error)
}

// JobUseCase is a Job use case.
type JobUseCase struct {
	repo JobRepo
	log  *log.Helper
}

// NewJobUseCase new a Job use case.
func NewJobUseCase(repo JobRepo, logger log.Logger) *JobUseCase {
	return &JobUseCase{repo: repo, log: log.NewHelper(logger)}
}
