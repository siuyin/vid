package main

import (
	"fmt"
	"time"

	"github.com/siuyin/expb"
	"github.com/siuyin/vid"
)

func main() {
	ts := expb.Intervals(time.Duration(2000000000), 3)
	t := time.NewTimer(0)
	<-t.C
	run(0)
	for i := 0; i < len(ts); i++ {
		t.Reset(ts[i])
		<-t.C
		run(i+1)
	}
}

func run(n int) {
	fmt.Printf("Run %d: %s\n", n, time.Now().Format("15:04:05.000"))
	vid.Capture(time.Now().Format("img-loc123-20060102-150405-SGT"))
}
