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

func SumAllTails(nums ...[]int) []int {
	var sums []int

	for _, numbers := range nums {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
