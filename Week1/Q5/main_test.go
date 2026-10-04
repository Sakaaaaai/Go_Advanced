package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestEscapeAnalysis(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "Stack value: 10") {
		t.Errorf("Expected output to contain 'Stack value: 10', got: %q", output)
	}

	if !strings.Contains(output, "Heap value: 20") {
		t.Errorf("Expected output to contain 'Heap value: 20', got: %q", output)
	}
}
