package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

func TestPprofServer(t *testing.T) {
	// Pprofサーバーが立ち上がるかテストする
	// メイン関数を非同期で少しだけ動かして、サーバーが起動するのを待つ
	go func() {
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		// mainを呼ぶが、時間がかかるので適当にタイムアウトさせるか途中で止めるのは難しい
		// 今回はメインアプリの動作をモック化していないため、直接httpを取得しにいくアプローチをとる
		main()

		w.Close()
		os.Stdout = oldStdout
		var buf bytes.Buffer
		io.Copy(&buf, r)
	}()

	// サーバーが起動するのを少し待つ
	time.Sleep(500 * time.Millisecond)

	// pprofのエンドポイントにアクセスできるか確認
	resp, err := http.Get("http://localhost:6060/debug/pprof/")
	if err != nil {
		t.Fatalf("Failed to connect to pprof server: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK from pprof server, got %v", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	if !strings.Contains(string(body), "Types of profiles available:") {
		t.Errorf("Response does not seem to be from pprof. Got:\n%s", string(body))
	}
}
