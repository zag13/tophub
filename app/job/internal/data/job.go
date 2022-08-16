package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"tophub/app/job/internal/biz"
)

var _ biz.JobRepo = (*jobRepo)(nil)

type jobRepo struct {
	data *Data
	log  *log.Helper
}

func NewJobRepo(data *Data, logger log.Logger) biz.JobRepo {
	return &jobRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "data/job")),
	}
}

func (j *jobRepo) CreateJob(ctx context.Context, job *biz.Job) (*biz.Job, error) {
	// TODO implement me
	panic("implement me")
}

func (j *jobRepo) UpdateJob(ctx context.Context, job *biz.Job) (*biz.Job, error) {
	// TODO implement me
	panic("implement me")
}

func (j *jobRepo) DeleteJob(ctx context.Context, i int64) (*biz.Job, error) {
	// TODO implement me
	panic("implement me")
}

func (j *jobRepo) GetJob(ctx context.Context, s string) ([]*biz.Job, error) {
	// TODO implement me
	panic("implement me")
}

func (j *jobRepo) ListJob(ctx context.Context) ([]*biz.Job, error) {
	// TODO implement me
	panic("implement me")
}
