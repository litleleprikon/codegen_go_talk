package grpc

import (
	"context"

	"github.com/litleleprikon/codegen_go_talk/model"
	v1 "github.com/litleleprikon/codegen_go_talk/pkg/api/grpc/coverage/v1"
	"github.com/litleleprikon/codegen_go_talk/repository"
)

var _ v1.CoverageServiceServer = &CoverageGRPCServer{}

type CoverageGRPCServer struct {
	v1.UnimplementedCoverageServiceServer
	repo repository.Coverage
}

func New(repo repository.Coverage) *CoverageGRPCServer {
	return &CoverageGRPCServer{
		repo: repo,
	}
}

// AddCommitCoverage implements v1.CoverageServiceServer.
func (s *CoverageGRPCServer) AddCommitCoverage(_ context.Context, r *v1.AddCommitCoverageRequest) (*v1.AddCommitCoverageResponse, error) {
	coverage := map[model.FileName][]model.LineCoverage{}

	for _, c := range r.Coverage {
		filename := model.FileName(c.File)
		coverage[filename] = append(coverage[filename], model.LineCoverage{
			Line:     model.LineNumber(c.Line),
			Coverage: model.LineCoverageType(c.Type),
		})
	}
	s.repo.AddCommitInfo(model.CommitHash(r.Hash), coverage)
	return &v1.AddCommitCoverageResponse{}, nil
}

// GetCommitCoverage implements v1.CoverageServiceServer.
func (s *CoverageGRPCServer) GetCommitCoverage(_ context.Context, r *v1.GetCommitCoverageRequest) (*v1.GetCommitCoverageResponse, error) {
	commit := model.CommitHash(r.Hash)
	info := s.repo.GetCommitInfo(commit)
	response := &v1.GetCommitCoverageResponse{}
	for file, fileCov := range info {
		for _, cov := range fileCov {
			response.Coverage = append(response.Coverage, &v1.LineCoverage{
				File: string(file),
				Line: int64(cov.Line),
				Type: v1.CoverageType(cov.Coverage),
			})
		}
	}
	return response, nil
}
