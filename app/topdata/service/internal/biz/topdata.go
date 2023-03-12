package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

// Topdata 定义业务中的数据结构
type Topdata struct {
	ID         int64
	Tab        string
	Position   int64
	Title      string
	URL        string
	Extra      string
	SpiderTime time.Time
}

type TopdataRepo interface {
	List(ctx context.Context, pageNum, pageSize int64) ([]*Topdata, error)
}

type TopdataUseCase struct {
	repo TopdataRepo
	log  *log.Helper
}

func NewTopdataUseCase(repo TopdataRepo, logger log.Logger) *TopdataUseCase {
	return &TopdataUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "topdata-service/biz"))}
}

func (uc *TopdataUseCase) List(ctx context.Context, pageNum, pageSize int64) ([]*Topdata, error) {
	return uc.repo.List(ctx, pageNum, pageSize)
}
