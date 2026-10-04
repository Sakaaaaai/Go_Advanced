package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

func TestFanOutFanIn(t *testing.T) {
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

	// 順序は保証されないので全ての値が含まれているかチェック
	expectedValues := []string{"1", "4", "9", "16", "25"}
	for _, v := range expectedValues {
		if !strings.Contains(output, v+"\n") {
			t.Errorf("Output misses expected value %s. Got:\n%s", v, output)
		}
	}

	// 5 items * 100ms = 500ms if single threaded.
	// With 2 workers, it should take ~300ms.
	if elapsed >= 450*time.Millisecond {
		t.Errorf("Execution seems too slow (%v), Fan-out might not be working correctly running in parallel.", elapsed)
	}
}
