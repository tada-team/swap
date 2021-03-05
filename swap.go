package swap

import "time"

func Bool(p *bool, v bool) func() {
	old := *p
	*p = v
	return func() { *p = old }
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

func Duration(p *time.Duration, v time.Duration) func() {
	old := *p
	*p = v
	return func() { *p = old }
}
