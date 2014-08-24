package kcapture

import (
	"testing"
)

func TestStart(t *testing.T) {
	/* expected TRUE test */
	actual := StartStreamer(nil)
	var expected error = nil
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	/* expected TRUE test */
	actual = StartServer(nil)
	expected = nil
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestStop(t *testing.T) {
	/* expected TRUE test */
	actual := StopStreamer()
	var expected error = nil
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	/* expected TRUE test */
	actual = StopServer()
	expected = nil
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
