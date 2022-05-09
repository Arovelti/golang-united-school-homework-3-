package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	var n, sum float32
	for _, v := range input {
		if v != 0 {
			n++
		}
		sum += v
	}
	result = sum / n
	return result
}
