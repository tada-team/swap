// +build !race

package swap

import "time"

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

func Bool(p *bool, v bool) func() {
	old := *p
	*p = v
	return func() { *p = old }
}

func True(p *bool) func() {
	return Bool(p, true)
}

func False(p *bool) func() {
	return Bool(p, false)
}

func String(p *string, v string) func() {
	old := *p
	*p = v
	return func() { *p = old }
}

func Strings(p *[]string, v []string) func() {
	old := *p
	*p = v
	return func() { *p = old }
}

func Int(p *int, v int) func() {
	old := *p
	*p = v
	return func() { *p = old }
}

func Int64(p *int64, v int64) func() {
	old := *p
	*p = v
	return func() { *p = old }
}

func Duration(p *time.Duration, v time.Duration) func() {
	old := *p
	*p = v
	return func() { *p = old }
}

func Time(p *time.Time, v time.Time) func() {
	old := *p
	*p = v
	return func() { *p = old }
}
