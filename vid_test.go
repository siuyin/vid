package vid

import (
	"testing"
	"time"
)

func TestCapture(t *testing.T) {
	output := "img-loc123-" + time.Now().Format("20060304-150405")
	if err := Capture(output); err != nil {
		t.Error(err)
	}
}
