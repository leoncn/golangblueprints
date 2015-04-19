package trace

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer

	tracer := New(&buf)

	if tracer == nil {
		t.Error("Return from New should not be null")
	} else {
		tracer.Trace("Hello trace.")

		if buf.String() != "Hello trace." {
			t.Errorf("Trace should not write '%s'.", buf.String())
		}
	}
}
