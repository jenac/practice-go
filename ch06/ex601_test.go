package ch06

import (
	"fmt"
	"os"
	"testing"
)

func TestHello(t *testing.T) {
	result := Hello()
	want := "Hello, World"
	if result == want {
		t.Logf("ok: Hello() = %v, want %v", result, want)
	} else {
		t.Errorf("fail: Hello() = %v, want %v", result, want)
	}

	want2 := "Hello, World"
	if result == want2 {
		t.Logf("ok: Hello() = %v, want %v", result, want)
	} else {
		t.Errorf("fail: Hello() = %v, want %v", result, want)
	}
}

func testHello(expected string) func(t *testing.T) {
	return func(t *testing.T) {
		if expected == "Hello, World" {
			t.Logf("Hello() = %v, want = %v", expected, "Hello World")
		} else {
			t.Errorf("Hello() = %v, want = %v", expected, "Hello World")
		}
	}
}

func TestHello2(t *testing.T) {
	t.Run("test for hello with run", testHello(Hello()))
}

func TestMain(m *testing.M) {
	fmt.Println("*****Befroe*****")
	code := m.Run()
	fmt.Println("******After*****")
	os.Exit(code)
}
