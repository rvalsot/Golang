package trace

import "io"

// Tracer is the interface that describes an object capable of tracing events througout code.
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {}

// New returns a tracer
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}
