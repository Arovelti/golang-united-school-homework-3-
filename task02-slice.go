package homework

func reverse(input []int64) (result []int64) {
	//Place your code here
	inputLen := len(input)
	result = make([]int64, inputLen)

	for index, num := range input {
		r := inputLen - index - 1
		result[r] = num
	}
	return result
}
