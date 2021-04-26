package recursion

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func CallRecur(n int){
	if n > 0 {
		CallRecur(n-1)
		fmt.Println(n)
	}
}

func TestRecursivity(t *testing.T) {
	CallRecur(3)
}


/*
	x1 = 0				func(3)
	x1 = 1	 		func(2) + x1= 19
	x1 = 2		func(1) + x1= 6
	x1 = 3	func(0) + x1= 3
				0
*/
func TestRecursionWithGlobalVariable(t *testing.T) {
	n:=3
	result := make([]int, n)
	RecursionWithStaticValue(3, result)
	fmt.Println(result)
	assert.Equal(t, result, []int{3,6,9})

}

var flagtestsFactorial = []struct {
	in  int
	out int
}{
	{0, 1},
	{1, 1},
	{3, 6},
	{10, 3628800},
}

func TestFactorial(t *testing.T) {
	for i, tt := range flagtestsFactorial {
		val := fmt.Sprintf("Test case # %v", i)
		t.Run(val, func(t *testing.T) {
			result := Factorial(tt.in)
			assert.Equal(t, tt.out, result)
		})
	}
}


var flagtestsPower = []struct {
	val  int
	exp int
	out int
}{
	{0, 0,1},
	{0, 1,0},
	{0, 5,0},
	{1, 0,1},
	{1, 5,1},
	{1, 4,1},
	{2, 0,1},
	{2, 1,2},
	{2, 3,8},
	{12, 4,20736},
	{12, 7,35831808},



}

func TestPower(t *testing.T) {
	for i, tt := range flagtestsPower {
		val := fmt.Sprintf("Test case # %v", i)
		t.Run(val, func(t *testing.T) {
			result := Power(tt.val, tt.exp)
			assert.Equal(t, tt.out, result)
		})
	}
}


var flagtestsTaylor = []struct {
	val  float32
	degree float32
	out float32
}{
	{1, 1,2},
	{1, 0,1},
	{1, 2,2.5},
	{3, 3,13},
	{1, 10,2.718282},


}

func TestTaylorSeries(t *testing.T) {

	for i, tt := range flagtestsTaylor {
		val := fmt.Sprintf("Test case # %v", i)
		t.Run(val, func(t *testing.T) {
			result := CalculateTaylorSeriesForExponential(tt.val, tt.degree)
			assert.Equal(t, tt.out, result)
		})
	}
}

var flagtestsTaylor2 = []struct {
	val  float32
	degree float32
	out float32
}{
	{1, 1,2},
	{1, 0,1},
	{1, 2,2.5},
	{3, 3,13},
	{1, 10,2.7182817},


}

func TestCalculateTaylorSeriesForExponentialWithHornerRule(t *testing.T) {

	for i, tt := range flagtestsTaylor2 {
		val := fmt.Sprintf("Test case # %v", i)
		t.Run(val, func(t *testing.T) {
			result := CalculateTaylorSeriesForExponentialWithHornerRule(tt.val, tt.degree)
			assert.Equal(t, tt.out, result)
		})
	}

	for i, tt := range flagtestsTaylor2 {
		val := fmt.Sprintf("Test case # %v", i)
		t.Run(val, func(t *testing.T) {
			result := CalculateTaylorSeriesForExponentialWithHornerRuleWithLoop(tt.val, tt.degree)
			assert.Equal(t, tt.out, result)
		})
	}
}


var flagtestsFibonacci = []struct {
	in  int
	out int
}{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
	{8, 21},
}

func TestFibonacci(t *testing.T) {

	for i, tt := range flagtestsFibonacci {
		val := fmt.Sprintf("Test case # %v", i)
		t.Run(val, func(t *testing.T) {
			result := FibonacciRecursively(tt.in)
			assert.Equal(t, tt.out, result)
		})
	}

	for i, tt := range flagtestsFibonacci {
		val := fmt.Sprintf("Test case # %v", i)
		t.Run(val, func(t *testing.T) {
			result := FibonacciIteratively(tt.in)
			assert.Equal(t, tt.out, result)
		})
	}

	for i, tt := range flagtestsFibonacci {
		val := fmt.Sprintf("Test case # %v", i)
		t.Run(val, func(t *testing.T) {
			result := FibonacciWithMemoization(tt.in)
			assert.Equal(t, tt.out, result)
		})
	}
}