package perfstat_test

import (
	"testing"
	"time"

	"github.com/allocz/perfstat"
)

func TestExampleSimpleUsage(t *testing.T) {
	var s perfstat.Stats
	s.PrintInterval = -1

	rootFunc := func() {
		end := s.StartTimer("rootFunc")

		f1 := func() {
			time.Sleep(time.Millisecond * 50)
		}
		f2 := func() {
			time.Sleep(time.Millisecond * 150)
		}
		f3 := func() {
			time.Sleep(time.Millisecond * 800)
		}

		endf := s.StartTimer("f1")
		f1()
		endf.Finish()

		endf = s.StartTimer("f2")
		f2()
		endf.Finish()

		endf = s.StartTimer("f3")
		f3()
		endf.Finish()

		end.Finish()
	}

	rootFunc()

	s.MaybePrint()
}

func TestExampleLoopUsage(t *testing.T) {
	var s perfstat.Stats
	s.PrintInterval = -1

	rootFunc := func() {
		end := s.StartTimer("rootFunc")

		f1 := func() {
			time.Sleep(time.Millisecond * 50)
		}
		f2 := func() {
			time.Sleep(time.Millisecond * 150)
		}
		f3 := func() {
			time.Sleep(time.Millisecond * 8)
		}

		endf := s.StartTimer("f1")
		f1()
		endf.Finish()

		endf = s.StartTimer("f2")
		f2()
		endf.Finish()

		endf = s.StartTimer("f3")
		for range 100 {
			f3()
		}
		// doing this instead of calling StartTimer() and endf() for 
		// each iteration reduces synchronization overhead
		endf.Finish()

		end.Finish()
	}

	rootFunc()

	s.MaybePrint()
}
