package sum

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(nums ...[]int) []int {
	var sums []int

	for _, numbers := range nums {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
