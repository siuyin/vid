package vid

import "testing"

func TestCapture(t *testing.T) {
	if err := Capture(); err != nil {
		t.Error(err)
	}
}
