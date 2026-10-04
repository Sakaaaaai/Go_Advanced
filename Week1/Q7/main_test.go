package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestMemoryLeakPprof(t *testing.T) {
	// テスト用に memoryLeak の無限ループは実行せず、サーバーの起動チェックと動作のみを確認する
	go func() {
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		// 今回は main() ではなく、一部だけ実行
		go func() {
			http.ListenAndServe(":6060", nil)
		}()
		
		go func() {
			data := make([]byte, 1024*1024)
			leakySlice = append(leakySlice, data)
		}()
		
		time.Sleep(2 * time.Second)

		w.Close()
		os.Stdout = oldStdout
		var buf bytes.Buffer
		io.Copy(&buf, r)
	}()

	time.Sleep(500 * time.Millisecond)

	resp, err := http.Get("http://localhost:6060/debug/pprof/heap")
	if err != nil {
		t.Fatalf("Failed to connect to pprof server: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK from pprof server, got %v", resp.StatusCode)
	}

	// Make sure the leakySlice has data
	if len(leakySlice) == 0 {
		t.Errorf("Expected leakySlice to have some elements, got 0")
	}
}
