package main

import (
	"testing"
)

func BenchmarkHandGetCompanyNames(b *testing.B) {
	var r []string
	for i := 0; i < b.N; i++ {
		r = byH.getCompanyNames()
	}
	_ = r
}

func BenchmarkLinqGetCompanyNamesT(b *testing.B) {
	var r []string
	for i := 0; i < b.N; i++ {
		r = byLi.getCompanyNamesT()
	}
	_ = r
}

func BenchmarkLinqGetCompanyNames(b *testing.B) {
	var r []string
	for i := 0; i < b.N; i++ {
		r = byLi.getCompanyNames()
	}
	_ = r
}
