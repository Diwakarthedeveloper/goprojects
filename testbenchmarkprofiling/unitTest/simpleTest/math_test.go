package math

import "testing"

func TestMul(t *testing.T) { // here funtion name will be Testfunctionname, Input parameter will always be *testing.T

	result := Mul(1, 2)
	if result != 3 {
		t.Errorf("Mul(1, 2) = %d; want 3", result) // if the out put of input 2, 3 is not 6 then throw error
	}
}
