package recursion

var x1 = 0

/*
	x1 = 0				func(3)
	x1 = 1	 		func(2) + x1= 19
	x1 = 2		func(1) + x1= 6
	x1 = 3	func(0) + x1= 3
				0
 */
func RecursionWithStaticValue(n int, result []int) int {
	if n > 0 {
		x1++
		val:= RecursionWithStaticValue(n-1, result) + x1
		result[n-1]= val
		return val
	}
	return 0
}


func Factorial(n int) int {
	result := 1
	if n == 0 {
		return result
	}
	for i:=1; i<=n; i++ {
		result = result*i
	}
	return result
}


func Power(m, exp int ) int {
	if exp == 0 {
		return 1
	}
	if exp == 1 {
		return m
	}
	if exp % 2 == 0 {
		return Power(m*m, exp/2)
	}
	return m*Power(m*m, (exp-1)/2)
}


type memo struct {
	f, p float32
}

func CalculateTaylorSeriesForExponential(x, degree float32) float32{
	return calculateTaylorSeriesForExponentialUtil(x, degree, &memo{1,1})
}

func calculateTaylorSeriesForExponentialUtil(x float32, degree float32, m *memo) float32 {
	if degree == 0 {
		return 1
	}
	result :=  calculateTaylorSeriesForExponentialUtil(x, degree-1, m)
	m.f = m.f*degree
	m.p = x*m.p
	return result + m.p/m.f
}

func CalculateTaylorSeriesForExponentialWithHornerRuleWithLoop(x, degree float32) float32 {
	result:=float32(1)
	for degree > 0 {
		result = 1 + x*result/degree
		degree--
	}
	return result
}

func CalculateTaylorSeriesForExponentialWithHornerRule(x, degree float32) float32 {
	s := float32(1)
	return calculateTaylorSeriesForExponentialWithHornerRuleUtil(x, degree, s)
}

func calculateTaylorSeriesForExponentialWithHornerRuleUtil(x, degree float32, s float32) float32 {
	if degree == 0 {
		return s
	}
	s = 1 + (x/degree)*s
	return calculateTaylorSeriesForExponentialWithHornerRuleUtil(x, degree-1, s)
}