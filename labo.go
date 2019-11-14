package main

import (
	"fmt"
	"os"

	"github.com/NasSilverBullet/go-bench-test/bfor"
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
	bf := &bfor.BFor{}
	s := bf.Gen()

	bf.LoopWithCopy(s)
	bf.LoopMapAccess(s)

	return nil
}
