package ReverseSlice

func reverse(input []int) []int {
	var output = make([]int, 0)
	if len(input) <= 0 {
		return output
	} else {
		for i := len(input) - 1; i >= 0; i-- {
			output = append(output, input[i])
		}
	}
	return output
}
