package main

func Sum(numbers []int) (sum int) {
	sum = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	//lengthOfNumbers := len(numbersToSum)
	//sums = make([]int, lengthOfNumbers)

	for _, numbers := range numbersToSum {
		sums = append(sums,Sum(numbers))
	}
	return sums
}
