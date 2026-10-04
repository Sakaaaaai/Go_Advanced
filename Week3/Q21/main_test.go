package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestWebSocketHandler(t *testing.T) {
	// グローバル変数のリセット
	readCount = 0

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "Upgraded to WebSocket!") {
		t.Errorf("Does not seem to upgrade to websocket")
	}

	if !strings.Contains(output, "Echoing message: Hello WebSocket!") {
		t.Errorf("Did not echo the correct message")
	}

	if !strings.Contains(output, "Read error (or disconnected): client disconnected") {
		t.Errorf("Did not handle disconnect error properly")
	}

	if !strings.Contains(output, "Connection closed.") {
		t.Errorf("Did not defer connection close")
	}
}
