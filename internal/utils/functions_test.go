package utils

import (
	"reflect"
	"testing"
)

func TestSplitStringSlice(t *testing.T) {
	solution := [][]string{{"one","two","three"},{"four","five","six"},{"seven","eight"}}
	in := []string{"one","two","three","four","five","six","seven","eight"}
	out, _ := SplitStringSlice(in, 3)
	if !reflect.DeepEqual(out, solution) {
		t.Errorf("Invalid split, current value: %v, need: %v", out, solution)
	}
}

func BenchmarkSplitStringSlice(b *testing.B) {
	in := []string{"one","two","three","four","five","six","seven","eight"}
	for i := 0; i < b.N; i++ {
		SplitStringSlice(in, 3)
	}
	b.ReportAllocs()
}

func TestSplitIntegerSlice(t *testing.T) {
	solution := [][]int{{1,2,3},{4,5,6},{7,8}}
	in := []int{1,2,3,4,5,6,7,8}
	out, _ := SplitIntegerSlice(in, 3)
	if !reflect.DeepEqual(out, solution) {
		t.Errorf("Invalid split, current value: %v, need: %v", out, solution)
	}
}

func BenchmarkSplitIntegerSlice(b *testing.B) {
	in := []int{1,2,3,4,5,6,7,8}
	for i := 0; i < b.N; i++ {
		SplitIntegerSlice(in, 3)
	}
	b.ReportAllocs()
}

func TestRevertMap(t *testing.T) {
	solution := map[string]string{"one": "n0", "two": "n1", "tree": "n2"}
	in := map[string]string{"n0": "one", "n1": "two", "n2": "tree"}
	out, _ := RevertMap(in)
	if !reflect.DeepEqual(out, solution) {
		t.Errorf("Invalid revert, current value: %v, need: %v", out, solution)
	}
}

func BenchmarkRevertMap(b *testing.B) {
	in := map[string]string{"n0": "one", "n1": "two", "n2": "tree"}
	for i := 0; i < b.N; i++ {
		RevertMap(in)
	}
	b.ReportAllocs()
}

func TestFilterIntegerSlice(t *testing.T) {
	solution := []int{2,3,4,7,8,9}
	in := []int{1,2,3,4,5,6,7,8,9}
	out, _ := FilterIntegerSlice(in)
	if !reflect.DeepEqual(out, solution) {
		t.Errorf("Invalid filter, current value: %v, need: %v", out, solution)
	}
}

func BenchmarkFilterIntegerSlice(b *testing.B) {
	in := []int{1,2,3,4,5,6,7,8,9}
	for i := 0; i < b.N; i++ {
		FilterIntegerSlice(in)
	}
	b.ReportAllocs()
}

func TestFilterStringSlice(t *testing.T) {
	solution := []string{"two","three","four","seven","eight"}
	in := []string{"one","two","three","four","five","six","seven","eight"}
	out, _ := FilterStringSlice(in)
	if !reflect.DeepEqual(out, solution) {
		t.Errorf("Invalid filter, current value: %v, need: %v", out, solution)
	}
}

func BenchmarkFilterStringSlice(b *testing.B) {
	in := []string{"one","two","three","four","five","six","seven","eight"}
	for i := 0; i < b.N; i++ {
		FilterStringSlice(in)
	}
	b.ReportAllocs()
}