package ch06

import "testing"

func TestHello(t *testing.T) {
	result:=Hello()
	want:="Hello, World"
	if result == want {
		t.Log("Hello() = %v, want %v", result, want)
	} else {
		t.Errorf("Hello() = %v, want %v", result, want)
	}
}
