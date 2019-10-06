package circuit

type Gate interface {
	AddConnection(Connection)
	Evaluate() Signal
}

type SignalGate struct {
	Connections []Connection
	Signal      Signal
}

func (gate *SignalGate) AddConnection(connection Connection) {
	gate.Connections = append(gate.Connections, connection)
}

func (gate SignalGate) Evaluate() Signal {
	return gate.Signal
}

func NewSignalGate(signal Signal) SignalGate {
	return SignalGate{[]Connection{}, signal}
}

type ORGate struct {
  Connections [2]Connection
  ConnectionCount int
}

func (gate *ORGate) AddConnection(connection Connection) {
  if gate.ConnectionCount < 2 {
    gate.Connections[gate.ConnectionCount] = connection
    gate.ConnectionCount++
  }
}

func (gate ORGate) Evaluate() Signal {
  for i := 0; i < len(gate.Connections); i++ {
    connection := gate.Connections[i]
    signal := connection.Evaluate()
    if signal == HI {
      return HI
    }
  }
  return LOW
}

func NewORGate() ORGate {
  return ORGate{[2]Connection{}, 0}
}

type ANDGate struct {
  Connections [2]Connection
  ConnectionCount int
}

func (gate *ANDGate) AddConnection(connection Connection) {
  if gate.ConnectionCount < 2 {
    gate.Connections[gate.ConnectionCount] = connection
    gate.ConnectionCount++
  }
}

func (gate ANDGate) Evaluate() Signal {
	for i := 0; i < len(gate.Connections); i++ {
		connection := gate.Connections[i]
		signal := connection.Evaluate()
		if signal == LOW {
			return LOW
		}
	}
	return HI
}

func NewANDGate() ANDGate {
	return ANDGate{[2]Connection{}, 0}
}
