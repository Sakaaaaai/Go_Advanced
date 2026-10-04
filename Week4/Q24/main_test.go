package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestMockingLogic(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "user={ID:1 Name:Alice}, err=<nil>") {
		t.Errorf("Normal case log is incorrect. Got: %s", output)
	}

	if !strings.Contains(output, "user={ID: Name:}, err=user not found") {
		t.Errorf("Error case log is incorrect. Got: %s", output)
	}
}
