package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

func TestWorkerPool(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	start := time.Now()
	main()
	elapsed := time.Since(start)

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "All jobs finished!") {
		t.Errorf("Output misses 'All jobs finished!'. Got:\n%s", output)
	}

	// 5 jobs taking 1 second each with 3 workers should take approximately 2 seconds (1s for first 3 jobs, 1s for next 2)
	// If it took 5 seconds, workers were not running in parallel correctly.
	if elapsed >= 4*time.Second {
		t.Errorf("Worker pool seems too slow (took %v), parallel execution might be failing", elapsed)
	}
}
