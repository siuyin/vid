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
	vid.Capture(fmt.Sprintf("img-loc123-%s",time.Now().Format("20060102-150405SGT")))
}
