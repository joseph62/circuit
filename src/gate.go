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

type ANDGate struct {
	Connections []Connection
}

func (gate *ANDGate) AddConnection(connection Connection) {
	gate.Connections = append(gate.Connections, connection)
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
	return ANDGate{[]Connection{}}
}
