package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestSyncOnceAndMap(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Check that initialization message appears ONLY once
	initMsg := "System initialized! (This should only appear once)"
	count := strings.Count(output, initMsg)
	if count != 1 {
		t.Errorf("Expected initialization message exactly once, found it %d times", count)
	}

	// Check if the result was correctly loaded from the sync.Map
	if !strings.Contains(output, "Result: Data for ID") {
		t.Errorf("Output misses expected map result string. Got:\n%s", output)
	}
}
