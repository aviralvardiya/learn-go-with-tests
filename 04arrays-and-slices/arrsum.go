package arrays

func ArrSum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return (sum)
}

func SumAll(numbersToSum ...[]int) []int {
	// legthOfnumbers := len(numbersToSum)
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, ArrSum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	// legthOfnumbers := len(numbersToSum)
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tails := numbers[1:]
			sums = append(sums, ArrSum(tails))
		}

	}

	return sums
}
