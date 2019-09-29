package circuit

import (
	"testing"
)

func TestSignalStringLOW(t *testing.T) {
	representation := SignalString(LOW)
	if representation != "LOW" {
		t.Errorf("SignalString(LOW) = %q; want \"LOW\"", representation)
	}
}

func TestSignalStringHI(t *testing.T) {
	representation := SignalString(HI)
	if representation != "HI" {
		t.Errorf("SignalString(HI) = %q; want \"HI\"", representation)
	}
}

func TestFlipSignalLOWToHI(t *testing.T) {
	flipped := FlipSignal(LOW)
	if flipped != HI {
		t.Errorf("FlipSignal(LOW) = %s; want HI", SignalString(flipped))
	}
}

func TestFlipSignalHIToLOW(t *testing.T) {
	flipped := FlipSignal(HI)
	if flipped != LOW {
		t.Errorf("FlipSignal(HI) = %s; want LOW", SignalString(flipped))
	}
}
