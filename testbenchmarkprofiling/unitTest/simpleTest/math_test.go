package math

import "testing"

func TestMul(t *testing.T) { // here funtion name will be Testfunctionname, Input parameter will always be *testing.T

	result := Mul(2, 3)
	if result != 6 {
		t.Errorf("Mul(2, 3) = %d; want 6", result) // if the out put of input 2, 3 is not 6 then throw error
	}
}
