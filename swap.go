package swap

import (
	"sync"
	"time"
)

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

func True(p *bool) func() { return Bool(p, true) }

func False(p *bool) func() { return Bool(p, false) }

func Bool(p *bool, v bool) func() {
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

func GetBool(p *bool) bool {
	mux.RLock()
	defer mux.RUnlock()
	return *p
}

func String(p *string, v string) func() {
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

func GetString(p *string) string {
	mux.RLock()
	defer mux.RUnlock()
	return *p
}

func Strings(p *[]string, v []string) func() {
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

func GetStrings(p *[]string) []string {
	mux.RLock()
	defer mux.RUnlock()
	return *p
}

func Int(p *int, v int) func() {
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

func GetInt(p *int) int {
	mux.RLock()
	defer mux.RUnlock()
	return *p
}

func Int64(p *int64, v int64) func() {
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

func GetInt64(p *int64) int64 {
	mux.RLock()
	defer mux.RUnlock()
	return *p
}

func Duration(p *time.Duration, v time.Duration) func() {
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

func GetDuration(p *time.Duration) time.Duration {
	mux.RLock()
	defer mux.RUnlock()
	return *p
}

func Time(p *time.Time, v time.Time) func() {
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

func GetTime(p *time.Time) time.Time {
	mux.RLock()
	defer mux.RUnlock()
	return *p
}
