package arrays

/*
An interesting property of arrays is that the size is encoded in its type.
If you try to pass an [4]int into a function that expects [5]int, it won't compile
*/
func SumFixed(numbers [5]int) int {
	sum := 0
	// for i := 0; i < 5; i++ {
	// 	sum += numbers[i]
	// }
	//range lets you iterate over an array. On each iteration, range returns two values - the index and the value
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// slice way
func Sum(numbers []int) int {
	sum := 0
	//range lets you iterate over an array. On each iteration, range returns two values - the index and the value
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	// result := make([]int, len(slices))
	// for i, sl := range slices {
	// result[i] = Sum(sl)
	// }
	var result []int
	for _, sl := range slices {
		result = append(result, Sum(sl))
	}
	return result
}

func SumAllTails(slices ...[]int) []int {
	var result []int
	for _, sl := range slices {
		var tail []int
		if len(sl) == 0 {
			tail = nil
		} else {
			tail = sl[1:]
		}
		result = append(result, Sum(tail))
	}
	return result
}
