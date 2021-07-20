package utils

import "errors"

var stringInnerList = []string{"one", "five", "six"}
var integerInnerList = []int{1,5,6}

func SplitStringSlice(in []string, size uint) ([][]string, error) {
	if in == nil {
		return nil, errors.New("invalid slice")
	}

	if size == 0 {
		return nil, errors.New("invalid size")
	}

	inSize := len(in)
	if int(size) >= inSize {
		return [][]string{in}, nil
	}

	out := make([][]string, 0, (inSize/int(size))+1)
	for startIndex, endIndex := 0, 0; startIndex < inSize; startIndex+=int(size) {
		endIndex = startIndex + int(size)
		if endIndex >= inSize {
			out = append(out, in[startIndex:])
		} else {
			out = append(out, in[startIndex:endIndex])
		}
	}

	return out, nil
}

func SplitIntegerSlice(in []int, size uint) ([][]int, error) {
	if in == nil {
		return nil, errors.New("invalid slice")
	}

	if size == 0 {
		return nil, errors.New("invalid size")
	}

	inSize := len(in)
	if int(size) >= inSize {
		return [][]int{in}, nil
	}

	out := make([][]int, 0, (inSize/int(size))+1)
	for startIndex, endIndex := 0, 0; startIndex < inSize; startIndex+=int(size) {
		endIndex = startIndex + int(size)
		if endIndex >= inSize {
			out = append(out, in[startIndex:])
		} else {
			out = append(out, in[startIndex:endIndex])
		}

	}

	return out, nil
}

func RevertMap(in map[string]string) (map[string]string, error) {
	if in == nil {
		return nil, errors.New("invalid map")
	}

	out := make(map[string]string, len(in))
	for key, value := range in {
		out[value] = key
	}

	return out, nil
}

func FilterStringSlice(in []string) ([]string, error) {
	if in == nil {
		return nil, errors.New("invalid slice")
	}

	out := make([]string, 0, len(in))
	lastIndexInList := len(stringInnerList) - 1
	for i := range in {
		for j := range stringInnerList {
			if stringInnerList[j] == in[i] {
				break
			}

			if lastIndexInList == j {
				out = append(out, in[i])
			}
		}
	}

	return out, nil
}

func FilterIntegerSlice(in []int) ([]int, error) {
	if in == nil {
		return nil, errors.New("invalid slice")
	}

	out := make([]int, 0, len(in))
	lastIndexInList := len(integerInnerList) - 1
	for i := range in {
		for j := range integerInnerList {
			if integerInnerList[j] == in[i] {
				break
			}

			if lastIndexInList == j {
				out = append(out, in[i])
			}
		}
	}

	return out, nil
}