package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestMigrationStrings(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()
	
	outputLower := strings.ToLower(output)

	if !strings.Contains(outputLower, "create table users") {
		t.Errorf("UpMigration does not seem to create 'users' table.")
	}
	
	if !strings.Contains(outputLower, "drop table if exists users") {
		t.Errorf("DownMigration does not seem to drop 'users' table correctly.")
	}
}
