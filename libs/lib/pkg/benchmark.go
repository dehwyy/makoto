// easy tool for benchmarking test in go

package lib

import (
	"fmt"
	"time"
)

type Benchmark struct {
	Name string

	start_time time.Time
	end_time   time.Time
}

func NewBenchmark(opts ...Benchmark) *Benchmark {
	benchmark_name := "func"
	if len(opts) > 0 && opts[0].Name != "" {
		benchmark_name = opts[0].Name
	}

	return &Benchmark{
		Name: benchmark_name,
	}
}

func (b *Benchmark) Start() {
	b.start_time = time.Now()
}

func (b *Benchmark) End() {
	b.end_time = time.Now()
}

func (b *Benchmark) EndAndSummarize(cases_count int) {
	b.end_time = time.Now()
	b.Summarize(cases_count)
}

// returns overall time and average time in ms
func (b *Benchmark) Summarize(cases_count int) {
	overall_time := b.end_time.Sub(b.start_time).Seconds() * 1000
	average_time := overall_time / float64(cases_count)

	// parse
	p := func(f float64) string {
		return fmt.Sprintf("%.4f", f)
	}

	fmt.Printf("%s: \ttime spend: %ss,\tavg time: %ss\n", b.Name, p(overall_time), p(average_time))
}
