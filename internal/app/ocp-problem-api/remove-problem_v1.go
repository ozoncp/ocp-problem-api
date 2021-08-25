package ocp_problem_api

import (
	"context"
	desc "github.com/ozoncp/ocp-problem-api/pkg/ocp-problem-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (pa *OcpProblemAPI) RemoveProblemV1(ctx context.Context, problemQuery *desc.ProblemQueryV1) (*desc.ProblemResultV1, error) {
	if err := pa.repo.RemoveEntity(ctx, problemQuery.Id); err != nil {
		pa.logError("RemoveProblemV1", problemQuery, err)
		return nil, status.Error(codes.NotFound, err.Error())
	}

	result := &desc.ProblemResultV1{
		Id: problemQuery.Id,
		IsSuccess: true,
		Text: "element was removed",
	}

	pa.logResult("RemoveProblemV1", problemQuery, result)

	return result, nil
}
