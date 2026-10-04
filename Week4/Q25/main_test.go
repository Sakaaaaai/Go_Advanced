package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestIntegrationMock(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "コンテナを破棄します") {
		t.Errorf("Cleanup was not executed. Output:\n%s", output)
	}

	if !strings.Contains(output, "CREATE TABLE users") {
		t.Errorf("Missing integration test logic. Output:\n%s", output)
	}
}
