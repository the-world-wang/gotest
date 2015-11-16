package even

import (
	"testing"
)

func TestEven(t *testing.T) {
	if !Even(10) {
		t.Log("10 must be even")
		t.Fail()
	}
}

func TestOdd(t *testing.T) {
	if !Odd(7) {
		t.Log("7 mus be odo")
		t.Fail()
	}
}
