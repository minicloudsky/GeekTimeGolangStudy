package fn__test

import (
	"math/rand"
	"testing"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)
}
