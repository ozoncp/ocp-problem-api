package repo

import (
	"context"
	"github.com/ozoncp/ocp-problem-api/internal/utils"
)

type repoChain struct {
	RepoRemover
	repoList []RepoRemover
}

func (rc *repoChain) AddEntities(ctx context.Context, problems []utils.Problem) error {
	var errMethod error
	for _, repo := range rc.repoList {
		if err := repo.AddEntities(ctx, problems); err != nil {
			errMethod = utils.NewWrappedError(err.Error(), errMethod)
		}
	}

	return errMethod
}
func (rc *repoChain) UpdateEntity(ctx context.Context, problem utils.Problem) error {
	var errMethod error
	for _, repo := range rc.repoList {
		if err := repo.UpdateEntity(ctx, problem); err != nil {
			errMethod = utils.NewWrappedError(err.Error(), errMethod)
		}
	}

	return errMethod
}

func (rc *repoChain) ListEntities(ctx context.Context, limit, offset uint64) ([]utils.Problem, error) {
	var errMethod error
	for _, repo := range rc.repoList {
		list, err := repo.ListEntities(ctx, limit, offset)
		if err != nil {
			errMethod = utils.NewWrappedError(err.Error(), errMethod)
		} else {
			return list, errMethod
		}
	}

	return nil, errMethod
}

func (rc *repoChain) DescribeEntity(ctx context.Context, entityId uint64) (*utils.Problem, error) {
	var errMethod error
	for _, repo := range rc.repoList {
		problem, err := repo.DescribeEntity(ctx, entityId)
		if err != nil {
			errMethod = utils.NewWrappedError(err.Error(), errMethod)
		} else {
			return problem, errMethod
		}
	}

	return nil, errMethod
}

func (rc *repoChain) RemoveEntity(ctx context.Context, entityId uint64) error {
	var errMethod error
	for _, repo := range rc.repoList {
		err := repo.RemoveEntity(ctx, entityId)
		if err != nil {
			errMethod = utils.NewWrappedError(err.Error(), errMethod)
		}
	}

	return errMethod
}

func NewRepoChain(repos ...RepoRemover) RepoRemover {
	return &repoChain{repoList: repos}
}