package repo

import "github.com/ozoncp/ocp-problem-api/internal/utils"

type repoError struct {
	problem *utils.Problem
	message string
	err error
}

func (re *repoError) Error() string {
	return re.message
}

func (re *repoError) Unwrap() error {
	return re.err
}

func NewRepoError(message string, problem *utils.Problem, err error) error {
	return &repoError{
		message: message,
		problem: problem,
		err: err,
	}
}
