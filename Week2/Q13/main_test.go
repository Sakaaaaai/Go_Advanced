package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestCacheLogic(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "Cache Miss!") {
		t.Errorf("Expected cache miss on first call, got:\n%s", output)
	}

	if !strings.Contains(output, "Cache Hit!") {
		t.Errorf("Expected cache hit on second call, got:\n%s", output)
	}
}
