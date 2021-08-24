package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/ozoncp/ocp-problem-api/internal/utils"
	"sync"
)

type fakeRepoProblem struct {
	*sync.Mutex
	store []utils.Problem
}

func (frp *fakeRepoProblem) AddEntities(_ context.Context, problems []utils.Problem) error {
	var err error
	frp.Lock()
	defer frp.Unlock()

	for _, problem := range problems {
		if p, _ := frp.describeEntity(problem.Id); p == nil {
			frp.store = append(frp.store, problem)
		} else {
			err = NewRepoError(fmt.Sprintf("duplicate problem #%d", problem.Id), &problem, err)
		}
	}

	return err
}

func (frp *fakeRepoProblem) ListEntities(_ context.Context, limit, offset uint64) ([]utils.Problem, error) {
	frp.Lock()
	defer frp.Unlock()

	if limit == 0 && offset == 0 {
		return frp.store, nil
	}

	storeSize := uint64(len(frp.store))
	if offset >= storeSize {
		return nil, errors.New("invalid offset")
	}

	diffSize := storeSize - offset
	if diffSize < limit || limit == 0 {
		return append(make([]utils.Problem, 0, diffSize), frp.store[offset:]...), nil
	}

	return append(make([]utils.Problem, 0, limit), frp.store[offset:(offset+limit)]...), nil
}

func (frp *fakeRepoProblem) describeEntity(entityId uint64) (*utils.Problem, error) {
	for _, problem := range frp.store {
		if problem.Id == entityId {
			return &problem, nil
		}
	}

	return nil, errors.New("problem not found")
}

func (frp *fakeRepoProblem) DescribeEntity(_ context.Context, entityId uint64) (*utils.Problem, error) {
	frp.Lock()
	defer frp.Unlock()

	return frp.describeEntity(entityId)
}

func (frp *fakeRepoProblem) RemoveEntity(_ context.Context, entityId uint64) error {
	frp.Lock()
	frp.Unlock()
	for i, problem := range frp.store {
		if problem.Id == entityId {
			frp.store = append(append(
				make([]utils.Problem, 0, len(frp.store)-1),
				frp.store[:i]...
			), append(
				make([]utils.Problem, 0, len(frp.store[i+1:])),
				frp.store[i+1:]...
			)...)

			return nil
		}
	}

	return errors.New("problem not found")
}

func NewFakeRepo() RepoRemover {
	return &fakeRepoProblem{Mutex: &sync.Mutex{}, store: []utils.Problem{}}
}
