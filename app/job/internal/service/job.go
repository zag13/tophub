package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"tophub/app/job/internal/biz"

	pb "tophub/api/job/v1"
)

type JobService struct {
	pb.UnimplementedJobServer

	oc  *biz.JobUseCase
	log *log.Helper
}

func NewJobService(oc *biz.JobUseCase, logger log.Logger) *JobService {
	return &JobService{
		oc:  oc,
		log: log.NewHelper(log.With(logger, "module", "service/job")),
	}
}

func (s *JobService) CreateJob(ctx context.Context, req *pb.CreateJobRequest) (*pb.CreateJobReply, error) {
	return &pb.CreateJobReply{}, nil
}
func (s *JobService) UpdateJob(ctx context.Context, req *pb.UpdateJobRequest) (*pb.UpdateJobReply, error) {
	return &pb.UpdateJobReply{}, nil
}
func (s *JobService) DeleteJob(ctx context.Context, req *pb.DeleteJobRequest) (*pb.DeleteJobReply, error) {
	return &pb.DeleteJobReply{}, nil
}
func (s *JobService) GetJob(ctx context.Context, req *pb.GetJobRequest) (*pb.GetJobReply, error) {
	return &pb.GetJobReply{}, nil
}
func (s *JobService) ListJob(ctx context.Context, req *pb.ListJobRequest) (*pb.ListJobReply, error) {
	return &pb.ListJobReply{}, nil
}
