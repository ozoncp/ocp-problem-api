package flusher

import (
	"context"
	"github.com/ozoncp/ocp-problem-api/internal/repo"
	"github.com/ozoncp/ocp-problem-api/internal/utils"
)

type Flusher interface {
	Flush(ctx context.Context, problems []utils.Problem) []utils.Problem
}

type flusher struct {
	chunkSize int
	problemRepo repo.Repo
}

func (f flusher) Flush(ctx context.Context, problems []utils.Problem) []utils.Problem {
	if problems == nil {
		return nil
	}

	problemsSize := len(problems)
	if problemsSize == 0 {
		return nil
	}

	if problemsSize < f.chunkSize {
		if err := f.problemRepo.AddEntities(ctx, problems); err != nil {
			return problems
		}

		return nil
	}

	returnList := make([]utils.Problem, 0, problemsSize)
	for startIndex, endIndex := 0, 0; startIndex < problemsSize; startIndex += f.chunkSize {
		endIndex = startIndex + f.chunkSize
		if endIndex > problemsSize {
			endIndex = problemsSize
		}

		if err := f.problemRepo.AddEntities(ctx, problems[startIndex:endIndex]); err != nil {
			returnList = append(returnList, problems[startIndex:endIndex]...)
		}
	}

	if len(returnList) == 0 {
		return nil
	}

	return returnList
}

func NewFlusher(chunkSize int, problemRepo repo.Repo) Flusher {
	return &flusher{
		chunkSize: chunkSize,
		problemRepo: problemRepo,
	}
}