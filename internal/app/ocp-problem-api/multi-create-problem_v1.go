package ocp_problem_api

import (
	"context"
	"github.com/ozoncp/ocp-problem-api/internal/utils"
	desc "github.com/ozoncp/ocp-problem-api/pkg/ocp-problem-api"
)

func (pa *OcpProblemAPI) MultiCreateProblemV1(
	ctx context.Context,
	problemList *desc.ProblemListV1,
	) (*desc.ListResultSaveV1, error) {

	problemSize := len(problemList.List)
	result := &desc.ListResultSaveV1{List: make([]*desc.ResultSaveV1, 0, problemSize)}
	problems := make([]utils.Problem, 0, problemSize)
	for _, problemV1 := range problemList.List {
		result.List = append(result.List, &desc.ResultSaveV1{Id: problemV1.Id})
		problems = append(problems, utils.Problem{
			Id: problemV1.Id,
			UserId: problemV1.UserId,
			Text: problemV1.Text,
		})
	}

	err := pa.repo.AddEntities(ctx, problems)
	if err != nil {
		pa.logError("CreateProblemV1", problemList, err)
		return nil, err
	}

	pa.logResult("CreateProblemV1", problemList, result)

	return result, nil
}
