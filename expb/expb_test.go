package expb

import (
	"testing"
	"time"
)

func TestIntervals(t *testing.T) {
	dat := []struct {
		intvl time.Duration
		n     int
		o     []time.Duration
	}{
		{time.Duration(1000000000), 4, []time.Duration{
			time.Duration(1000000000), time.Duration(2000000000),
			time.Duration(4000000000), time.Duration(8000000000)}},

		{time.Duration(1000000000), 2, []time.Duration{
			time.Duration(1000000000), time.Duration(2000000000)}},

		{time.Duration(500000000), 3, []time.Duration{
			time.Duration(500000000), time.Duration(1000000000),
			time.Duration(2000000000)}},
	}
	for i, d := range dat {
		if o := Intervals(d.intvl, d.n); !cmp(o, d.o) {
			t.Errorf("case: %d: %v, %v", i, o, d.o)
		}
	}
}
func cmp(a, b []time.Duration) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
