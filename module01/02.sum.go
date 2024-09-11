package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	// sum := 0
	// for _, val := range numbers {
	// 	sum += val
	// }
	// return sum
	if numbers == nil || len(numbers) <= 0 {
		return 0
	} 
	return numbers[0] + Sum(numbers[1:])
}

// Sum([]int{}) => 0
// Sum([]int{3,2,1}) => 3 + Sum([]int{2,1}) => 3 + 2 + Sum([]int{1}) => 3 + 2 + 1 + Sum([]int{})