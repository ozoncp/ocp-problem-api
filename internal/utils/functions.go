package utils

var stringInnerList = []string{"one", "five", "six"}
var integerInnerList = []int{1,5,6}

func SplitStringSlice(in []string, size int) [][]string {
	inSize := len(in)
	if size >= inSize {
		return [][]string{in}
	}

	out := make([][]string, 0, (inSize/size)+1)
	for startIndex, endIndex := 0, 0; startIndex < inSize; startIndex+=size {
		endIndex = startIndex + size
		if endIndex >= inSize {
			out = append(out, in[startIndex:])
		} else {
			out = append(out, in[startIndex:endIndex])
		}
	}

	return out
}

func SplitIntegerSlice(in []int, size int) [][]int {
	inSize := len(in)
	if size >= inSize {
		return [][]int{in}
	}

	out := make([][]int, 0, (inSize/size)+1)
	for startIndex, endIndex := 0, 0; startIndex < inSize; startIndex+=size {
		endIndex = startIndex + size
		if endIndex >= inSize {
			out = append(out, in[startIndex:])
		} else {
			out = append(out, in[startIndex:endIndex])
		}

	}

	return out
}

func RevertMap(in map[string]string) map[string]string {
	out := make(map[string]string, len(in))
	for key, value := range in {
		out[value] = key
	}

	return out
}

func FilterStringSlice(in []string) []string {
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

	return out
}

func FilterIntegerSlice(in []int) []int {
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

	return out
}