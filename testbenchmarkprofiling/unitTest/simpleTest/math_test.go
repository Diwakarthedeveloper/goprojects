package math

import "testing"

func TestMul(t *testing.T) {
	result := Mul(2, 3)
	if result != 6 {
		t.Errorf("Mul(2, 3) = %d; want 6", result)
	}
}
