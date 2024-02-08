package rest

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/litleleprikon/codegen_go_talk/model"
	v1 "github.com/litleleprikon/codegen_go_talk/pkg/api/rest/coverage/v1"
	"github.com/litleleprikon/codegen_go_talk/repository"
)

var _ v1.ServerInterface = &CoverageServer{}

type CoverageServer struct {
	repo repository.Coverage
}

func New(repo repository.Coverage) *CoverageServer {
	return &CoverageServer{
		repo: repo,
	}
}

// AddCommitFilesCoverage implements v1.ServerInterface.
func (s *CoverageServer) AddCommitFilesCoverage(ctx echo.Context, commit string) error {
	reqData := v1.AddCommitFilesCoverageJSONRequestBody{}
	ctx.Bind(&reqData)
	coverage := map[model.FileName][]model.LineCoverage{}
	for file, fileCov := range reqData.Coverage {
		for _, cov := range fileCov {
			coverage[model.FileName(file)] = append(coverage[model.FileName(file)], model.LineCoverage{
				Line:     model.LineNumber(cov.Line),
				Coverage: model.NewLineCoverageType(string(cov.Coverage)),
			})
		}
	}
	s.repo.AddCommitInfo(model.CommitHash(commit), coverage)
	ctx.JSON(http.StatusOK, map[string]any{"ok": true})
	return nil
}

// ListCommitFilesCoverage implements v1.ServerInterface.
func (s *CoverageServer) ListCommitFilesCoverage(ctx echo.Context, commit string) error {
	coverage := s.repo.GetCommitInfo(model.CommitHash(commit))
	response := v1.CommitCoverage{}
	for file, fileCov := range coverage {
		for _, cov := range fileCov {
			response.Coverage[string(file)] = append(response.Coverage[string(file)], v1.LineCoverage{
				Line:     int(cov.Line),
				Coverage: v1.LineCoverageCoverage(cov.Coverage.String()),
			})
		}
	}
	ctx.JSON(http.StatusOK, coverage)
	return nil
}
