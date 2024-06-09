package calculator

import "testing"

func CalculatorAddCommon(a, b int, t *testing.T) {
	calculator := MakeCalculator()
	calculator.Assign(a, b)
	expect, result := a+b, calculator.Add()
	if result != expect {
		t.Fatalf("Add result wrong, expect: %d, get: %d", expect, result)
	}
}

func CalculatorDivideCommon(a, b int, t *testing.T) {
	calculator := MakeCalculator()
	calculator.Assign(a, b)
	expect, result := a/b, calculator.Divide()
	if result != expect {
		t.Fatalf("Divide result wrong, expect: %d, get: %d", expect, result)
	}
}

func TestCalculator_Add_Positives(t *testing.T) {
	CalculatorAddCommon(1, 2, t)
}

func TestCalculator_Add_Negatives(t *testing.T) {
	CalculatorAddCommon(-10, -2, t)
}

func TestCalculator_Add_MixResultPositive(t *testing.T) {
	CalculatorAddCommon(10, -2, t)
}

func TestCalculator_Add_MixResultNegative(t *testing.T) {
	CalculatorAddCommon(-10, 2, t)
}

func TestCalculator_Divide_Positives(t *testing.T) {
	CalculatorDivideCommon(4, 2, t)
}

func TestCalculator_DividebyZero(t *testing.T) {
	t.Run("Test Panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expect Panic but not")
			}
		}()

		calculator := MakeCalculator()
		calculator.Assign(1, 0)
		calculator.Divide()
	})
}
