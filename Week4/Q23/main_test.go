package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestDependencyInjection(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "Service is working with its injected repository!") {
		t.Errorf("Service was not correctly constructed and invoked via DI.")
	}

	srv := InitializeService()
	if srv == nil || srv.repo == nil || srv.repo.db == nil {
		t.Errorf("Dependencies were not correctly injected inside the Service object tree.")
	}
}
