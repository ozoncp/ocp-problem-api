package ocp_problem_api

import (
	"context"
	desc "github.com/ozoncp/ocp-problem-api/pkg/ocp-problem-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (pa *OcpProblemAPI) ListProblemsV1(ctx context.Context, listQuery *desc.ProblemListQueryV1) (*desc.ProblemListV1, error) {
	result := &desc.ProblemListV1{
		List: []*desc.ProblemV1{},
	}

	list, err := pa.repo.ListEntities(ctx, listQuery.Limit, listQuery.Offset)
	if err != nil {
		pa.logError("ListProblemsV1", listQuery, err)
		return nil, status.Error(codes.Unknown, err.Error())
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

	return result, nil
}
