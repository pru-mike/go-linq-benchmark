package main

import "testing"

func BenchmarkLinqFindTAvgEmployUSA(b *testing.B) {
	var r float64
	for i := 0; i < b.N; i++ {
		r = byLi.findTAvgEmployUSA()
	}
	_ = r
}

func BenchmarkHandFindAvgEmployeesUSA(b *testing.B) {
	var r float64
	for i := 0; i < b.N; i++ {
		r = byH.findAvgEmployeesUSA()
	}
	_ = r
}
