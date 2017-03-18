package trace

// Tracer is the interface that describes an object capable of tracing events througout code.
type Tracer interface {
	Trace(...interface{})
}
