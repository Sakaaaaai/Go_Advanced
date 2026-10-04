package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestPubSubMock(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "API: User 1 registered.") {
		t.Errorf("Output misses main API log")
	}

	if !strings.Contains(output, "Sending email to: user1@example.com") ||
		!strings.Contains(output, "Sending email to: user2@example.com") {
		t.Errorf("Subscriber did not process emails correctly. Got:\n%s", output)
	}

	if !strings.Contains(output, "Subscriber closed.") {
		t.Errorf("Subscriber did not close properly")
	}
}
