package others

func GetFibonacciValue(fibonacciIndex int) int {
	if fibonacciIndex < 0 {
		panic("negative value")
	}
	if fibonacciIndex == 0 {
		return 0
	}

	if fibonacciIndex == 1 {
		return 1
	}

	return GetFibonacciValue(fibonacciIndex-1) + GetFibonacciValue(fibonacciIndex-2)
}
