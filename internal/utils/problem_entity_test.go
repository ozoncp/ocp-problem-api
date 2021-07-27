package utils

import (
	"errors"
	"reflect"
	"testing"
)

func TestNewProblem(t *testing.T) {
	solution := &Problem{
		Id: 1,
		UserId: 2,
		Text: "some text",
	}

	result := NewProblem(1,2, "some text")
	if *solution != *result {
		t.Errorf("Invalid result: %#v, need: %#v", *result, *solution)
	}
}

func TestGetMapWithIdKey(t *testing.T) {
	solution := map[uint64]Problem{
		1: {1,2, "some text1"},
		2: {2,2, "some text2"},
		3: {3,2, "some text3"},
	}

	in := []Problem{
		{1,2, "some text1"},
		{2,2, "some text2"},
		{3,2, "some text3"},
	}

	result, _ := GetMapWithIdKey(in)

	if !reflect.DeepEqual(solution, result) {
		t.Errorf("Invalid result map: %v, need: %v", result, solution)
	}
}

func TestGetMapWithIdKeyNilValue(t *testing.T) {
	checkError := NewExtendedError("problem list is not init")
	_, err := GetMapWithIdKey(nil)
	if !errors.Is(checkError, err) {
		t.Error("Unknown error!")
	}
}

func TestGetMapWithIdKeyEmptyValue(t *testing.T) {
	checkError := NewExtendedError("problem list is empty")
	_, err := GetMapWithIdKey([]Problem{})
	if !errors.Is(checkError, err) {
		t.Error("Unknown error!")
	}
}

func TestSplitToBulks(t *testing.T) {
	solution := [][]Problem{
		{
			{1,2,"some text1"},
			{2,2,"some text2"},
		},
		{
			{3,2,"some text3"},
			{4,2,"some text4"},
		},
		{
			{5,2,"some text5"},
		},
	}

	in := []Problem{
		{1,2,"some text1"},
		{2,2,"some text2"},
		{3,2,"some text3"},
		{4,2,"some text4"},
		{5,2,"some text5"},
	}

	result := SplitToBulks(in, 2)
	if !reflect.DeepEqual(solution, result) {
		t.Errorf("Invalid result batch slice: %v, need: %v", result, solution)
	}
}

func TestProblem_String(t *testing.T) {
	solution := "Id: 1, UserId: 2, Text: some text"
	in := NewProblem(1,2,"some text")
	result := in.String()

	if solution != result {
		t.Errorf("Invalid result: %#v, need: %#v", result, solution)
	}
}

func TestProblem_Bytes(t *testing.T) {
	solution := []byte("Id: 1, UserId: 2, Text: some text")
	in := NewProblem(1,2,"some text")
	result := in.Bytes()

	if !reflect.DeepEqual(solution, result) {
		t.Errorf("Invalid result: %v, need: %v", result, solution)
	}
}

func TestProblem_Clone(t *testing.T) {
	in := NewProblem(1, 2, "some text")
	result := in.Clone()
	if *in != *result {
		t.Error("Values is not equal")
	}

	if in == result {
		t.Error("Invalid address new value")
	}

	result.Id = 123
	if in.Id == result.Id {
		t.Error("Invalid clone operation")
	}
}

func TestProblemCollection_GetMapWithIdKey(t *testing.T) {
	solution := map[uint64]Problem{
		5: {5, 2, "some text1"},
		6: {6, 2, "some text2"},
		7: {7, 2, "some text3"},
	}

	inCollection := ProblemCollection{
		{5, 2, "some text1"},
		{6, 2, "some text2"},
		{7, 2, "some text3"},
	}

	result, _ := inCollection.GetMapWithIdKey()
	if !reflect.DeepEqual(solution, result) {
		t.Errorf("Invalid result: %v, need: %v", result, solution)
	}
}

func TestProblemCollection_SplitToBulks(t *testing.T) {
	solution := []ProblemCollection{
		{
			{1,2,"some text1"},
			{2,2,"some text2"},
		},
		{
			{3,2,"some text3"},
			{4,2,"some text4"},
		},
		{
			{5,2,"some text5"},
		},
	}

	inCollection := ProblemCollection{
		{1,2,"some text1"},
		{2,2,"some text2"},
		{3,2,"some text3"},
		{4,2,"some text4"},
		{5,2,"some text5"},
	}

	result := inCollection.SplitToBulks(2)
	if !reflect.DeepEqual(solution, result) {
		t.Errorf("Invalid result batch slice: %v, need: %v", result, solution)
	}
}
