package arrays

func Sum(numbers [5]int) int {
	sum := 0
	for i, _ := range numbers {
		sum += numbers[i]
	}
	return sum
}
