package ocp_problem_api

import (
	"context"
	desc "github.com/ozoncp/ocp-problem-api/pkg/ocp-problem-api"
)

func (pa *OcpProblemAPI) RemoveProblemV1(ctx context.Context, problemQuery *desc.ProblemQueryV1) (*desc.ProblemResultV1, error) {
	if err := pa.repo.RemoveEntity(ctx, problemQuery.Id); err != nil {
		pa.logError("RemoveProblemV1", problemQuery, err)
		return nil, err
	}

	result := &desc.ProblemResultV1{
		Id: problemQuery.Id,
		IsSuccess: true,
		Text: "element was removed",
	}

	pa.logResult("RemoveProblemV1", problemQuery, result)

	return result, nil
}
