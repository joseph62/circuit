package circuit

type Signal uint64

const (
	HI Signal = iota
	LOW
)

func FlipSignal(signal Signal) Signal {
	if signal == HI {
		return LOW
	} else {
		return HI
	}
}

func SignalString(signal Signal) string {
	if signal == HI {
		return "HI"
	} else {
		return "LOW"
	}
}
