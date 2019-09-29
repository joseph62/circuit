package circuit

import (
	"testing"
)

func TestNewConnection(t *testing.T) {
	source := NewSignalGate(HI)
	destination := NewSignalGate(LOW)
	connection := NewConnection(&source, &destination)
	if source.Evaluate() != connection.Source.Evaluate() {
		t.Errorf("Connection source not assined as expected")
	}

	if destination.Evaluate() != connection.Destination.Evaluate() {
		t.Errorf("Connection destination not assined as expected")
	}
}
