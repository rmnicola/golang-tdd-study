package main

func Sum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}
	return
}

func SumAll(numbersToSum ...[]int) (result []int) {
	for _, numbers := range numbersToSum {
		result = append(result, Sum(numbers))
	}
	return
}

func SumAllTails(numbersToSum ...[]int) (result []int) {
	for _, numbers := range numbersToSum {
		tail := numbers[1:]
		result = append(result, Sum(tail))
	}
	return
}
