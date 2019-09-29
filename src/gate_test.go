package circuit

import (
  "testing"
)

func TestSignalGate(t *testing.T) {
  gate := SignalGate{[]Connection{}, HI}
  signal := gate.Evaluate()
  if signal != HI {
    t.Errorf("gate.Evaluate() = %s; expected HI", SignalString(HI))
  }
}

func TestANDGateOnlyHI(t *testing.T) {
  andGate := NewANDGate()
  signalGateOne := NewSignalGate(HI)
  signalGateTwo := NewSignalGate(HI)
  NewConnection(&signalGateOne, &andGate)
  NewConnection(&signalGateTwo, &andGate)
  signal := andGate.Evaluate()
  if signal != HI {
    t.Errorf("ANDGate with only HI input = %s; expected HI", SignalString(signal))
  }
}

func TestANDGateHIAndLOW(t *testing.T) {
  andGate := NewANDGate()
  signalGateOne := NewSignalGate(HI)
  signalGateTwo := NewSignalGate(LOW)
  NewConnection(&signalGateOne, &andGate)
  NewConnection(&signalGateTwo, &andGate)
  signal := andGate.Evaluate()
  if signal != LOW {
    t.Errorf("ANDGate with HI and LOW input = %s; expected LOW", SignalString(signal))
  }
}
