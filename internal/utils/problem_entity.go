package utils

import (
	"fmt"
)

type extendedStringError struct {
	s string
}

func (e *extendedStringError) Error() string {
	return e.s
}

func (e *extendedStringError) Is(target error) bool {
	return e.Error() == target.Error()
}

func NewExtendedError(message string) error{
	return &extendedStringError{
		s: message,
	}
}

type Problem struct {
	Id uint64
	UserId uint64
	Text string
}

func (p *Problem) String() string  {
	return fmt.Sprintf("Id: %v, UserId: %v, Text: %v", p.Id, p.UserId, p.Text)
}

func (p *Problem) Bytes() []byte  {
	return []byte(p.String())
}

func (p *Problem) Clone() *Problem  {
	return NewProblem(p.Id, p.UserId, p.Text)
}

type ProblemCollection []Problem

func (pc ProblemCollection) SplitToBulks(butchSize uint) []ProblemCollection {
	problemBatchList := SplitToBulks(pc, butchSize)
	result := make([]ProblemCollection, 0, len(problemBatchList))
	for _, problemList := range problemBatchList {
		result = append(result, problemList)
	}

	return result
}

func (pc ProblemCollection) GetMapWithIdKey() (map[uint64]Problem, error) {
	return GetMapWithIdKey(pc)
}

func NewProblem(id uint64, userId uint64, text string) *Problem {
	return &Problem{
		Id: id,
		UserId: userId,
		Text: text,
	}
}

func SplitToBulks(problems []Problem, butchSize uint) [][]Problem  {
	if problems == nil {
		return nil
	}

	inListSize := len(problems)
	if int(butchSize) >= inListSize {
		return append(
			make([][]Problem, 0, 1),
			append(make([]Problem, 0, inListSize), problems...),
			)
	}

	outListSize := (inListSize/int(butchSize)) + 1
	outList := make([][]Problem, 0, outListSize)
	for startIndex, endIndex := 0, 0; startIndex < inListSize; startIndex+=int(butchSize) {
		endIndex = startIndex + int(butchSize)
		if endIndex > inListSize {
			endIndex = inListSize
		}

		outList = append(outList, append(make([]Problem, 0, int(butchSize)), problems[startIndex:endIndex]...))
	}

	return outList
}

func GetMapWithIdKey(problems []Problem) (map[uint64]Problem, error) {
	if problems == nil {
		return nil, NewExtendedError("problem list is not init")
	}

	if len(problems) == 0 {
		return nil, NewExtendedError("problem list is empty")
	}

	outList := make(map[uint64]Problem, len(problems))
	for _, currentProblem := range problems {
		outList[currentProblem.Id] = currentProblem
	}

	return outList, nil
}
