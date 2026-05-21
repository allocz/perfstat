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

func (s *Stats) StartTimer(label string) func(count uint64) {
	return nop 
}

func (s *Stats) MaybePrint() {
	// NOP
}
