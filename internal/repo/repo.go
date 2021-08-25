package repo

import (
	"context"
	"github.com/ozoncp/ocp-problem-api/internal/utils"
)

type Repo interface {
	AddEntities(ctx context.Context, problems []utils.Problem) error
	UpdateEntity(ctx context.Context, problems utils.Problem) error
	ListEntities(ctx context.Context, limit, offset uint64) ([]utils.Problem, error)
	DescribeEntity(ctx context.Context, entityId uint64) (*utils.Problem, error)
}

type RepoRemover interface {
	Repo
	RemoveEntity(ctx context.Context, entityId uint64) error
}
