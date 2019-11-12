package bmap

import "math/rand"

var benchcount = 100000

type M struct {
	Count int
}

type Foo struct {
	Id         uint
	ForeignKey uint
}

func (b *M) GenHugeMap() map[uint]Foo {
	c := b.Count
	if c == 0 {
		c = benchcount
	}

	m := make(map[uint]Foo, c)

	for i := 0; i < c; i++ {
		m[uint(i+1)] = Foo{uint(rand.Int()), uint(i + 1)}
	}

	return m
}
