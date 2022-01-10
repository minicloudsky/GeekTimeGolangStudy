package stringfn_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	t.Logf("parts: %v", parts)
	t.Logf("join: %v", strings.Join(parts, "-"))
}
func TestingConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str: ", s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log("i: ", 10+i)
	}
}
