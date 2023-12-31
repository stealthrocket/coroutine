package coroutine

import (
	"reflect"
	"testing"
)

func TestLocalStorage(t *testing.T) {
	execute(42, func() {
		if v := load(); !reflect.DeepEqual(v, 42) {
			t.Errorf("wrong value: %v", v)
		}
	})
}

//go:noinline
func weirdLoop(n int, f func()) int {
	if n == 0 {
		f()
		return 0
	} else {
		return weirdLoop(n-1, f) + 1 // just in case Go ever implements tail recursion
	}
}

func TestLocalStorageGrowStack(t *testing.T) {
	execute("hello", func() {
		weirdLoop(100e3, func() {
			if v := load(); v != "hello" {
				t.Errorf("wrong value: %v", v)
			}
		})
	})
}

func BenchmarkLocalStorage(b *testing.B) {
	execute("hello", func() {
		for i := 0; i < b.N; i++ {
			load()
		}
	})
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New[any, any](func() { _ = i }).Next()
	}
}
