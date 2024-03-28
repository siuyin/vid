package vid

import (
	"testing"
	"time"
)

func TestCapture(t *testing.T) {
	t.Skip("Skipping test. Uncomment to run test")
	output := "img-loc123-" + time.Now().Format("20060101-150405")
	if err := Capture(output); err != nil {
		t.Error(err)
	}
}

func TestFrames(t *testing.T) {
	output := "vid-loc123-" + time.Now().Format("20060102-150405")
	if err := Frames(output, 50); err != nil {
		t.Error(err)
	}
}
