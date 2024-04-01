package main

import (
	"fmt"
	"time"

	"github.com/siuyin/expb"
	"github.com/siuyin/vid"
)

func main() {
	ts := expb.Intervals(time.Duration(1000000000), 4)
	run(0)
	t := time.NewTimer(ts[0])
	for i := 1; i < len(ts); i++ {
		<-t.C
		t.Reset(ts[i])
		run(i)
	}
}

func run(n int) {
	fmt.Printf("Run %d: %s\n", n, time.Now().Format("15:04:05.000"))
	vid.Capture(time.Now().Format("img-loc123-20060102-150405-SGT"))
}
