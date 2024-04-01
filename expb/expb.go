// package expb provides exponential back-off functions
package expb

import (
	"math/rand"
	"time"
)

// Intervals returns a slice of n numbers of exponentially backed-off durations.
// Each subsequent duration is double the previous duration.
func Intervals(intvl time.Duration, n int) []time.Duration {
	return IntvlRateFudge(intvl, n, 2.0, 0.00)
}

// IntvlRateFudge returns a slice of numbers of exponentially backed-off durations.
// Each subsequent duration is rate * the previous duration,
// multiplied by 1+fudge*random[-1.0,1.0).
func IntvlRateFudge(intvl time.Duration, n int, rate, fudge float32) []time.Duration {
	durs := []time.Duration{}
	b := time.Duration(float32(intvl) * (1.0 + fudge*(rand.Float32()*2-1.0)))
	for i := 0; i < n; i++ {
		durs = append(durs, b)
		b = time.Duration(rate * float32(b))
	}
	return durs

}
