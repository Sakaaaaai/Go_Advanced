package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestTracingSimulation(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "[Trace] Span 'FetchDatabase' started.") {
		t.Errorf("Tracer start is missing or incorrect name")
	}

	if !strings.Contains(output, "[Trace] Error recorded on span 'FetchDatabase': database timeout") {
		t.Errorf("Error was not recorded onto the span")
	}

	if !strings.Contains(output, "[Trace] Span 'FetchDatabase' ended.") {
		t.Errorf("Span was not deferred to end")
	}
}
