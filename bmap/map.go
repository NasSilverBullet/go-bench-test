package bmap

import "math/rand"

var benchcount = 100000

type BMap struct {
	Count int
}

type Foo struct {
	Id         uint
	ForeignKey uint
}

func (b *BMap) getBenchCount() int {
	if b.Count > 0 {
		return b.Count
	}
	return benchcount
}

func (b *BMap) Gen() map[uint]Foo {
	c := b.getBenchCount()

	m := make(map[uint]Foo, c)
	for i := 0; i < c; i++ {
		m[uint(i+1)] = Foo{uint(rand.Int()), uint(i + 1)}
	}

	return m
}

func (b *BMap) GetLoop(m map[uint]Foo) {
	c := b.getBenchCount()

	for i := 0; i < c; i++ {
		_ = m[uint(i+1)]
	}
}
