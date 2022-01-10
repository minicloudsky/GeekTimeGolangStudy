package string_test

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello"
	t.Log("len: ", len(s))
	// s[1] = '3'
	s = "\xE4\xBA\xBB\xFF"
	t.Log(s)
	t.Log("len: ", len(s))
}

func TestUTF8(t *testing.T) {
	s := "中"
	t.Log(" 中 len； ", len(s))
	// c := []rune[s]
	// t.Logf("中 unicode %x",c[0])
	// t.Logf("中 UTF8 %x",s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("test string to rune %[1]c %[1]c", c)
	}
}
