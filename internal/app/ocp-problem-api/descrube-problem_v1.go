package ocp_problem_api

import (
	"context"
	desc "github.com/ozoncp/ocp-problem-api/pkg/ocp-problem-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (pa *OcpProblemAPI) DescribeProblemV1(ctx context.Context, problemQuery *desc.ProblemQueryV1) (*desc.ProblemV1, error) {
	problem, err := pa.repo.DescribeEntity(ctx, problemQuery.Id)
	if err != nil {
		pa.logError("DescribeProblemV1", problemQuery, err)
		return nil, status.Error(codes.NotFound, err.Error())
	}

	result := &desc.ProblemV1{
		Id: problem.Id,
		UserId: problem.UserId,
		Text: problem.Text,
	}

	pa.logResult("DescribeProblemV1", problemQuery, result)

	return result, nil
}
