package arrays

func Reduce[A any](collection []A, accumulator func(A, A) A, initialValue A) A {
	var result = initialValue
	for _, x := range collection {
		result = accumulator(result, x)
	}
	return result
}


func SumGenerics(numbers []int) int {
	add := func(acc, x int) int { return acc + x }
	return Reduce(numbers, add, 0)
}

func SumAllTailsGenerics(numbers ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, SumGenerics(tail))
		}
	}

	return Reduce(numbers, sumTail, []int{})
}