package bmap_test

import (
	"testing"

	"github.com/NasSilverBullet/go-bench-test/bmap"
)

func BenchmarkBMap_GenHugeMap(b *testing.B) {
	bm := &bmap.BMap{}
	m := bm.Gen()
	for n := 0; n < b.N; n++ {
		bm.GetLoop(m)
	}
}
