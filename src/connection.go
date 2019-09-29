package circuit

type Connection struct {
	Source      Gate
	Destination Gate
}

func (c Connection) Evaluate() Signal {
	return c.Source.Evaluate()
}

func NewConnection(source Gate, destination Gate) Connection {
	connection := Connection{source, destination}
	source.AddConnection(connection)
	destination.AddConnection(connection)
	return connection
}
