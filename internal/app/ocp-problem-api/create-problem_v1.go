package ocp_problem_api

import (
	"context"
	"github.com/ozoncp/ocp-problem-api/internal/utils"
	desc "github.com/ozoncp/ocp-problem-api/pkg/ocp-problem-api"
)

func (pa *OcpProblemAPI) CreateProblemV1(ctx context.Context, problem *desc.ProblemV1) (*desc.ResultSaveV1, error) {
	err := pa.repo.AddEntities(ctx, []utils.Problem{
		{
			Id:     problem.Id,
			UserId: problem.UserId,
			Text:   problem.Text,
		},
	})

	if err != nil {
		pa.logError("CreateProblemV1", problem, err)
		return nil, err
	}

	result := &desc.ResultSaveV1{Id: problem.Id}
	pa.logResult("CreateProblemV1", problem, result)

	return result, nil
}
