package repo

import (
	"errors"
	"github.com/ozoncp/ocp-problem-api/internal/utils"
)

type repoError struct {
	problem *utils.Problem
	err error
}

func (re *repoError) Error() string {
	return re.err.Error()
}

func (re *repoError) Unwrap() error {
	return errors.Unwrap(re.err)
}

func NewRepoError(message string, problem *utils.Problem, err error) error {
	return &repoError{
		problem: problem,
		err: utils.NewWrappedError(message, err),
	}
}
