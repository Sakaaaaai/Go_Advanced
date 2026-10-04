package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestGRPCClient(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "User found: Alice (from mock server)") {
		t.Errorf("Expected output. Got:\n%s", output)
	}
}
