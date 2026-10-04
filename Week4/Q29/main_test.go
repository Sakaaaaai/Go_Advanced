package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestGRPCErrorHandling(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "code = 3 desc = ID is required") {
		t.Errorf("InvalidArgument error missing or incorrect. Got: %s", output)
	}

	if !strings.Contains(output, "code = 5 desc = User not found") {
		t.Errorf("NotFound error missing or incorrect. Got: %s", output)
	}

	if !strings.Contains(output, "user=&{Id:1 Name:Alice}") && !strings.Contains(output, "Id:1 Name:Alice") {
		t.Errorf("Success case missing or incorrect user data. Got: %s", output)
	}
}
