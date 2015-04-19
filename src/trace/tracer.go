// trace project trace.go
package trace

import (
	"fmt"
	"io"
)

type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	w io.Writer
}

func (t tracer) Trace(a ...interface{}) {
	t.w.Write([]byte(fmt.Sprint(a...)))
}

func New(w io.Writer) Tracer {
	return &tracer{w}
}
