package perfstat_test

import (
	"testing"

	"github.com/allocz/perfstat"
)

// As a Global to avoid compiler optimizing this out
var s perfstat.Stats

func BenchmarkPerfstatStartTimer(b *testing.B) {
	s = perfstat.Stats{}
	for range b.N {
		endR := s.StartTimer("root")

		end := s.StartTimer("label1")
		end.Finish()

		end = s.StartTimer("label2")
		end.Finish()

		end = s.StartTimer("label3")
		end.Finish()

		endR.Finish()
	}
}
