package bmap_test

import (
	"reflect"
	"testing"

	"github.com/NasSilverBullet/go-bench-test/bmap"
)

func Test_genHugeMap(t *testing.T) {
	tests := []struct {
		name string
		want map[uint]bmap.Foo
	}{
		{"Success", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bmap.GenHugeMap(0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genHugeMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
