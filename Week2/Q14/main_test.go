package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestDistributedLock(t *testing.T) {
	rdb := &redisClient{locks: make(map[string]bool)}
	
	processCalled := false
	task := func() {
		processCalled = true
	}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// 1st Execution (should succeed)
	success1 := ExecuteWithLock(rdb, task)
	if !success1 || !processCalled {
		t.Errorf("First execution failed. success: %v, processCalled: %v", success1, processCalled)
	}
	
	// Ensure lock was released
	if rdb.locks["lock:billing"] {
		t.Errorf("Lock was not released after first execution")
	}

	// 2nd Execution with simulated pre-existing lock (should fail)
	rdb.locks["lock:billing"] = true
	processCalled = false
	
	success2 := ExecuteWithLock(rdb, task)
	if success2 || processCalled {
		t.Errorf("Second execution should have been skipped. success: %v, processCalled: %v", success2, processCalled)
	}

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "Lock acquired!") || !strings.Contains(output, "Already running on another instance") {
		t.Errorf("Output misses expected log messages. Got:\n%s", output)
	}
}
