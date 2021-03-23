package swap

import "testing"

func TestChain(t *testing.T) {
	x := 2
	y := 3

	t.Run("defer", func(t *testing.T) {
		if x != 2 {
			t.Fatal("x must be 2")
		}
		defer Chain(
			Int(&x, 10),
			Int(&y, 20),
		)()
		if x != 10 {
			t.Fatal("x must be 10")
		}
	})

	if x != 2 {
		t.Fatal("x must be 2")
	}
}

func TestInt(t *testing.T) {
	x := 2
	t.Run("defer", func(t *testing.T) {
		if x != 2 {
			t.Fatal("x must be 2")
		}
		defer Int(&x, 10)()
		if x != 10 {
			t.Fatal("x must be 10")
		}
	})
	if x != 2 {
		t.Fatal("x must be 2")
	}
}
