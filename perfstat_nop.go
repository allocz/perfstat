//go:build !perfstat
package perfstat

import "time"

type Stats struct {
	PrintInterval time.Duration
}

func (s *Stats) Update(label string, totalCost, count uint64) {
	// NOP
}

func nop(count uint64) {
	// NOP
}

func (s *Stats) StartTimer(label string) Timer {
	return Timer{}
}

func (s *Stats) MaybePrint() {
	// NOP
}

func (s *Stats) Reset() {
	// NOP
}

type Timer struct {}

func (t *Timer) Finish(count ...uint64) {
	// NOP
}
