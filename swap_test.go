package swap

import "testing"

func TestValueAndChain(t *testing.T) {
	someInt := 2
	someString := "3"
	someStrings := []string{"3"}
	someBool := true

	t.Run("defer", func(t *testing.T) {
		if someInt != 2 {
			t.Fatal("someInt must be 2")
		}
		defer Chain(
			Value(&someInt, 10),
			Value(&someString, "20"),
			Value(&someBool, false),
			Value(&someStrings, []string{"fff"}),
		)()
		if someInt != 10 {
			t.Fatal("someInt must be 10")
		}
	})

	if someInt != 2 {
		t.Fatal("someInt must be 2")
	}
}
