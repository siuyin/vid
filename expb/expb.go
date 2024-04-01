// package expb provides exponential back-off functions
package expb

import "time"

func Intervals(intvl time.Duration, n int) []time.Duration {
	durs := []time.Duration{}
	b := intvl
	for i := 0; i < n; i++ {
		durs = append(durs, b)
		b = 2 * b
	}
	return durs
}
