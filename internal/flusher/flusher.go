package flusher

import (
	"context"
	"errors"
	"github.com/opentracing/opentracing-go"
	"github.com/ozoncp/ocp-problem-api/internal/repo"
	"github.com/ozoncp/ocp-problem-api/internal/utils"
	"strconv"
)

type Flusher interface {
	Flush(ctx context.Context, problems []utils.Problem) []utils.Problem
	FlushWithError(ctx context.Context, problems []utils.Problem) ([]utils.Problem, error)
}

type flusher struct {
	chunkSize int
	problemRepo repo.Repo
}

func (f *flusher) Flush(ctx context.Context, problems []utils.Problem) []utils.Problem {
	result, _ := f.FlushWithError(ctx, problems)

	return result
}

func (f *flusher) FlushWithError(ctx context.Context, problems []utils.Problem) ([]utils.Problem, error) {
	var methodErr error

	tracer := opentracing.GlobalTracer()
	parentSpan := opentracing.SpanFromContext(ctx)

	if problems == nil {
		methodErr = errors.New("problem list is empty")
		return nil, methodErr
	}

	problemsSize := len(problems)
	if problemsSize == 0 {
		span := tracer.StartSpan("FlushWithError", opentracing.ChildOf(parentSpan.Context()))
		defer span.Finish()

		span.SetBaggageItem("save-count", strconv.Itoa(problemsSize))
		span.SetBaggageItem("result", "error")

		methodErr = errors.New("problem list is empty")
		return nil, methodErr
	}

	if problemsSize < f.chunkSize {
		span := tracer.StartSpan("FlushWithError", opentracing.ChildOf(parentSpan.Context()))
		defer span.Finish()

		span.SetBaggageItem("save-count", strconv.Itoa(problemsSize))
		if err := f.problemRepo.AddEntities(ctx, problems); err != nil {
			span.SetBaggageItem("result", "error")
			return problems, err
		}

		span.SetBaggageItem("result", "success")

		return nil, methodErr
	}

	returnList := make([]utils.Problem, 0, problemsSize)
	for startIndex, endIndex := 0, 0; startIndex < problemsSize; startIndex += f.chunkSize {
		span := tracer.StartSpan("FlushWithError", opentracing.ChildOf(parentSpan.Context()))

		endIndex = startIndex + f.chunkSize
		if endIndex > problemsSize {
			endIndex = problemsSize
		}

		problemsToSave := problems[startIndex:endIndex]
		span.SetBaggageItem("save-count", strconv.Itoa(problemsSize))
		if err := f.problemRepo.AddEntities(ctx, problemsToSave); err != nil {
			span.SetBaggageItem("result", "error")
			methodErr = utils.NewWrappedError(err.Error(), methodErr)
			returnList = append(returnList, problems[startIndex:endIndex]...)
		} else {
			span.SetBaggageItem("result", "success")
		}

		span.Finish()
	}

	if len(returnList) == 0 {
		return nil, methodErr
	}

	return returnList, methodErr
}

func NewFlusher(chunkSize int, problemRepo repo.Repo) Flusher {
	return &flusher{
		chunkSize: chunkSize,
		problemRepo: problemRepo,
	}
}