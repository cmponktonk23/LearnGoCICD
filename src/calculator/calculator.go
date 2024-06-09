package calculator

type Calculator struct {
	a int
	b int
}

func (cal *Calculator) Assign(a, b int) {
	cal.a, cal.b = a, b
}

func (cal *Calculator) Add() int {
	return cal.a + cal.b
}

func (cal *Calculator) Divide() int {
	if cal.b == 0 {
		panic("divide by 0")
	}
	return cal.a / cal.b
}

func MakeCalculator() Calculator {
	return Calculator{0, 0}
}
