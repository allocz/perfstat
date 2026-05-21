//go:build perfstat
package perfstat

import (
	"bytes"
	"fmt"
	"slices"
	"sync"
	"time"
)

type Stats struct {
	mu sync.Mutex
	data map[string]statItem
	lastPrint time.Time

	PrintInterval time.Duration
}

type statItem struct {
	totalCost uint64
	count uint64
}

func (s *Stats) Update(label string, totalCost, count uint64) {
	s.mu.Lock()

	if s.data == nil {
		s.data = make(map[string]statItem, 1)
	}

	item := s.data[label]

	item.totalCost += totalCost
	item.count += count

	s.data[label] = item

	s.mu.Unlock()
}

func (s *Stats) StartTimer(label string) func(count uint64) {
	start := time.Now()
	return func(count uint64) {
		s.Update(label, uint64(time.Since(start).Nanoseconds()), count)
	}
}

func (s *Stats) MaybePrint() {
	s.mu.Lock()

	if s.data == nil {
		s.mu.Unlock()
		return
	}

	if s.PrintInterval == 0 {
		s.PrintInterval = time.Second * 10
	}

	if s.lastPrint.IsZero() {
		s.lastPrint = time.Now()
	}

	if time.Since(s.lastPrint) < s.PrintInterval {
		s.mu.Unlock()
		return
	}

	s.lastPrint = time.Now()

	items := make([]reportItem, len(s.data))
	var i = 0;
	for k, v := range s.data {
		items[i] = reportItem{
			label:     k,
			totalCost: v.totalCost,
			count:     v.count,
		}
		i++
	}

	fmt.Println(reportItems(items))

	s.mu.Unlock()
}

type reportItem struct {
	label string
	totalCost uint64
	count uint64
}

type reportItems []reportItem

func (s reportItems) String() string {
	var buf bytes.Buffer
	var it *reportItem

	sLen := len(s)
	if sLen == 0 {
		return ""
	}

	slices.SortFunc([]reportItem(s), func(a, b reportItem) int {
		// sort descending
		return int(b.totalCost - a.totalCost)
	})

	it = &s[0]
	fmt.Fprintf(&buf, "%s totalCost %d iterCost %d relativeCost %.4f\n",
		it.label, it.totalCost, it.totalCost/it.count, 1.0)	
	
	for i := 1; i < sLen; i++ {
		it = &s[i]
		fmt.Fprintf(&buf,
			"\t%s totalCost %d iterCost %d relativeCost %.4f\n",
			it.label, it.totalCost, it.totalCost/it.count,
			float64(it.totalCost)/float64(s[0].totalCost))	
	}
	
	return buf.String()
}
