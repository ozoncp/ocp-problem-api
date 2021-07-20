package utils

import (
	"reflect"
	"testing"
)

func TestSplitStringSlice(t *testing.T) {
	solution := [][]string{{"one","two","three"},{"four","five","six"},{"seven","eight"}}
	in := []string{"one","two","three","four","five","six","seven","eight"}
	out := SplitStringSlice(in, 3)
	if !reflect.DeepEqual(out, solution) {
		t.Errorf("Invalid split, current value: %v, need: %v", out, solution)
	}
}

func TestSplitIntegerSlice(t *testing.T) {
	solution := [][]int{{1,2,3},{4,5,6},{7,8}}
	in := []int{1,2,3,4,5,6,7,8}
	out := SplitIntegerSlice(in, 3)
	if !reflect.DeepEqual(out, solution) {
		t.Errorf("Invalid split, current value: %v, need: %v", out, solution)
	}
}

func TestRevertMap(t *testing.T) {
	solution := map[string]string{"one": "n0", "two": "n1", "tree": "n2"}
	in := map[string]string{"n0": "one", "n1": "two", "n2": "tree"}
	out := RevertMap(in)
	if !reflect.DeepEqual(out, solution) {
		t.Errorf("Invalid revert, current value: %v, need: %v", out, solution)
	}
}

func TestFilterStringSlice(t *testing.T) {
	solution := []int{2,3,4,7,8,9}
	in := []int{1,2,3,4,5,6,7,8,9}
	out := FilterIntegerSlice(in)
	if !reflect.DeepEqual(out, solution) {
		t.Errorf("Invalid filter, current value: %v, need: %v", out, solution)
	}
}

func TestFilterIntegerSlice(t *testing.T) {
	solution := []string{"two","three","four","seven","eight"}
	in := []string{"one","two","three","four","five","six","seven","eight"}
	out := FilterStringSlice(in)
	if !reflect.DeepEqual(out, solution) {
		t.Errorf("Invalid filter, current value: %v, need: %v", out, solution)
	}
}