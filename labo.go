package main

import (
	"fmt"
	"os"

	"github.com/NasSilverBullet/go-bench-test/bmap"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stdout, "Success!")
	os.Exit(0)
}

func run() error {
	b := &bmap.BMap{}
	hm := b.GenHugeMap()
	_ = hm

	return nil
}
