# swap variables

```go
// myconfig.go
var myConfig struct {
	Foo string
	Bar int
}

func init() {
	readConfig(&myConfig)
}
```

```go
// test.go
import "github.com/tada-team/swap"

func TestWithSwap(t *testing.T) {
	defer swap.Value(&myConfig.Foo, "test value")()
	defer swap.Value(&myConfig.Bar, 42)()
	// ...test cases...
}

// more sugar
func TestWithSwapChain(t *testing.T) {
	defer swap.Chain(
		swap.Value(&myConfig.Foo, "test value"),
		swap.Value(&myConfig.Bar, 42),
	)()  
	// ...test cases...
}
```
