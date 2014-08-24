package kcapture

import (
	"./kcapture"
	"testing"
)

func TestStart(t *testing.T) {
	/* expected TRUE test */
	actual := kcapture.Start("ffmpeg", nil)
	var expected error = nil
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	/* expected TRUE test */
	actual = kcapture.Start("ffserver", nil)
	expected = nil
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestStop(t *testing.T) {
	/* expected TRUE test */
	actual := kcapture.Stop("ffmpeg")
	var expected error = nil
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	/* expected TRUE test */
	actual = kcapture.Stop("ffserver")
	expected = nil
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
