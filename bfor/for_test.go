package bfor_test

import (
	"testing"

	"github.com/NasSilverBullet/go-bench-test/bfor"
)

func BenchmarkBFor_LoopWithCopy(b *testing.B) {
	bf := &bfor.BFor{100000}

	s := bf.Gen()

	for n := 0; n < b.N; n++ {
		bf.LoopWithCopy(s)
	}
}

func BenchmarkBFor_LoopMapAccess(b *testing.B) {
	bf := &bfor.BFor{100000}

	s := bf.Gen()

	for n := 0; n < b.N; n++ {
		bf.LoopMapAccess(s)
	}
}
