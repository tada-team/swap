package swap

import "sync"

var mux sync.RWMutex

func Chain(fns ...func()) func() {
	return func() {
		for i := range fns {
			fns[i]()
		}
	}
}

func RevChain(fns ...func()) func() {
	return func() {
		for i := len(fns) - 1; i >= 0; i-- {
			fns[i]()
		}
	}
}

func True(p *bool) func() { return Value(p, true) }

func False(p *bool) func() { return Value(p, false) }

func Value[T any](p *T, v T) func() {
	mux.Lock()
	defer mux.Unlock()

	old := *p
	*p = v

	return func() {
		mux.Lock()
		defer mux.Unlock()
		*p = old
	}
}

func GetValue[T any](p *T) T {
	mux.RLock()
	defer mux.RUnlock()
	return *p
}
