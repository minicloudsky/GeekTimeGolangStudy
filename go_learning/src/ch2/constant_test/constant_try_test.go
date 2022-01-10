package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Satursday
	Sunday
)
const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstanttry(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConstanttry1(t testing.T) {
	t.Log(Readable, Writable, Executable)
}
