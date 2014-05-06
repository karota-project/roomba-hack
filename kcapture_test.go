package kcapture

import (
	"./kcapture"
	"testing"
)

func TestStart(t *testing.T) {
	/* expected TRUE test */
	actual := kcapture.Start("ffmpeg", nil)
	expected := true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	/* expected TRUE test */
	actual2 := kcapture.Start("ffserver", nil)
	expected2 := true
	if actual2 != expected {
		t.Errorf("got %v\nwant %v", actual2, expected2)
	}
}

func TestStop(t *testing.T) {
	/* expected TRUE test */
	actual := kcapture.Stop("ffmpeg")
	expected := true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	/* expected TRUE test */
	actual2 := kcapture.Stop("ffserver")
	expected2 := true
	if actual2 != expected {
		t.Errorf("got %v\nwant %v", actual2, expected2)
	}
}
