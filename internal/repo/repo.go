package repo

import (
	"github.com/ozoncp/ocp-problem-api/internal/utils"
)

type Repo interface {
	AddEntities(problems []utils.Problem) error
	ListEntities(limit, offset uint64) ([]utils.Problem, error)
	DescribeEntity(entityId uint64) (*utils.Problem, error)
}
