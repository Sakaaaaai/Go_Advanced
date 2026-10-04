package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestGRPCStreaming(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "[Stream Sending...] Log 1: Server started") ||
	   !strings.Contains(output, "[Stream Sending...] Log 4: Shutting down") {
		t.Errorf("Stream did not send the expected logs. Got: %s", output)
	}

	if !strings.Contains(output, "Stream finished successfully.") {
		t.Errorf("Stream did not finish properly")
	}
}
