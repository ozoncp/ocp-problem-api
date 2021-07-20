package utils

var stringInnerList = []string{"one", "five", "six"}
var integerInnerList = []int{1,5,6}

func SplitStringSlice(in []string, size int) (out [][]string) {
	inSize := len(in)
	if size >= inSize {
		out = append(out, in)
		return
	}

	endIndex := 0
	for startIndex := 0; startIndex < inSize; startIndex+=size {
		endIndex = startIndex + size
		if endIndex >= inSize {
			out = append(out, in[startIndex:])
		} else {
			out = append(out, in[startIndex:endIndex])
		}
	}

	return
}

func SplitIntegerSlice(in []int, size int) (out [][]int) {
	inSize := len(in)
	if size >= inSize {
		out = append(out, in)
		return
	}

	endIndex := 0
	for startIndex := 0; startIndex < inSize; startIndex+=size {
		endIndex = startIndex + size
		if endIndex >= inSize {
			out = append(out, in[startIndex:])
		} else {
			out = append(out, in[startIndex:endIndex])
		}

	}

	return
}

func RevertMap(in map[string]string) map[string]string {
	out := make(map[string]string)
	for key := range in {
		out[in[key]] = key
	}

	return out
}

func FilterStringSlice(in []string) (out []string) {
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

	return
}

func FilterIntegerSlice(in []int) (out []int) {
	lastIndexInList := len(integerInnerList) - 1
	for i := range in {
		for j := range stringInnerList {
			if integerInnerList[j] == in[i] {
				break
			}

			if lastIndexInList == j {
				out = append(out, in[i])
			}
		}
	}

	return
}