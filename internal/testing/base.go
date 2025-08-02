package testing

import "testing"

func SetBase() base { return base{} }

type base struct{}

func (b base) Benchmark(bench *testing.B) { bench.ReportAllocs() }
