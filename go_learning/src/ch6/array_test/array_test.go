package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 3, 4, 5}
	arr2 := [...]int{1, 3, 5, 7, 9}
	t.Log(arr[1], arr[2])
	t.Log(arr1)
	t.Log(arr2)
}

func TestArraytrave(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 9}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	for i, val := range arr3 {
		t.Log(i, " ", val)
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 3, 5, 7, 9, 11}
	arr3_sec := arr3[0:3]
	t.Log(arr3_sec)
	t.Log(arr3_sec[len(arr3_sec)-1])
}
