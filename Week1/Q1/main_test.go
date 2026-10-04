package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestPipeline(t *testing.T) {
	// Unit tests for the stages
	in := generator(5, 6)
	out := squarer(in)

	first := <-out
	if first != 25 {
		t.Errorf("Expected 25, got %d", first)
	}

	second := <-out
	if second != 36 {
		t.Errorf("Expected 36, got %d", second)
	}

	_, ok := <-out
	if ok {
		t.Errorf("Expected out channel to be closed")
	}

	// Output test for main
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "4") || !strings.Contains(output, "9") || !strings.Contains(output, "16") {
		t.Errorf("Output is missing expected values. Got:\n%s", output)
	}
}
