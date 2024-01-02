package main

func Sum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}
	return
}

func SumAll(numbersToSum ...[]int) (result []int) {
	result = make([]int, len(numbersToSum))
	for i, numbers := range numbersToSum {
		result[i] = Sum(numbers)
	}
	return
}
