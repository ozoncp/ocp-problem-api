package ocp_problem_api

import (
	"context"
	desc "github.com/ozoncp/ocp-problem-api/pkg/ocp-problem-api"
)

func (pa *OcpProblemAPI) ListProblemsV1(_ context.Context, listQuery *desc.ProblemListQueryV1) (*desc.ProblemListV1, error) {
	result := &desc.ProblemListV1{
		List: []*desc.ProblemV1{},
	}

	list, err := pa.repo.ListEntities(0, 0)
	if err != nil {
		pa.logError("ListProblemsV1", listQuery, err)
		return nil, err
	}

	filter := listQuery.Filter
	hasIdsFilter := filter !=nil && filter.Ids != nil && len(filter.Ids) > 0
	hasUserIdsFilter := filter !=nil && filter.UserIds != nil && len(filter.UserIds) > 0

	filterIsSuccess := true
	for _, problem := range list {
		if hasIdsFilter {
			for _, id := range filter.Ids {
				if id == problem.Id {
					filterIsSuccess = true
					break
				}
				filterIsSuccess = false
			}

			if !filterIsSuccess {
				continue
			}
		}

		if hasUserIdsFilter {
			for _, userId := range filter.UserIds {
				if userId == problem.UserId {
					filterIsSuccess = true
					break
				}
				filterIsSuccess = false
			}
		}

		if filterIsSuccess {
			result.List = append(result.List, &desc.ProblemV1{
				Id: problem.Id,
				UserId: problem.UserId,
				Text: problem.Text,
			})
		}
	}

	pa.logResult("ListProblemsV1", listQuery, result)
	//
	//if listQuery.Limit > 0 {
	//	result.List = result.List[listQuery.Offset:(listQuery.Offset+listQuery.Limit)]
	//}

	return result, nil
}
