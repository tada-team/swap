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

// test.go
import "github.com/tada-team/swap"

func TestWithSwap(t *testing.T) {
	defer swap.Bool(&myConfig.Foo, "test value")()
	defer swap.Int(&myConfig.Bar, 42)()
	// ...test cases...
}

// more sugar
func TestWithSwapChain(t *testing.T) {
    defer swap.Chain(
    	swap.Bool(&myConfig.Foo, "test value"),
        swap.Int(&myConfig.Bar, 42),
    )()  
    // ...test cases...
}
```
